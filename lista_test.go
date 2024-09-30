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

func TestVaciar(t *testing.T) {
}

func TestVerPrimero(t *testing.T) {
}

func TestVerUltimo(t *testing.T) {
}

func TestLargo(t *testing.T) {
}

func TestVerUltimoPrimero(t *testing.T) {
}

func TestVolumen(t *testing.T) {
}

func TestVolumenFloats(t *testing.T) {
}

func TestComportamiento(t *testing.T) {
}

func TestComportamientoFloats(t *testing.T) {
}

func TestIterarInterno(t *testing.T) {
}

func TestIterarInternoCorte(t *testing.T) {
}

func TestIterarInternoVacia(t *testing.T) {
}

func TestIterarInternoVolumen(t *testing.T) {
}

func TestIterarExterno(t *testing.T) {
}

func TestIterarExternoCorte(t *testing.T) {
}

func TestIterarExternoVacia(t *testing.T) {

}

func TestIteradorExternoInsertarVacia(t *testing.T) {

}

func TestIteradorExternoInsertar(t *testing.T) {

}

func TestIteradorExternoInsertarVolumen(t *testing.T) {

}

func TestIteradorExternoBorrar(t *testing.T) {

}

func TestIteradorExternoInsertarFinal(t *testing.T) {

}

func TestIteradorExternoInsertarInicio(t *testing.T) {

}

func TestIteradorExternoInsertarMedio(t *testing.T) {

}

func TestIteradorExternoBorrarInicio(t *testing.T) {

}

func TestIteradorExternoBorrarMedio(t *testing.T) {

}

func TestIteradorExternoBorrarFinal(t *testing.T) {

}

func TestIteradorExternoBorrarVolumen(t *testing.T) {

}

func TestIteradorExternoVaciar(t *testing.T) {

}
