package lista

type listaEnlazada[T any] struct {
}

type nodo[T any] struct {
}

func (lista *listaEnlazada[T]) EstaVacia() bool {
	return false
}

func (lista *listaEnlazada[T]) InsertarPrimero(valor T) {
}

func (lista *listaEnlazada[T]) InsertarUltimo(valor T) {
}

func (lista *listaEnlazada[T]) BorrarPrimero() T {
	return nil
}

func (lista *listaEnlazada[T]) VerPrimero() T {
	return nil
}

func (lista *listaEnlazada[T]) VerUltimo() T {
	return nil
}

func (lista *listaEnlazada[T]) Largo() int {
	return 0
}

func (lista *listaEnlazada[T]) Iterar(visitar func(T) bool) {
}

func (lista *listaEnlazada[T]) Iterador() IteradorLista[T] {
	return nil
}

func CrearListaEnlazada[T any]() Lista[T] {
	return nil
}

func crearNuevoNodo[T any](valor T) *nodo[T] {
	return nil
}

type iteradorListaEnlazada[T any] struct {
}

func (iterador *iteradorListaEnlazada[T]) VerActual() T {
	return nil
}

func (iterador *iteradorListaEnlazada[T]) HaySiguiente() bool {
	return false
}

func (iterador *iteradorListaEnlazada[T]) Siguiente() {
}

func (iterador *iteradorListaEnlazada[T]) Insertar(valor T) {
}

func (iterador *iteradorListaEnlazada[T]) Borrar() T {
	return nil
}

func CrearIteradorListaEnlazada[T any]() IteradorLista[T] {
	return nil
}
