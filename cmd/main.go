package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

const FILE_NAME = "tasks.json"

type Task struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}

type TasksList struct {
	tasks []Task
}

// add
func (l *TasksList) addTask(t Task) {
	l.tasks = append(l.tasks, t)
}

// complete task
func (l *TasksList) completeTask(index int) {
	l.tasks[index].Completed = true
}

// update
func (l *TasksList) updateTask(index int, t Task) {
	l.tasks[index] = t
}

// remove
func (l *TasksList) deleteTask(index int) {
	l.tasks = append(l.tasks[:index], l.tasks[index+1:]...)
}

// list all task
func (l *TasksList) listAll() {
	if len(l.tasks) == 0 {
		fmt.Println("Lista vacia.")
		return
	}

	fmt.Println("Lista de tareas:")
	fmt.Println("=======================================")

	for i, t := range l.tasks {
		fmt.Printf("%d. %s - %s - completado: %t\n", i+1, t.Name, t.Description, t.Completed)
	}

	fmt.Println("=======================================")
}

// load tasks from json file
func (l *TasksList) loadFile() error {
	file, err := os.Open(FILE_NAME)
	if os.IsNotExist(err) {
		return nil
	}

	if err != nil {
		return err
	}

	defer file.Close()

	decoder := json.NewDecoder(file)

	err = decoder.Decode(&l.tasks)
	if err != nil {
		return err
	}

	return nil
}

// save / refresh tasks list in json file
func (l *TasksList) saveInFile() error {
	if len(l.tasks) == 0 {
		fmt.Print("Lista vacia. No hay registros que guardar")
	}

	file, err := os.Create(FILE_NAME)
	if err != nil {
		return err
	}

	defer file.Close()

	encoder := json.NewEncoder(file)

	err = encoder.Encode(l.tasks)
	if err != nil {
		return err
	}

	fmt.Println("Guardado exitoso.")

	return nil
}

func main() {
	list := TasksList{}
	reader := bufio.NewReader(os.Stdin) //To read many characters without limitations as fmt.

	err := list.loadFile()
	if err != nil {
		log.Panic(err)
	}

	for {
		var option int

		fmt.Println("Seleccione una opción:\n",
			"1. Agregar tarea\n",
			"2. Marcar tarea como completada\n",
			"3. Editar tarea\n",
			"4. Eliminar tarea\n",
			"5. Listar tareas\n",
			"6. Guardar / actualizar lista de tareas en archivo\n",
			"7. Salir")

		fmt.Print("Ingrese la opción: ")

		_, err := fmt.Scanln(&option)
		if err != nil {
			log.Println(err)
			return
		}

		switch option {
		case 1:
			var t Task

			fmt.Print("Ingrese el nombre de la tarea: ")
			t.Name, _ = reader.ReadString('\n')
			fmt.Print("Ingrese la descripcion de la tarea: ")
			t.Description, _ = reader.ReadString('\n')

			list.addTask(t)

			fmt.Println("Tarea agregada correctamente.")
		case 2:
			var index int

			fmt.Print("Ingrese numero de la tarea que desea completar: ")
			fmt.Scanln(&index)

			list.completeTask(index - 1)

			fmt.Println("Tarea marcada como completada correctamente.")
		case 3:
			var (
				index int
				t     Task
			)

			fmt.Print("Ingrese numero de la tarea que desea editar: ")
			fmt.Scanln(&index)

			fmt.Print("Ingrese el nombre de la tarea: ")
			t.Name, _ = reader.ReadString('\n')
			fmt.Print("Ingrese la descripcion de la tarea: ")
			t.Description, _ = reader.ReadString('\n')

			list.updateTask(index-1, t)

			fmt.Println("Tarea editada correctamente.")
		case 4:
			var index int

			fmt.Print("Ingrese numero de la tarea que desea eliminar: ")
			fmt.Scanln(&index)

			list.deleteTask(index - 1)

			fmt.Println("Tarea eliminada correctamente.")
		case 5:
			list.listAll()
		case 6:
			err := list.saveInFile()
			if err != nil {
				fmt.Println("Hubo un error, intentelo nuevamente: ", err)
			}
		case 7:
			fmt.Println("Saliendo...")
			return
		default:
			fmt.Println("Opción invalida.")
		}
	}
}
