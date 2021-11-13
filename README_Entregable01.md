Endpoints: 

localhost:8000/pokemons => Se mostrara la informacion completa del csv que se esta utilizando. 
localhost:8000/pokemon/{id} => se mostrara solo el pokemon en cuestion, o un erro saldra. 



NOTA. Se tomo como ejemplo entrada de inforamcion de pokemones (un pokedex, numero y Nombre pokemon) para realizar la logica de negocios de este entregable. 


## First Deliverable (due November 12th 23:59PM)

Based on the self-study material and mentorship covered until this deliverable, we suggest you perform the following:

- Create an API
==> Se realizo las bases de la API (Router, Controller, Models, Servicios)
- Add an endpoint to read from a CSV file
====> No realice endpoint, pero genere un servicio que se encarga de leer un csv y transformarlo a un array de estructura. 
- The CSV should have any information, for example:

```txt
1,bulbasaur
2,ivysaur
3,venusaur
```

- The items in the CSV must have an ID element (int value)
==> El servicio de ReadCSV se encarga de analizar que se obtenga un ID (entero en la primera columna)

- The endpoint should get information from the CSV by some field ***(example: ID)***
==> Se realizo un endpoint donde envias un ID y lo busca. 

- The result should be displayed as a response
==> El resultado de los endpoints (De toda la informacion y una informacion en especifico) son en formato JSON. 

- Clean architecture proposal
==> Se tomo como base algunas caracteristicas de esta architectura (domin, infrastructure, interface, etc)

- Use best practices
- Handle the Errors ***(CSV not valid, error connection, etc)***
==> Se manejaron errores especificos como son: Parseo incorrecto, Invalidez a la hora de abrir archivo, archivo CSV no tiene ID .


TODO 

```Todo
1. Generar  Endpoint exclusivamente para procesar un archivo CSV que se envia mediante POST. 
2. Evitar repetir el codigo para el response del HTTP. 
3. Hacer un Middleware para majear multiples errores. (Customizados )
4. Investigar redundancia a al hora de importar el mismo paquete en archivos diferentes. 
```


Test
```Test
Run Server: go run main.go 

Test Endpoint pokemons: 
- curl -v 'http://localhost:8000/pokemons'

Test endpoint pokemon/{id}
- curl -v 'http://localhost:8000/pokemon/122'
```
