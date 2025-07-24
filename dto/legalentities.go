package dto

type LegalEntityDTO struct {
    UUID    string `json:"uuid"`
    Name    string `json:"name"`
    INN     string `json:"inn"`
    KPP     string `json:"kpp"`
    OGRN    string `json:"ogrn"`
    Address string `json:"address"`
}
