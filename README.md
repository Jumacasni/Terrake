# Terrake

**Autor:** Juan Manuel Castillo Nievas

## Descripción del proyecto

La descripción de este proyecto se puede consultar en [este enlace](https://github.com/Jumacasni/Terrake/blob/main/docs/descripcion_proyecto.md).

## Arquitectura

Al principio podría pensarse una arquitectura monolítica, de forma que agruparíamos toda la funcionalidad en una misma base de código, pero una vez que se han estudiado las distintas funcionalidades que integra el producto completo, se ha visto que pueden diferenciarse claramente cuatro microservicios y que cada uno de ellos realiza una tarea distinta:
* **Catálogo de terremotos**
* **Cuentas de usuario**
* **Reporte de terremotos**
* **Envío de notificaciones**

Cada microservicio realiza una funcionalidad que no depende de otro pero que uniendo todos estos resulta en el producto total, con lo cual se ha elegido finalmente una **arquitectura basada en microservicios**. Como desarrollador de este proyecto, esta arquitectura me permite una mayor facilidad en el avance de cada microservicio, pues cada uno tiene su propio código y esto hace que resulten más sencillos y ordenados los cambios que puedan surgir en cada uno. Lo que más me beneficia de esta arquitectura es que el software va a ser fácilmente personalizable y escalable cuando se despliegue en la nube, al contrario que ocurriría si se hubiera optado por la arquitectura monolítica.

### Herramientas

* **Lenguaje:** en un principio se pensó *Node.js* por la experiencia que tengo en este lenguaje, pero finalmente se va a realizar en **Go** ya que nunca he tenido la oportunidad de aprenderlo y este proyecto es el momento para hacerlo.

## Planificación del proyecto

El desarrollo de este proyecto se va a dividir en las siguientes fases:

### *Milestone* [**catálogo de terremotos**](https://github.com/Jumacasni/Terrake/milestone/3): primera fase

En esta primera fase se va a desarrollar este *milestone* que da como resultado el producto mínimo viable en el que los usuarios pueden realizar consultas de terremotos. La historia de usuario que cubre este *milestone* es:

* [[HU1] Consultar terremotos](https://github.com/Jumacasni/Terrake/issues/46)

### *Milestone* [**cuentas de usuario**](https://github.com/Jumacasni/Terrake/milestone/4): segunda fase

Este *milestone* da como resultado final el producto mínimo viable que incluye la gestión de las cuentas de usuario. Los usuarios podrán registrarse en el sistema, además de poder modificar su información o eliminar su cuenta. Esta fase incluye las siguientes historias de usuario: 

* [[HU3] Registrar usuario](https://github.com/Jumacasni/Terrake/issues/48)
* [[HU4] Modificar cuenta de usuario](https://github.com/Jumacasni/Terrake/issues/49)
* [[HU5] Eliminar cuenta de usuario](https://github.com/Jumacasni/Terrake/issues/50)

### *Milestone* [**notificaciones**](https://github.com/Jumacasni/Terrake/milestone/6): tercera fase

Una vez que se pueden consultar terremotos y los usuarios pueden crear sus cuentas, el siguiente producto mínimo viable incluiría la posibilidad de que estos usuarios registrados puedan recibir notificaciones cada vez que ocurra un terremoto. Los usuarios pueden recibir notificaciones de cada terremoto que ocurra o bien pueden restringir estas notificaciones de forma que sólo reciban aquellas que cumplan con los criterios seleccionados. Por ejemplo: *recibir notificaciones de los terremotos cuya magnitud sea mayor a 2.5 y hayan ocurrido en Andalucía*. Esta fase incluye las siguientes historias de usuario:

* [[HU6] Activar notificaciones](https://github.com/Jumacasni/Terrake/issues/51)
* [[HU7] Restringir las notificaciones](https://github.com/Jumacasni/Terrake/issues/52)
* [[HU8] Desactivar las notificaciones](https://github.com/Jumacasni/Terrake/issues/53)

### *Milestone* [**reporte de terremotos**](https://github.com/Jumacasni/Terrake/milestone/5): cuarta fase

En esta última fase de desarrollo se creará el producto final que incluirá la posibilidad de reportar que se ha sentido un terremoto. Los usuarios podrán indicar dónde se encuentran en este momento y la intensidad con que han sentido el terremoto con el fin de ayudar a hacer mapas de distribución de la intensidad de un terremoto. La historia de usuario que abarca esta última fase es:

* [[HU2] Reportar un terremoto](https://github.com/Jumacasni/Terrake/issues/47)

## Clases creadas

Se han creado las siguientes clases:
- [terremoto.go](https://github.com/Jumacasni/Terrake/blob/main/src/terremoto/terremoto.go): se crea la clase **terremoto**, enlazada a [[HU1]](https://github.com/Jumacasni/Terrake/issues/29)
- [tipomagnitud.go](https://github.com/Jumacasni/Terrake/blob/main/src/terremoto/tipomagnitud/tipomagnitud.go): contiene el *enum* con los tipos de magnitudes que puede tener un terremoto, creado junto con la clase **terremoto** en [[HU1]](https://github.com/Jumacasni/Terrake/issues/29)
- [controlador.go](https://github.com/Jumacasni/Terrake/blob/main/src/controlador.go): se crea la clase **controlador** para crear o destruir terremotos, enlazada a [[HU1]](https://github.com/Jumacasni/Terrake/issues/29)

### Comprobación de la sintaxis

Para la comprobación de que una clase está sintácticamente correcta se ha utilizado el comando ``gofmt -e <fichero.go>``, cuya salida indica los errores sintácticos encontrados.

## Licencia

El código de este repositorio está bajo la licencia [GPLv3](./LICENSE).
