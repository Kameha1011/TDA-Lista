package lista

const (
	_LARGO_INICIAL = 0
)

type nodoLista[T any] struct {
	dato      T
	siguiente *nodoLista[T]
}

func crearNuevoNodoLista[T any](valor T) *nodoLista[T] {
	nuevoNodo := new(nodoLista[T])
	nuevoNodo.dato = valor
	nuevoNodo.siguiente = nil
	return nuevoNodo
}

type listaEnlazada[T any] struct {
	primero *nodoLista[T]
	ultimo  *nodoLista[T]
	largo   int
}

func panicListaVacia[T any](lista *listaEnlazada[T]) {
	if lista.EstaVacia() {
		panic("La lista esta vacia")
	}
}

func CrearListaEnlazada[T any]() Lista[T] {
	lista := new(listaEnlazada[T])
	lista.primero = nil
	lista.ultimo = nil
	lista.largo = _LARGO_INICIAL
	return lista
}

func (lista *listaEnlazada[T]) EstaVacia() bool {
	return lista.primero == nil
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
	} else {
		lista.ultimo.siguiente = nuevoNodo
	}

	lista.ultimo = nuevoNodo
	lista.largo++
}

func (lista *listaEnlazada[T]) BorrarPrimero() T {
	panicListaVacia(lista)
	dato := lista.primero.dato

	if lista.primero.siguiente == lista.ultimo {
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

func (lista *listaEnlazada[T]) Iterador() IteradorLista[T] {
	iterador := crearIteradorListaEnlazada[T](lista)
	return iterador
}

type iteradorListaEnlazada[T any] struct {
	actual   *nodoLista[T]
	anterior *nodoLista[T]
	lista    *listaEnlazada[T]
}

func crearIteradorListaEnlazada[T any](lista *listaEnlazada[T]) *iteradorListaEnlazada[T] {
	iterador := new(iteradorListaEnlazada[T])
	iterador.actual = lista.primero
	iterador.anterior = nil
	iterador.lista = lista
	return iterador
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

	if iterador.anterior == nil {
		iterador.lista.primero = nuevoNodo
	} else {
		iterador.anterior.siguiente = nuevoNodo
	}

	if iterador.actual == nil {
		iterador.lista.ultimo = nuevoNodo
	}

	iterador.actual = nuevoNodo
	iterador.lista.largo++
}

func (iterador *iteradorListaEnlazada[T]) Borrar() T {
	panicIteradorFinalizado(iterador)
	valor := iterador.actual.dato

	if iterador.actual == iterador.lista.primero {
		iterador.lista.primero = iterador.actual.siguiente
	} else {
		iterador.anterior.siguiente = iterador.actual.siguiente
	}

	if iterador.actual == iterador.lista.ultimo {
		iterador.lista.ultimo = iterador.anterior
	}

	iterador.actual = iterador.actual.siguiente
	iterador.lista.largo--
	return valor

}
