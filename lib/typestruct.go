package lib

import "github.com/google/uuid"

type Token struct {
	ID       uuid.UUID  `json:"id"`
	EMAIL    string     `json:"email"`
	ID_EMP   *uuid.UUID `json:"id_emp"`
	ID_GRUPO *uuid.UUID `json:"id_grupo"`
}

type Carga struct {
	LABEL string `json:"label"`
	VALUE string `json:"value"`
	KEY   string `json:"key"`
	ID    string `json:"id"`
}

type Lookup struct {
	NAME string `json:"name"`
	ID   string `json:"id"`
}
