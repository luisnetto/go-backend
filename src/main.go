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
