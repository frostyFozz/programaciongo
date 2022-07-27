slices

Son apuntadores a arrays (no almacenan ningún dato), permiten trabajar con arrays de forma dinámica.

Nota:

El ultimo elemento del rango del slice es excluyente.
Modificar elementos del slice, en verdad es modificar elementos del array donde apunta el slice.
Otras sintaxis:

name_array[ : indice_final]: el índice inicial es el 0.
name_array[indice_inicial : ]: el índice final es el tamaño del array.
name_array[ : ]: Toma todos los elementos del array.