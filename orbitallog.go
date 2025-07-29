package orbitallog

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sync"
	"time"
)

type Logger struct {
	mu      sync.Mutex
	file    *os.File
	logger  *log.Logger
	logDir  string
	prefix  string
	maxAge  time.Duration
	logPath string
	logDate string
}

// New cria o logger com limpeza inicial
func New(logDir, prefix string, maxAge time.Duration) (*Logger, error) {
	// Garante que a pasta existe
	if err := os.MkdirAll(logDir, 0755); err != nil {
		return nil, fmt.Errorf("erro ao criar pasta de logs: %w", err)
	}

	l := &Logger{
		logDir: logDir,
		prefix: prefix,
		maxAge: maxAge,
	}

	// Limpa arquivos antigos
	if err := l.cleanupOldFiles(); err != nil {
		return nil, err
	}

	// Abre ou cria o log de hoje
	if err := l.openLogFile(); err != nil {
		return nil, err
	}

	return l, nil
}

// openLogFile abre ou cria o log de hoje
func (l *Logger) openLogFile() error {
	l.logDate = time.Now().Format("02-01-2006")
	l.logPath = filepath.Join(l.logDir, fmt.Sprintf("%s_%s.log", l.prefix, l.logDate))

	file, err := os.OpenFile(l.logPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("erro ao abrir log: %w", err)
	}
	l.file = file
	l.logger = log.New(file, "", log.LstdFlags)
	return nil
}

// cleanupOldFiles remove arquivos mais antigos que maxAge
func (l *Logger) cleanupOldFiles() error {
	files, err := filepath.Glob(filepath.Join(l.logDir, fmt.Sprintf("%s_*.log", l.prefix)))
	if err != nil {
		return err
	}

	cutoff := time.Now().Add(-l.maxAge)
	for _, f := range files {
		info, err := os.Stat(f)
		if err == nil && info.ModTime().Before(cutoff) {
			_ = os.Remove(f)
		}
	}
	return nil
}

// Printf escreve no log com segurança para concorrência
func (l *Logger) Printf(format string, v ...any) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.logger.Printf(format, v...)
}

// Close fecha o arquivo de log
func (l *Logger) Close() error {
	l.mu.Lock()
	defer l.mu.Unlock()
	if l.file != nil {
		return l.file.Close()
	}
	return nil
}
