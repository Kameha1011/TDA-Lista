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
	lista.InsertarUltimo(2)
	require.False(t, lista.EstaVacia())
	require.Equal(t, 1, lista.VerPrimero())
	lista.InsertarPrimero(3)
	require.Equal(t, 3, lista.VerPrimero())
	require.Equal(t, 2, lista.VerUltimo())
	require.Equal(t, 3, lista.Largo())
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
	lista.InsertarPrimero(3)
	lista.InsertarPrimero(4)
	require.Equal(t, 4, lista.BorrarPrimero())
	lista.InsertarPrimero(5)
	require.Equal(t, 5, lista.BorrarPrimero())
	require.Equal(t, 3, lista.BorrarPrimero())
	require.Equal(t, 2, lista.BorrarPrimero())
	require.Equal(t, 1, lista.BorrarPrimero())
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
	lista := TDALista.CrearListaEnlazada[int]()
	require.True(t, lista.EstaVacia())
	lista.InsertarPrimero(1)
	require.Equal(t, 1, lista.VerPrimero())
	lista.InsertarPrimero(2)
	require.Equal(t, 2, lista.VerPrimero())
	lista.InsertarPrimero(3)
	require.Equal(t, 3, lista.VerPrimero())

}

func TestVerUltimo(t *testing.T) { //omar
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarUltimo(1)
	lista.InsertarUltimo(2)
	require.Equal(t, 2, lista.VerUltimo())
}

func TestLargo(t *testing.T) { //valentin
	lista := TDALista.CrearListaEnlazada[int]()
	require.Equal(t, 0, lista.Largo())
	lista.InsertarPrimero(1)
	require.Equal(t, 1, lista.Largo())
	lista.InsertarPrimero(2)
	require.Equal(t, 2, lista.Largo())
	lista.InsertarPrimero(3)
	require.Equal(t, 3, lista.Largo())
	lista.BorrarPrimero()
	require.Equal(t, 2, lista.Largo())
	lista.BorrarPrimero()
	require.Equal(t, 1, lista.Largo())
	lista.BorrarPrimero()
	require.Equal(t, 0, lista.Largo())
}

func TestVerUltimoPrimero(t *testing.T) { // omar
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarPrimero(1)
	require.Equal(t, 1, lista.VerUltimo())
	require.Equal(t, 1, lista.VerPrimero())
	require.Equal(t, lista.VerUltimo(), lista.VerPrimero())
}

func TestVolumen(t *testing.T) { // valentin
	lista := TDALista.CrearListaEnlazada[int]()
	for i := 0; i < 100000; i++ {
		lista.InsertarPrimero(i)
		require.Equal(t, i, lista.VerPrimero())
		require.Equal(t, 0, lista.VerUltimo())
		require.Equal(t, i+1, lista.Largo())
	}
	for i := 0; i < 100000; i++ {
		require.Equal(t, 99999-i, lista.BorrarPrimero())
		require.Equal(t, 99999-i, lista.Largo())
	}

	require.True(t, lista.EstaVacia())

	for i := 0; i < 100000; i++ {
		lista.InsertarUltimo(i)
		require.Equal(t, 0, lista.VerPrimero())
		require.Equal(t, i, lista.VerUltimo())
		require.Equal(t, i+1, lista.Largo())
	}
	for i := 0; i < 100000; i++ {
		require.Equal(t, i, lista.BorrarPrimero())
		require.Equal(t, 99999-i, lista.Largo())
	}

	require.True(t, lista.EstaVacia())

}

func TestVolumenFloats(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[float64]()
	for i := 0; i < 100000; i++ {
		lista.InsertarPrimero(float64(i))
		require.Equal(t, float64(i), lista.VerPrimero())
		require.Equal(t, 0.0, lista.VerUltimo())
	}
	for i := 0; i < 100000; i++ {
		require.Equal(t, float64(99999-i), lista.BorrarPrimero())
		require.Equal(t, 99999-i, lista.Largo())
	}

	for i := 0; i < 100000; i++ {
		lista.InsertarUltimo(float64(i))
		require.Equal(t, 0.0, lista.VerPrimero())
		require.Equal(t, float64(i), lista.VerUltimo())
	}

	for i := 0; i < 100000; i++ {
		require.Equal(t, float64(i), lista.BorrarPrimero())
		require.Equal(t, 99999-i, lista.Largo())
	}

	require.True(t, lista.EstaVacia())
}

func TestComportamiento(t *testing.T) { // valentin
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarPrimero(1)
	lista.InsertarPrimero(2)
	lista.InsertarPrimero(3)
	lista.InsertarPrimero(4)

	require.Equal(t, 4, lista.VerPrimero())
	require.Equal(t, 1, lista.VerUltimo())
	require.Equal(t, 4, lista.Largo())

	require.Equal(t, 4, lista.BorrarPrimero())
	require.Equal(t, 3, lista.VerPrimero())
	require.Equal(t, 1, lista.VerUltimo())
	require.Equal(t, 3, lista.Largo())

	require.False(t, lista.EstaVacia())

	lista.InsertarUltimo(5)
	require.Equal(t, 3, lista.VerPrimero())
	require.Equal(t, 5, lista.VerUltimo())
	require.Equal(t, 4, lista.Largo())

	lista.InsertarUltimo(6)
	lista.InsertarUltimo(7)
	require.Equal(t, 3, lista.VerPrimero())
	require.Equal(t, 7, lista.VerUltimo())
	require.Equal(t, 6, lista.Largo())

	require.Equal(t, 3, lista.BorrarPrimero())
	require.Equal(t, 2, lista.BorrarPrimero())
	require.Equal(t, 1, lista.BorrarPrimero())
	require.Equal(t, 5, lista.BorrarPrimero())
	require.Equal(t, 6, lista.BorrarPrimero())
	require.Equal(t, 7, lista.BorrarPrimero())

	require.True(t, lista.EstaVacia())
	require.Equal(t, 0, lista.Largo())

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
	var arr []int
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarPrimero(1)
	lista.InsertarPrimero(2)
	lista.InsertarPrimero(3)
	lista.InsertarPrimero(4)
	lista.InsertarPrimero(5)
	lista.Iterar(func(v int) bool {
		arr = append(arr, v)
		return true
	})
	require.Equal(t, []int{5, 4, 3, 2, 1}, arr)
}

func TestIterarInternoCorte(t *testing.T) { // omar
	var arr []int
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarPrimero(1)
	lista.InsertarPrimero(2)
	lista.InsertarPrimero(3)
	lista.InsertarPrimero(4)
	lista.InsertarPrimero(5)
	lista.Iterar(func(v int) bool {
		arr = append(arr, v)
		return v != 3
	})
	require.Equal(t, []int{5, 4, 3}, arr)
}

func TestIterarInternoVacia(t *testing.T) { // valentin
	suma := 0
	lista := TDALista.CrearListaEnlazada[int]()
	lista.Iterar(func(v int) bool {
		suma += v
		return true
	})
	require.Equal(t, 0, suma)
}

func TestIterarInternoVolumen(t *testing.T) { // omar
}

func TestIterarExterno(t *testing.T) { // valentin
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarPrimero(1)
	lista.InsertarPrimero(2)
	lista.InsertarPrimero(3)
	lista.InsertarPrimero(4)
	lista.InsertarPrimero(5)
	suma := 0
	iter := lista.Iterador()
	for iter.HaySiguiente() {
		suma += iter.VerActual()
		iter.Siguiente()
	}
	require.Equal(t, 15, suma)
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Siguiente() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Borrar() })
}

func TestIterarExternoCorte(t *testing.T) { //valentin
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarPrimero(1)
	lista.InsertarPrimero(2)
	lista.InsertarPrimero(3)
	lista.InsertarPrimero(4)
	lista.InsertarPrimero(5)
	suma := 0
	iter := lista.Iterador()
	for iter.HaySiguiente() {
		suma += iter.VerActual()
		if iter.VerActual() == 3 {
			iter.Siguiente()
			break
		}
		iter.Siguiente()
	}
	require.Equal(t, 12, suma)
	suma = 0
	for iter.HaySiguiente() {
		suma += iter.VerActual()
		iter.Siguiente()
	}
	require.Equal(t, 3, suma)
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Siguiente() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Borrar() })

}

func TestIterarExternoVacia(t *testing.T) { // valentin

	sum := 0
	lista := TDALista.CrearListaEnlazada[int]()
	for iter := lista.Iterador(); iter.HaySiguiente(); iter.Siguiente() {
		sum += iter.VerActual()
	}
	require.Equal(t, 0, sum)
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { lista.Iterador().VerActual() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { lista.Iterador().Siguiente() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { lista.Iterador().Borrar() })
}

func TestIteradorExternoInsertarVacia(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	iter := lista.Iterador()
	iter.Insertar(1)
	require.Equal(t, 1, lista.VerPrimero())
	require.Equal(t, 1, lista.VerUltimo())
	require.Equal(t, 1, lista.Largo())

}

func TestIteradorExternoInsertar(t *testing.T) { // valentin
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarPrimero(1)
	lista.InsertarPrimero(2)
	lista.InsertarPrimero(3)
	lista.InsertarPrimero(4)
	lista.InsertarPrimero(5)
	sum := 0
	for iter := lista.Iterador(); iter.HaySiguiente(); iter.Siguiente() {
		sum += iter.VerActual()
	}
	require.Equal(t, 15, sum)
	for iter2 := lista.Iterador(); iter2.HaySiguiente(); iter2.Siguiente() {
		if iter2.VerActual() == 3 {
			iter2.Insertar(6)
			break
		}
	}
	sum2 := 0
	for iter3 := lista.Iterador(); iter3.HaySiguiente(); iter3.Siguiente() {
		sum2 += iter3.VerActual()
	}
	require.Equal(t, 21, sum2)
}

func TestIteradorExternoInsertarFinal(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarUltimo(1)
	lista.InsertarUltimo(2)
	lista.InsertarUltimo(3)
	lista.InsertarUltimo(4)
	sum := 0
	for iter := lista.Iterador(); iter.HaySiguiente(); iter.Siguiente() {
		sum += iter.VerActual()
	}
	require.Equal(t, 10, sum)
	largo := lista.Largo()
	iter := lista.Iterador()
	for i := 0; i <= largo; i++ {
		if i == largo {
			iter.Insertar(5)
		}
		iter.Siguiente()
	}
	sum2 := 0
	for iter2 := lista.Iterador(); iter2.HaySiguiente(); iter2.Siguiente() {
		sum2 += iter2.VerActual()
	}
	require.Equal(t, 15, sum2)
	require.Equal(t, 5, lista.Largo())
	require.Equal(t, 5, lista.VerUltimo())

}

func TestIteradorExternoInsertarInicio(t *testing.T) { // valentin
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarUltimo(1)
	lista.InsertarUltimo(2)
	lista.InsertarUltimo(3)
	require.Equal(t, 1, lista.VerPrimero())
	largo := lista.Largo()
	iter := lista.Iterador()
	for i := 0; i < largo; i++ {
		if i == 0 {
			iter.Insertar(4)
		}
		iter.Siguiente()
	}
	require.Equal(t, 4, lista.VerPrimero())

}

func TestIteradorExternoInsertarVolumen(t *testing.T) { // omar
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarPrimero(1)
	lista.InsertarPrimero(2)
	lista.InsertarPrimero(3)
	for iter := lista.Iterador(); iter.HaySiguiente(); iter.Siguiente() {
		if iter.VerActual() == 3 {
			for i := 0; i < 100000; i++ {
				iter.Insertar(i)
			}
			break
		}
	}
	require.Equal(t, 100003, lista.Largo())
	require.Equal(t, 1, lista.VerUltimo())
	require.Equal(t, 99999, lista.VerPrimero())

	lista2 := TDALista.CrearListaEnlazada[int]()
	lista2.InsertarPrimero(1)
	lista2.InsertarPrimero(2)
	lista2.InsertarPrimero(3)
	iter2 := lista2.Iterador()
	largo := lista2.Largo()
	for i := 0; i <= largo; i++ {
		if i == 3 {
			for j := 0; j < 100000; j++ {
				iter2.Insertar(j)
			}
		}
		iter2.Siguiente()
	}

	require.Equal(t, 100003, lista2.Largo())
	require.Equal(t, 0, lista2.VerUltimo())
	require.Equal(t, 3, lista2.VerPrimero())
}

func TestIteradorExternoBorrar(t *testing.T) { // valentin
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarPrimero(1)
	lista.InsertarPrimero(2)
	lista.InsertarPrimero(3)

	sum := 0
	for iter := lista.Iterador(); iter.HaySiguiente(); iter.Siguiente() {
		sum += iter.VerActual()
	}
	require.Equal(t, 6, sum)

	guardar := 0
	for iter2 := lista.Iterador(); iter2.HaySiguiente(); iter2.Siguiente() {
		if iter2.VerActual() == 3 {
			guardar = iter2.Borrar()
			break
		}
	}
	require.Equal(t, 3, guardar)

	sum2 := 0
	for iter3 := lista.Iterador(); iter3.HaySiguiente(); iter3.Siguiente() {
		sum2 += iter3.VerActual()
	}
	require.Equal(t, 3, sum2)

}

func TestIteradorExternoBorrar1(t *testing.T) { // omar
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarPrimero(1)
	iter := lista.Iterador()
	iter.Borrar()
	require.True(t, lista.EstaVacia())
	iter.Insertar(1)
	require.Equal(t, 1, lista.VerPrimero())
}

func TestIteradorExternoBorrarInicio(t *testing.T) { // valentin
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarUltimo(1)
	lista.InsertarUltimo(2)
	lista.InsertarUltimo(3)
	require.Equal(t, 1, lista.VerPrimero())

	inicioOriginal := lista.VerPrimero()

	iter := lista.Iterador()

	for i := 0; i < lista.Largo(); i++ {
		if i == 0 {
			iter.Borrar()
		}
		iter.Siguiente()
	}

	require.Equal(t, 1, inicioOriginal)

	require.Equal(t, 2, lista.VerPrimero())

	require.Equal(t, 2, lista.Largo())

	sum := 0
	for iter2 := lista.Iterador(); iter2.HaySiguiente(); iter2.Siguiente() {
		sum += iter2.VerActual()
	}
	require.Equal(t, 5, sum)

}

func TestIteradorExternoVaciar(t *testing.T) { // valentin
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarPrimero(1)
	lista.InsertarPrimero(2)
	lista.InsertarPrimero(3)
	lista.InsertarPrimero(4)

	iter := lista.Iterador()
	for iter.HaySiguiente() {
		iter.Borrar()
	}

	require.True(t, lista.EstaVacia())
	require.Equal(t, 0, lista.Largo())
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { lista.Iterador().VerActual() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { lista.Iterador().Siguiente() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { lista.Iterador().Borrar() })

}

func TestIteradorExternoBorrarFinal(t *testing.T) { //Borrar en el ultimo elemento, sin que finalize el iterador
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarUltimo(1)
	lista.InsertarUltimo(2)
	lista.InsertarUltimo(3)
	require.Equal(t, 3, lista.VerUltimo())
	borrado := 0

	iter := lista.Iterador()
	for iter.HaySiguiente() {
		if iter.VerActual() == 3 {
			borrado = iter.Borrar()
		} else {
			iter.Siguiente()
		}
	}

	require.Equal(t, 3, borrado)
	require.Equal(t, 2, lista.VerUltimo())
	require.Equal(t, 2, lista.Largo())

}

func TestIteradorExternoBorrarVolumen(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarPrimero(0)
	iter := lista.Iterador()

	for i := 0; i < 100000; i++ {
		iter.Insertar(i)
	}

	for iter.HaySiguiente() {
		iter.Borrar()
	}

	require.Equal(t, 0, lista.Largo())
	require.True(t, lista.EstaVacia())
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Siguiente() })

}
