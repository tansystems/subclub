package panel

import (
	"fmt"
	"net/http"
)

// Модуль панели автора: web-интерфейс для управления подписками и контентом.

// AuthorPanelHandler — заглушка панели автора
func AuthorPanelHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Панель автора: здесь будет управление подписчиками и контентом.")
}
