package parser

type KeywordType = int

const (
	UNKNOWN_KEYWORD KeywordType = iota
	GLOBAL_KEYWORD
	LET_KEYWORD
	CONST_KEYWORD
	IF_KEYWORD
	ELSE_KEYWORD
	ELSEIF_KEYWORD
	SWITCH_KEYWORD
	LOOP_KEYWORD
	BREAK_KEYWORD
	CONTINUE_KEYWORD
	RETURN_KEYWORD
	OBJECT_KEYWORD
	IMPORT_KEYWORD
	FUNCTION_KEYWORD
	NAMESPACE_KEYWORD
	USING_KEYWORD
	CLASS_KEYWORD
	MODEL_CLASS_KEYWORD
	DB_CLASS_KEYWORD
	TABLE_CLASS_KEYWORD
	PRIMARY_KEY_KEYWORD
	FOREIGN_KEY_KEYWORD
	BLUEPRINT_KEYWORD
	ENUM_KEYWORD
	PUBLIC_KEYWORD
	PRIVATE_KEYWORD
	PROTECTED_KEYWORD
	ABSTRACT_KEYWORD
	LATE_KEYWORD
	VIRTUAL_KEYWORD
	INTERNAL_KEYWORD
	REF_KEYWORD
	EMBED_KEYWORD
	EXTENSION_KEYWORD
	OBSERVE_KEYWORD
	UI_KEYWORD
)
