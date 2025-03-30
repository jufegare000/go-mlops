Para el despligue efectivo de la siguiente aplicación en java se listan los siguientes requerimientos:
1. Tener un contenedor con el proyecto con las siguientes características:
    a. Tener instalado golang en su máquina
2. Desplegar una base de datos mysql con las siguientes características:
    a. *usuario*: root
    b. *contraseña*: root
    c. *host*: el-que-kubernetes-entregue.
    d. *Tablas de la base de datos*:
    - **author**:
	    - id: int
	    - name: varchar(128)
	    - birth_date: date
	    - death_date: date
	    - nationality: varchar(64)
	- **literary_work**:
		- id: int
	    - name: varchar(128)
		- publication_date: date
	- **famous_quote**:
		- id: int
		- literary_work_id: int
		- description: varchar(2048)
