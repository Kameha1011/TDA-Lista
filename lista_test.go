package lista_test

import (
	TDALista "tdas/lista"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestListaVacia(t *testing.T) { //omar
	lista := TDALista.CrearListaEnlazada[int]()
	require.True(t, lista.EstaVacia())
	require.Panics(t, func() { lista.BorrarPrimero() })
	require.Panics(t, func() { lista.VerPrimero() })
	require.Panics(t, func() { lista.VerUltimo() })
	require.Equal(t, 0, lista.Largo())
}

func TestInsertarPrimero(t *testing.T) { //valentin
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarPrimero(1)
	require.False(t, lista.EstaVacia())
	require.Equal(t, 1, lista.VerPrimero())
}

func TestInsertarUltimo(t *testing.T) { //omar
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarUltimo(1)
	require.False(t, lista.EstaVacia())
	require.Equal(t, 1, lista.VerPrimero())
}

func TestBorrarPrimero(t *testing.T) { //valentin
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

func TestVaciar(t *testing.T) { //omar
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarPrimero('h')
	lista.InsertarPrimero('o')
	lista.InsertarPrimero('l')
	lista.InsertarPrimero('a')
	lista.BorrarPrimero()
	lista.BorrarPrimero()
	lista.BorrarPrimero()
	lista.BorrarPrimero()
	require.True(t, lista.EstaVacia())
	require.Panics(t, func() { lista.BorrarPrimero() })
	require.Panics(t, func() { lista.VerPrimero() })
	require.Panics(t, func() { lista.VerUltimo() })
	require.Equal(t, 0, lista.Largo())
}

func TestVerPrimero(t *testing.T) { //valentin
}

func TestVerUltimo(t *testing.T) { //omar
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarUltimo(1)
	lista.InsertarUltimo(2)
	require.Equal(t, 2, lista.VerUltimo())
}

func TestLargo(t *testing.T) { //valentin
}

func TestVerUltimoPrimero(t *testing.T) { // omar
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarPrimero(1)
	require.Equal(t, 1, lista.VerUltimo())
	require.Equal(t, 1, lista.VerPrimero())
	require.Equal(t, lista.VerUltimo(), lista.VerPrimero())
}

func TestVolumen(t *testing.T) { // valentin
}

func TestVolumenFloats(t *testing.T) { // omar
}

func TestComportamiento(t *testing.T) { // valentin
}

func TestComportamientoFloats(t *testing.T) { // omar
	lista := TDALista.CrearListaEnlazada[float64]()
	// comportamiento nuevo
	require.True(t, lista.EstaVacia())
	require.Panics(t, func() { lista.BorrarPrimero() })
	require.Panics(t, func() { lista.VerPrimero() })
	require.Panics(t, func() { lista.VerUltimo() })
	require.Equal(t, 0, lista.Largo())
	//comportamiento lleno
	lista.InsertarPrimero(1)
	require.Equal(t, float64(1), lista.BorrarPrimero())
	lista.InsertarUltimo(1)
	require.Equal(t, 1, lista.Largo())
	lista.InsertarUltimo(2)
	require.Equal(t, float64(2), lista.VerUltimo())
	require.Equal(t, float64(1), lista.VerPrimero())
	// comportamiento vaciando
	lista.BorrarPrimero()
	lista.BorrarPrimero()
	require.True(t, lista.EstaVacia())
	require.Panics(t, func() { lista.BorrarPrimero() })
	require.Panics(t, func() { lista.VerPrimero() })
	require.Panics(t, func() { lista.VerUltimo() })
	require.Equal(t, 0, lista.Largo())
}

func TestIterarInterno(t *testing.T) { // valentin
}

func TestIterarInternoCorte(t *testing.T) { // omar
}

func TestIterarInternoVacia(t *testing.T) { // valentin
}

func TestIterarInternoVolumen(t *testing.T) { // omar
}

func TestIterarExterno(t *testing.T) { // valentin
}

func TestIterarExternoCorte(t *testing.T) { // omar
}

func TestIterarExternoVacia(t *testing.T) { // valentin

}

func TestIteradorExternoInsertarVacia(t *testing.T) { // omar

}

func TestIteradorExternoInsertar(t *testing.T) { // valentin

}

func TestIteradorExternoInsertarVolumen(t *testing.T) { // omar

}

func TestIteradorExternoBorrar(t *testing.T) { // valentin

}

func TestIteradorExternoInsertarFinal(t *testing.T) { // omar

}

func TestIteradorExternoInsertarInicio(t *testing.T) { // valentin

}

func TestIteradorExternoInsertarMedio(t *testing.T) { // omar

}

func TestIteradorExternoBorrarInicio(t *testing.T) { // valentin

}

func TestIteradorExternoBorrarMedio(t *testing.T) { // omar

}

func TestIteradorExternoBorrarFinal(t *testing.T) { // valentin

}

func TestIteradorExternoBorrarVolumen(t *testing.T) { // omar

}

func TestIteradorExternoVaciar(t *testing.T) { // valentin

}
