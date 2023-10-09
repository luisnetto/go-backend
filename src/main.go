package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Tarefa struct {
	ID        string `json:"id"`
	Descricao string `json:"descricao"`
	Concluida bool   `json:"concluida"`
}

var tarefas = make(map[string]Tarefa)

func main() {

	router := gin.Default()

	router.POST("/tarefas", criarTarefa)
	router.GET("/tarefas", buscarTarefas)
	router.GET("/tarefas/:id", buscarTarefa)
	router.DELETE("/tarefas/:id", excluirTarefa)
	router.PUT("/tarefas/:id", editarTarefa)
	router.GET("contagem-tarefas", contagemTarefas)
	router.Run(":3000")
}

func criarTarefa(c *gin.Context) {
	var tarefa Tarefa
	if err := c.ShouldBindJSON(&tarefa); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	tarefa.ID = uuid.New().String()
	tarefas[tarefa.ID] = tarefa
	c.JSON(http.StatusCreated, tarefa)
}

func buscarTarefas(c *gin.Context) {
	itens := make([]Tarefa, 0, len(tarefas))
	for _, item := range tarefas {
		itens = append(itens, item)
	}

	c.JSON(http.StatusOK, itens)
}

func buscarTarefa(c *gin.Context) {
	id := c.Param("id")
	tarefa, existe := tarefas[id]
	if !existe {
		c.JSON(http.StatusNotFound, gin.H{"error": "Tarefa não Encontrada!"})
		return
	}
	c.JSON(http.StatusOK, tarefa)
}

func excluirTarefa(c *gin.Context) {
	id := c.Param("id")
	_, existe := tarefas[id]
	if !existe {
		c.JSON(http.StatusNotFound, gin.H{"error": "Tarefa não Encontrada!"})
		return
	}
	delete(tarefas, id)
	c.JSON(http.StatusOK, gin.H{"message": "Tarefa excluída"})
}

func editarTarefa(c *gin.Context) {
	var tarefa Tarefa
	id := c.Param("id")
	if err := c.ShouldBindJSON(&tarefa); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	_, existe := tarefas[id]
	if !existe {
		c.JSON(http.StatusNotFound, gin.H{"error": "Tarefa não Encontrada!"})
		return
	}
	tarefas[id] = tarefa
	c.JSON(http.StatusOK, tarefa)
}
func contagemTarefas(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"quantidade_tarefas": len(tarefas)})
}
