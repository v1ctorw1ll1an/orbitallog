package orbitallog

import (
	"os"
	"path/filepath"
	"testing"
	"time"
)

func TestNewLoggerCreatesFile(t *testing.T) {
	dir := t.TempDir()
	prefix := "testapp"
	maxAge := 24 * time.Hour

	logger, err := New(dir, prefix, maxAge)
	if err != nil {
		t.Fatalf("erro ao criar logger: %v", err)
	}
	defer logger.Close()

	// Espera que o arquivo de log do dia exista
	expectedFile := filepath.Join(dir, prefix+"_"+time.Now().Format("02-01-2006")+".log")
	if _, err := os.Stat(expectedFile); os.IsNotExist(err) {
		t.Errorf("arquivo de log não foi criado: %s", expectedFile)
	}
}

func TestLoggerWritesToFile(t *testing.T) {
	dir := t.TempDir()
	logger, _ := New(dir, "testapp", 24*time.Hour)
	defer logger.Close()

	logger.Printf("Mensagem de teste %d", 123)

	logFile := filepath.Join(dir, "testapp_"+time.Now().Format("02-01-2006")+".log")
	content, err := os.ReadFile(logFile)
	if err != nil {
		t.Fatalf("erro ao ler arquivo de log: %v", err)
	}

	if string(content) == "" {
		t.Errorf("conteúdo do log está vazio")
	}
}

func TestCleanupOldFiles(t *testing.T) {
	dir := t.TempDir()
	prefix := "oldlog"
	maxAge := 24 * time.Hour

	// Criar um arquivo antigo
	oldFile := filepath.Join(dir, prefix+"_01-01-2000.log")
	if err := os.WriteFile(oldFile, []byte("antigo"), 0644); err != nil {
		t.Fatalf("erro ao criar log antigo: %v", err)
	}

	// Forçar modificação antiga
	oldTime := time.Now().Add(-48 * time.Hour)
	if err := os.Chtimes(oldFile, oldTime, oldTime); err != nil {
		t.Fatalf("erro ao alterar data do log antigo: %v", err)
	}

	logger, _ := New(dir, prefix, maxAge)
	defer logger.Close()

	// Rodar limpeza manual
	if err := logger.cleanupOldFiles(); err != nil {
		t.Fatalf("erro na limpeza: %v", err)
	}

	// Arquivo deve ter sido removido
	if _, err := os.Stat(oldFile); err == nil {
		t.Errorf("arquivo antigo não foi removido")
	}
}

func TestLoggerClose(t *testing.T) {
	dir := t.TempDir()
	logger, _ := New(dir, "closeapp", 24*time.Hour)
	if err := logger.Close(); err != nil {
		t.Errorf("erro ao fechar logger: %v", err)
	}
}
