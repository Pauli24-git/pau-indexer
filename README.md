**pau-indexer**

**Reto de aprendizaje - CHALLENGE**

Este repositorio tiene 3 partes principales: INDEXER, API y VISUALIZADOR(front)

**INDEXER**
Este programa recibe como parametro un directorio donde estan alojados los archivos que se van a procesar
Se debe tener una instancia de ZincSearch corriendo, en el archivo .env van los datos de login de esa instancia
Una vez procesado los archivos, se hace una sola llamada a la API enviando todos los documentos para su indexación.

**API**
Esta compuesta por 1 solo endpoint que recibe el término a buscar y opcionalmente definir si ese término se busca en algún campo en 
especifico

**VISUALIZADOR**
Realizado en Vue3 aplicando Tailwind para el estilo de la página
Se debe ingresar el término a buscar en la barra de busqueda y opcionalmente se puede escoger el campo en el dropdown de la derecha