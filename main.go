package main

import (
	"bytes"
	"net/http"

	"github.com/jung-kurt/gofpdf"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.POST("/account", func(c echo.Context) error {
		return c.String(http.StatusOK, "Cadastrando uma conta")

	})

	e.GET("/accounts", func(c echo.Context) error {
		return c.String(http.StatusOK, "Listando todas as contas")

	})

	e.GET("/accounts/:accountId", func(c echo.Context) error {
		accountId := c.Param("accountId")
		return c.String(http.StatusOK, "Listando a conta "+accountId)

	})

	e.PUT("/accounts/:accountId", func(c echo.Context) error {
		accountId := c.Param("accountId")
		return c.String(http.StatusOK, "Alterando uma conta existente "+accountId)

	})

	e.GET("/accounts/:accountId/cards", func(c echo.Context) error {
		accountId := c.Param("accountId")
		return c.String(http.StatusOK, "Listando todos os cartões de uma conta "+accountId)

	})

	e.POST("/card", func(c echo.Context) error {
		return c.String(http.StatusOK, "Cadastrando um cartão")

	})

	e.GET("/cards", func(c echo.Context) error {
		return c.String(http.StatusOK, "Listando todos os cartões")

	})

	e.GET("/cards/:cardId", func(c echo.Context) error {
		cardId := c.Param("cardId")
		return c.String(http.StatusOK, "Listando um cartão "+cardId)

	})

	e.PUT("/cards/:cardId", func(c echo.Context) error {
		cardId := c.Param("cardId")
		return c.String(http.StatusOK, "Alterando um cartão existente "+cardId)

	})

	e.GET("/cards/:cardId/transactions", func(c echo.Context) error {
		cardId := c.Param("cardId")
		return c.String(http.StatusOK, "Listando todas as transações de um cartao "+cardId)

	})

	e.POST("/transaction", func(c echo.Context) error {
		return c.String(http.StatusOK, "Cadastrando uma transação ")

	})

	e.GET("/transactions", func(c echo.Context) error {
		return c.String(http.StatusOK, "Listando todas as transações ")

	})

	e.GET("/transactions/:transactionId", func(c echo.Context) error {
		transactionId := c.Param("transactionId")
		return c.String(http.StatusOK, "Listando uma transação pelo Id da transação "+transactionId)

	})

	e.GET("/contas/:accountId/transacoes.pdf", GetTransactionsPDF)

	e.Logger.Fatal(e.Start(":8080"))

}

func GetTransactionsPDF(c echo.Context) error {
	accountID := c.Param("accountId")

	transactions := []struct {
		Data  string
		Valor string
		Loja  string
	}{
		{"2023-09-01", "199.50", "Apple Store"},
		{"2023-09-05", "27.50", "Netflix"},
	}

	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 16)
	pdf.Cell(40, 10, "Transações da Conta "+accountID)
	pdf.Ln(12)

	pdf.SetFont("Arial", "B", 12)
	pdf.Cell(40, 10, "Data")
	pdf.Cell(60, 10, "Loja")
	pdf.Cell(40, 10, "Valor")
	pdf.Ln(12)

	pdf.SetFont("Arial", "", 12)
	for _, tx := range transactions {
		pdf.Cell(40, 10, tx.Data)
		pdf.Cell(60, 10, tx.Loja)
		pdf.Cell(40, 10, tx.Valor)
		pdf.Ln(10)
	}

	var buf bytes.Buffer
	err := pdf.Output(&buf)
	if err != nil {
		return err
	}

	return c.Blob(http.StatusOK, "application/pdf", buf.Bytes())
}
