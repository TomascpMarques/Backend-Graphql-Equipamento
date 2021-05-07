package mongodbhandle

import "github.com/tomascpmarques/PAP/backend/robinservicodocumentacao/resolvedschema"

// ParseTipoDeRegisto -
func ParseTipoDeRegisto(alvo map[string]interface{}) interface{} {
	switch alvo["tipo_de_registo"] {
	// Para o tipo Item
	case "Repositorio":
		return resolvedschema.RepositorioParaStruct(&alvo)
	case "FicheiroMetaData":
		return resolvedschema.FicheiroMetaDataParaStruct(&alvo)
	case "FicheiroConteudo":
		return resolvedschema.FicheiroConteudoParaStruct(&alvo)
	}

	return nil
}

// MongoRecordsParssedArrays :
// 	Utiliza o vetor de maps fornecido para traduzir e formatar os com campos JSON custom para melhor relacionamento
func MongoRecordsParssedArrays(registos []map[string]interface{}) []map[string]interface{} {
	returns := make([]map[string]interface{}, 0)
	// Itera sobre todos os maps
	for _, valor := range registos {
		// Se a converssão for bem sucedida
		if registos := ParseTipoDeRegisto(valor); registos != nil {
			// Adiciona a informação misc (não presente na defenição da struct), e a struct
			returns = append(returns, map[string]interface{}{
				"misc": map[string]interface{}{
					"id":              valor["_id"],
					"tipo_de_registo": valor["tipo_de_registo"],
				},
				"registo": registos,
			})
			continue
		} else {
			returns = append(returns, map[string]interface{}{
				"misc": map[string]interface{}{
					"info_converssao": "Tipo de registo não implementado ou erro de conversão",
				},
				"registo": valor,
			})
		}
	}
	return returns
}