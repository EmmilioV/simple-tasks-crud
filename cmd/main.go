package main

import (
	"bufio"
	"fmt"
	"os"
)

type Task struct {
	name        string
	description string
	completed   bool
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
	l.tasks[index].completed = true
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
		fmt.Printf("%d. %s - %s - completado: %t\n", i+1, t.name, t.description, t.completed)
	}

	fmt.Println("=======================================")
}

func main() {
	list := TasksList{}
	reader := bufio.NewReader(os.Stdin) //Para leer varios caracteres sin tantas limitaciones como fmt.

	for {
		var option int

		fmt.Println("Seleccione una opción:\n",
			"1. Agregar tarea\n",
			"2. Marcar tarea como completada\n",
			"3. Editar tarea\n",
			"4. Eliminar tarea\n",
			"5. listar tareas\n",
			"6. Salir")

		fmt.Print("Ingrese la opción: ")
		fmt.Scanln(&option)

		switch option {
		case 1:
			var t Task

			fmt.Print("Ingrese el nombre de la tarea: ")
			t.name, _ = reader.ReadString('\n')
			fmt.Print("Ingrese la descripcion de la tarea: ")
			t.description, _ = reader.ReadString('\n')

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
			t.name, _ = reader.ReadString('\n')
			fmt.Print("Ingrese la descripcion de la tarea: ")
			t.description, _ = reader.ReadString('\n')

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
			fmt.Println("Saliendo...")
			return
		default:
			fmt.Println("Opción invalida.")
		}
	}
}
