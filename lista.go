package lista

type Lista[T any] interface {

	// EstaVacia devuelve verdadero si la lista no tiene elementos , false en caso contrario.
	EstaVacia() bool

	// InsertarPrimero inserta el elemento pasado por parámetro al inicio de la lista
	InsertarPrimero(T)

	// InsertarUltima inserta el elemento pasado por parámetro al final de la lista
	InsertarUltimo(T)

	// BorrarPrimero borra el primer elemento de la lista y lo devuelve. Si la lista vacia arrojará panic.
	BorrarPrimero() T

	// VerPrimero retorna el primer elemento de la lista. Si la lista vacia arrojará panic.
	VerPrimero() T

	// VerUltimo retorna el ultimo elemento de la lista. Si la lista vacia arrojará panic.
	VerUltimo() T

	// Lago devuelve la cantidad de elementos de la lista.
	Largo() int

	/*
	  Iterar itera por todos los elementos de la lista. Recibe como parámetro una función que va a recibir el elemento actual de la iteración y devuelve un booleano el cual si es verdadero
	  continúa la iteración y si es falso detiene la iteración.
	*/
	Iterar(visitar func(T) bool)

	// Devuelve una instancia de un Iterador para la lista
	Iterador() IteradorLista[T]
}

type IteradorLista[T any] interface {
	// VerActual retorna el elemento actual del iterador.
	VerActual() T

	// HaySiguiente retorna verdadero si hay mas elementos por iterar, del contrario retorna falso.
	HaySiguiente() bool

	// Siguiente se mueve al elemento que le sigue al actual del iterador.
	Siguiente()

	// Insertar recibe un elemento por parámetro y lo inserta en la posicion del actual del iterador,
	Insertar(T)

	//Borrar borra el elemento actual del iterador y lo devuelve
	Borrar() T
}
