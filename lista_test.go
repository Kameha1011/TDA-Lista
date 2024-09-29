package lista_test

import (
	TDALista "tdas/lista"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestListaVacia(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	require.True(t, lista.EstaVacia())
}

func TestInsertarPrimero(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarPrimero(1)
	require.False(t, lista.EstaVacia())
	require.Equal(t, 1, lista.VerPrimero())
}

func TestInsertarUltimo(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarUltimo(1)
	require.False(t, lista.EstaVacia())
	require.Equal(t, 1, lista.VerPrimero())
}

func TestBorrarPrimero(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarPrimero(1)
	lista.InsertarPrimero(2)
	require.Equal(t, 1, lista.VerUltimo())
	lista.InsertarPrimero(3)
	lista.BorrarPrimero()
	require.Equal(t, 2, lista.VerPrimero())
	lista.BorrarPrimero()
	require.Equal(t, 1, lista.VerPrimero())
	lista.BorrarPrimero()
	require.True(t, lista.EstaVacia())
}
