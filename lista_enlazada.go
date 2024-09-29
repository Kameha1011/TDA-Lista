package lista

type listaEnlazada[T any] struct {
	primero *nodoLista[T]
	ultimo  *nodoLista[T]
	largo   int
}

type nodoLista[T any] struct {
	dato      T
	siguiente *nodoLista[T]
}

func (lista *listaEnlazada[T]) EstaVacia() bool {
	return lista.primero == nil
}

func panicListaVacia[T any](lista *listaEnlazada[T]) {
	if lista.EstaVacia() {
		panic("La lista esta vacia")
	}
}

func (lista *listaEnlazada[T]) InsertarPrimero(valor T) {
	nuevoNodo := crearNuevoNodoLista(valor)
	if lista.EstaVacia() {
		lista.ultimo = nuevoNodo
	}
	nuevoNodo.siguiente = lista.primero
	lista.primero = nuevoNodo
	lista.largo++
}

func (lista *listaEnlazada[T]) InsertarUltimo(valor T) {
	nuevoNodo := crearNuevoNodoLista(valor)
	if lista.EstaVacia() {
		lista.primero = nuevoNodo
	}
	lista.ultimo.siguiente = nuevoNodo
	lista.ultimo = nuevoNodo
	lista.largo++
}

func (lista *listaEnlazada[T]) BorrarPrimero() T {
	panicListaVacia(lista)
	dato := lista.primero.dato
	if lista.largo == 2 {
		lista.primero = lista.ultimo
	} else {
		lista.primero = lista.primero.siguiente
	}
	lista.largo--
	return dato
}

func (lista *listaEnlazada[T]) VerPrimero() T {
	panicListaVacia(lista)
	return lista.primero.dato
}

func (lista *listaEnlazada[T]) VerUltimo() T {
	panicListaVacia(lista)
	return lista.ultimo.dato
}

func (lista *listaEnlazada[T]) Largo() int {
	return lista.largo
}

func (lista listaEnlazada[T]) Iterar(visitar func(T) bool) {
	condicion := true
	nodoActual := lista.primero
	for nodoActual != nil && condicion {
		condicion = visitar(nodoActual.dato)
		nodoActual = nodoActual.siguiente
	}
}

func CrearListaEnlazada[T any]() Lista[T] {
	lista := new(listaEnlazada[T])
	lista.primero = nil
	lista.ultimo = nil
	lista.largo = 0
	return lista
}

func crearNuevoNodoLista[T any](valor T) *nodoLista[T] {
	nuevoNodo := new(nodoLista[T])
	nuevoNodo.dato = valor
	nuevoNodo.siguiente = nil
	return nuevoNodo
}

func (lista *listaEnlazada[T]) Iterador() IteradorLista[T] {
	iterador := crearIteradorListaEnlazada[T](lista)
	return iterador
}

type iteradorListaEnlazada[T any] struct {
	actual   *nodoLista[T]
	anterior *nodoLista[T]
	lista    *listaEnlazada[T]
}

func panicIteradorFinalizado[T any](iterador *iteradorListaEnlazada[T]) {
	if iterador.actual == nil {
		panic("El iterador termino de iterar")
	}
}

func (iterador *iteradorListaEnlazada[T]) VerActual() T {
	panicIteradorFinalizado(iterador)
	return iterador.actual.dato
}

func (iterador *iteradorListaEnlazada[T]) HaySiguiente() bool {
	return iterador.actual != nil
}

func (iterador *iteradorListaEnlazada[T]) Siguiente() {
	panicIteradorFinalizado(iterador)
	iterador.anterior = iterador.actual
	iterador.actual = iterador.actual.siguiente
}

func (iterador *iteradorListaEnlazada[T]) Insertar(valor T) {
	nuevoNodo := crearNuevoNodoLista(valor)
	nuevoNodo.siguiente = iterador.actual

	if iterador.anterior != nil {
		iterador.anterior.siguiente = nuevoNodo
		iterador.lista.primero = nuevoNodo
	}
	if iterador.actual == nil {
		iterador.lista.ultimo = nuevoNodo
	}
	// if iterador.actual == nil && iterador.lista.primero == iterador.actual {
	// 	iterador.lista.primero = nuevoNodo
	// }
	iterador.actual = nuevoNodo
	iterador.lista.largo++
}

func (iterador *iteradorListaEnlazada[T]) Borrar() T {
	panicIteradorFinalizado(iterador)
	valor := iterador.actual.dato
	if iterador.anterior == nil {
		iterador.lista.primero = iterador.actual.siguiente
	} else {
		iterador.anterior.siguiente = iterador.actual.siguiente
	}

	if iterador.actual == iterador.lista.ultimo {
		iterador.lista.ultimo = iterador.anterior
		iterador.actual = iterador.anterior

	} else {
		iterador.actual = iterador.actual.siguiente
	}
	return valor

}

func crearIteradorListaEnlazada[T any](lista *listaEnlazada[T]) *iteradorListaEnlazada[T] {
	iterador := new(iteradorListaEnlazada[T])
	iterador.actual = lista.primero
	iterador.anterior = nil
	return iterador
}
