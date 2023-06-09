﻿package main

import (
	"PARTNER/lib"
	"PARTNER/rotas"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

var (
	logger *lib.Logger
)

func ImprimirNome() {
	fmt.Println(lib.GetLog())
	fmt.Println("Port:" + lib.GetPort())
}

func main() {
	logger = lib.GetLogger("main")
	ImprimirNome()
	err := lib.LerEnv()
	if err != nil {
		logger.Errorf("Erro ao ler .env : %v", err)
		panic(err.Error())
	}
	// Inicializar o roteador
	router := gin.Default()
	router.Use(lib.CorsMiddleware())
	router.POST("/signin", rotas.PostSignIn)
	router.POST("/signemp", rotas.PostSignEmp)
	router.GET("/signemp", rotas.GetSignEmp)
	router.POST("/CLUB0001/carga/", rotas.CLUB0001_carga)
	router.POST("/CLUB0001/lookup/", rotas.CLUB0001_lookup)
	router.POST("/CLUB0001/lista/", rotas.CLUB0001_lista)
	router.POST("/CLUB0001/insert/", rotas.CLUB0001_insert)
	router.GET("/CLUB0001/row/:id", rotas.CLUB0001_row_id)
	router.PUT("/CLUB0001/update/:id", rotas.CLUB0001_update_id)
	router.DELETE("/CLUB0001/delete/:id", rotas.CLUB0001_delete_id)
	router.POST("/CLUB0002/carga/", rotas.CLUB0002_carga)
	router.POST("/CLUB0002/lookup/", rotas.CLUB0002_lookup)
	router.POST("/CLUB0002/lista/", rotas.CLUB0002_lista)
	router.POST("/CLUB0002/insert/", rotas.CLUB0002_insert)
	router.GET("/CLUB0002/row/:id", rotas.CLUB0002_row_id)
	router.PUT("/CLUB0002/update/:id", rotas.CLUB0002_update_id)
	router.DELETE("/CLUB0002/delete/:id", rotas.CLUB0002_delete_id)
	router.POST("/CLUB0003/carga/", rotas.CLUB0003_carga)
	router.POST("/CLUB0003/lookup/", rotas.CLUB0003_lookup)
	router.POST("/CLUB0003/lista/", rotas.CLUB0003_lista)
	router.POST("/CLUB0003/insert/", rotas.CLUB0003_insert)
	router.GET("/CLUB0003/row/:id", rotas.CLUB0003_row_id)
	router.PUT("/CLUB0003/update/:id", rotas.CLUB0003_update_id)
	router.DELETE("/CLUB0003/delete/:id", rotas.CLUB0003_delete_id)
	router.POST("/CLUB0004/carga/", rotas.CLUB0004_carga)
	router.POST("/CLUB0004/lookup/", rotas.CLUB0004_lookup)
	router.POST("/CLUB0004/lista/", rotas.CLUB0004_lista)
	router.POST("/CLUB0004/insert/", rotas.CLUB0004_insert)
	router.GET("/CLUB0004/row/:id", rotas.CLUB0004_row_id)
	router.PUT("/CLUB0004/update/:id", rotas.CLUB0004_update_id)
	router.DELETE("/CLUB0004/delete/:id", rotas.CLUB0004_delete_id)
	router.POST("/EMPRESAS/carga/", rotas.EMPRESAS_carga)
	router.POST("/EMPRESAS/lookup/", rotas.EMPRESAS_lookup)
	router.POST("/EMPRESAS/lista/", rotas.EMPRESAS_lista)
	router.POST("/EMPRESAS/insert/", rotas.EMPRESAS_insert)
	router.GET("/EMPRESAS/row/:id", rotas.EMPRESAS_row_id)
	router.PUT("/EMPRESAS/update/:id", rotas.EMPRESAS_update_id)
	router.DELETE("/EMPRESAS/delete/:id", rotas.EMPRESAS_delete_id)
	router.POST("/FIN0001/carga/", rotas.FIN0001_carga)
	router.POST("/FIN0001/lookup/", rotas.FIN0001_lookup)
	router.POST("/FIN0001/lista/", rotas.FIN0001_lista)
	router.POST("/FIN0001/insert/", rotas.FIN0001_insert)
	router.GET("/FIN0001/row/:id", rotas.FIN0001_row_id)
	router.PUT("/FIN0001/update/:id", rotas.FIN0001_update_id)
	router.DELETE("/FIN0001/delete/:id", rotas.FIN0001_delete_id)
	router.POST("/FIN0002/carga/", rotas.FIN0002_carga)
	router.POST("/FIN0002/lookup/", rotas.FIN0002_lookup)
	router.POST("/FIN0002/lista/", rotas.FIN0002_lista)
	router.POST("/FIN0002/insert/", rotas.FIN0002_insert)
	router.GET("/FIN0002/row/:id", rotas.FIN0002_row_id)
	router.PUT("/FIN0002/update/:id", rotas.FIN0002_update_id)
	router.DELETE("/FIN0002/delete/:id", rotas.FIN0002_delete_id)
	router.POST("/FIN0003/carga/", rotas.FIN0003_carga)
	router.POST("/FIN0003/lookup/", rotas.FIN0003_lookup)
	router.POST("/FIN0003/lista/", rotas.FIN0003_lista)
	router.POST("/FIN0003/insert/", rotas.FIN0003_insert)
	router.GET("/FIN0003/row/:id", rotas.FIN0003_row_id)
	router.PUT("/FIN0003/update/:id", rotas.FIN0003_update_id)
	router.DELETE("/FIN0003/delete/:id", rotas.FIN0003_delete_id)
	router.POST("/FIN0004/carga/", rotas.FIN0004_carga)
	router.POST("/FIN0004/lookup/", rotas.FIN0004_lookup)
	router.POST("/FIN0004/lista/", rotas.FIN0004_lista)
	router.POST("/FIN0004/insert/", rotas.FIN0004_insert)
	router.GET("/FIN0004/row/:id", rotas.FIN0004_row_id)
	router.PUT("/FIN0004/update/:id", rotas.FIN0004_update_id)
	router.DELETE("/FIN0004/delete/:id", rotas.FIN0004_delete_id)
	router.POST("/FIN0005/carga/", rotas.FIN0005_carga)
	router.POST("/FIN0005/lookup/", rotas.FIN0005_lookup)
	router.POST("/FIN0005/lista/", rotas.FIN0005_lista)
	router.POST("/FIN0005/insert/", rotas.FIN0005_insert)
	router.GET("/FIN0005/row/:id", rotas.FIN0005_row_id)
	router.PUT("/FIN0005/update/:id", rotas.FIN0005_update_id)
	router.DELETE("/FIN0005/delete/:id", rotas.FIN0005_delete_id)
	router.POST("/FIN0006/carga/", rotas.FIN0006_carga)
	router.POST("/FIN0006/lookup/", rotas.FIN0006_lookup)
	router.POST("/FIN0006/lista/", rotas.FIN0006_lista)
	router.POST("/FIN0006/insert/", rotas.FIN0006_insert)
	router.GET("/FIN0006/row/:id", rotas.FIN0006_row_id)
	router.PUT("/FIN0006/update/:id", rotas.FIN0006_update_id)
	router.DELETE("/FIN0006/delete/:id", rotas.FIN0006_delete_id)
	router.POST("/FIN0007/carga/", rotas.FIN0007_carga)
	router.POST("/FIN0007/lookup/", rotas.FIN0007_lookup)
	router.POST("/FIN0007/lista/", rotas.FIN0007_lista)
	router.POST("/FIN0007/insert/", rotas.FIN0007_insert)
	router.GET("/FIN0007/row/:id", rotas.FIN0007_row_id)
	router.PUT("/FIN0007/update/:id", rotas.FIN0007_update_id)
	router.DELETE("/FIN0007/delete/:id", rotas.FIN0007_delete_id)
	router.POST("/FIN0008/carga/", rotas.FIN0008_carga)
	router.POST("/FIN0008/lookup/", rotas.FIN0008_lookup)
	router.POST("/FIN0008/lista/", rotas.FIN0008_lista)
	router.POST("/FIN0008/insert/", rotas.FIN0008_insert)
	router.GET("/FIN0008/row/:id", rotas.FIN0008_row_id)
	router.PUT("/FIN0008/update/:id", rotas.FIN0008_update_id)
	router.DELETE("/FIN0008/delete/:id", rotas.FIN0008_delete_id)
	router.POST("/FIN0009/carga/", rotas.FIN0009_carga)
	router.POST("/FIN0009/lookup/", rotas.FIN0009_lookup)
	router.POST("/FIN0009/lista/", rotas.FIN0009_lista)
	router.POST("/FIN0009/insert/", rotas.FIN0009_insert)
	router.GET("/FIN0009/row/:id", rotas.FIN0009_row_id)
	router.PUT("/FIN0009/update/:id", rotas.FIN0009_update_id)
	router.DELETE("/FIN0009/delete/:id", rotas.FIN0009_delete_id)
	router.POST("/TB0001/carga/", rotas.TB0001_carga)
	router.POST("/TB0001/lookup/", rotas.TB0001_lookup)
	router.POST("/TB0001/lista/", rotas.TB0001_lista)
	router.POST("/TB0001/insert/", rotas.TB0001_insert)
	router.GET("/TB0001/row/:id", rotas.TB0001_row_id)
	router.PUT("/TB0001/update/:id", rotas.TB0001_update_id)
	router.DELETE("/TB0001/delete/:id", rotas.TB0001_delete_id)
	router.POST("/UF/carga/", rotas.UF_carga)
	router.POST("/UF/lookup/", rotas.UF_lookup)
	router.POST("/UF/lista/", rotas.UF_lista)
	router.POST("/UF/insert/", rotas.UF_insert)
	router.GET("/UF/row/:id", rotas.UF_row_id)
	router.PUT("/UF/update/:id", rotas.UF_update_id)
	router.DELETE("/UF/delete/:id", rotas.UF_delete_id)
	router.POST("/USER/carga/", rotas.USER_carga)
	router.POST("/USER/lookup/", rotas.USER_lookup)
	router.POST("/USER/lista/", rotas.USER_lista)
	router.POST("/USER/insert/", rotas.USER_insert)
	router.GET("/USER/row/:id", rotas.USER_row_id)
	router.PUT("/USER/update/:id", rotas.USER_update_id)
	router.DELETE("/USER/delete/:id", rotas.USER_delete_id)
	router.POST("/USEREMP/carga/", rotas.USEREMP_carga)
	router.POST("/USEREMP/lookup/", rotas.USEREMP_lookup)
	router.POST("/USEREMP/lista/", rotas.USEREMP_lista)
	router.POST("/USEREMP/insert/", rotas.USEREMP_insert)
	router.GET("/USEREMP/row/:id", rotas.USEREMP_row_id)
	router.PUT("/USEREMP/update/:id", rotas.USEREMP_update_id)
	router.DELETE("/USEREMP/delete/:id", rotas.USEREMP_delete_id)
	// Iniciar o servidor                      
	log.Fatal(router.Run(":" + lib.GetPort()))
}
