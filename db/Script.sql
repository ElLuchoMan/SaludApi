CREATE SCHEMA acceso_historia;
CREATE SCHEMA odontologia;
CREATE SCHEMA medicina;
CREATE SCHEMA psicologia;

/* Drop Tables */

DROP TABLE IF EXISTS acceso_historia.AccesoHistoriaClinica CASCADE
;

DROP TABLE IF EXISTS odontologia.Anamnesis CASCADE
;

DROP TABLE IF EXISTS medicina.Antecedente CASCADE
;

DROP TABLE IF EXISTS psicologia.AntecedentePsicologico CASCADE
;

DROP TABLE IF EXISTS psicologia.ComportamientoConsulta CASCADE
;

DROP TABLE IF EXISTS psicologia.ComposicionFamiliar CASCADE
;

DROP TABLE IF EXISTS medicina.ConsultaFisioterapia CASCADE
;

DROP TABLE IF EXISTS odontologia.ControlPlaca CASCADE
;

DROP TABLE IF EXISTS medicina.DiagnosticoMedicina CASCADE
;

DROP TABLE IF EXISTS odontologia.DiagnosticoOdontologia CASCADE
;

DROP TABLE IF EXISTS psicologia.DiagnosticoPsicologia CASCADE
;

DROP TABLE IF EXISTS medicina.Examen CASCADE
;

DROP TABLE IF EXISTS odontologia.ExamenDental CASCADE
;

DROP TABLE IF EXISTS odontologia.ExamenesComplementarios CASCADE
;

DROP TABLE IF EXISTS odontologia.ExamenEstomatologico CASCADE
;

DROP TABLE IF EXISTS medicina.HistoriaClinica CASCADE
;

DROP TABLE IF EXISTS medicina.HojaHistoria CASCADE
;

DROP TABLE IF EXISTS psicologia.Limites CASCADE
;

DROP TABLE IF EXISTS odontologia.Odontograma CASCADE
;

DROP TABLE IF EXISTS medicina.Sistemas CASCADE
;

DROP TABLE IF EXISTS medicina.TipoAntecedente CASCADE
;

DROP TABLE IF EXISTS psicologia.TipoAntecedentePsicologico CASCADE
;

DROP TABLE IF EXISTS medicina.TipoExamen CASCADE
;

DROP TABLE IF EXISTS odontologia.TipoOdontograma CASCADE
;

DROP TABLE IF EXISTS psicologia.ValoracionInterpersonal CASCADE
;

/* Create Tables */

CREATE TABLE acceso_historia.AccesoHistoriaClinica
(
	id_acceso serial NOT NULL,
	fecha_acceso timestamp without time zone NULL,
	profesional_id integer NULL,
	id_historia_clinica integer NULL
)
;

CREATE TABLE odontologia.Anamnesis
(
	id_anamnesis serial NOT NULL,
	id_historia_clinica integer NULL,
	tratamiento text NULL,
	medicamentos text NULL,
	alergias text NULL,
	hemorragias text NULL,
	irradiaciones text NULL,
	sinusitis text NULL,
	enfermedad_respiratoria text NULL,
	cardiopatias text NULL,
	diabetes text NULL,
	fiebre_reumatica text NULL,
	hepatitis text NULL,
	hipertension text NULL,
	antecedente_familiar text NULL,
	cepillado integer NULL,
	ceda integer NULL,
	enjuague integer NULL,
	dulces text NULL,
	fuma text NULL,
	chicle text NULL,
	otras text NULL,
	ultima_visita date NULL
)
;

CREATE TABLE medicina.Antecedente
(
	id_antecedente serial NOT NULL,
	id_tipo_antecedente integer NULL,
	observaciones text NULL,
	id_historia_clinica integer NULL
)
;

CREATE TABLE psicologia.AntecedentePsicologico
(
	id_antecedente serial NOT NULL,
	actual_somatico text NULL,
	pasado_somatico text NULL,
	id_tipo_antecedente_psicologico integer NULL,
	id_historia_clinica integer NULL
)
;

CREATE TABLE psicologia.ComportamientoConsulta
(
	id_comportamiento_consulta serial NOT NULL,
	id_hoja_historia integer NULL,
	problematica text NULL,
	afrontamiento text NULL,
	comportamiento text NULL
)
;

CREATE TABLE psicologia.ComposicionFamiliar
(
	id_composicion_familiar serial NOT NULL,
	id_hoja_historia integer NULL,
	observaciones text NULL
)
;

CREATE TABLE medicina.ConsultaFisioterapia
(
	id_consulta_fisioterapia serial NOT NULL,
	id_hoja_historia integer NULL,
	motivo_consulta text NULL,
	valoracion text NULL,
	plan_manejo text NULL,
	evolucion text NULL,
	observaciones text NULL
)
;

CREATE TABLE odontologia.ControlPlaca
(
	id_control_placa serial NOT NULL,
	indice_anterior integer NULL,
	indice_actual integer NULL,
	fecha date NULL,
	id_hoja_historia integer NULL,
	vestibulares text NULL,
	observaciones text NULL,
	id_tipo_odontograma integer NULL
)
;

CREATE TABLE medicina.DiagnosticoMedicina
(
	id_diagnostico serial NOT NULL,
	nombre text NULL,
	descripcion text NULL,
	numero integer NULL,
	activo boolean NULL,
	fecha_creacion timestamp without time zone NULL,
	fecha_modificacion timestamp without time zone NULL,
	plan_de_manejo text NULL,
	analisis text NULL
)
;

CREATE TABLE odontologia.DiagnosticoOdontologia
(
	id_diagnostico serial NOT NULL,
	dianostico text NULL,
	pronostico text NULL,
	evolucion text NULL,
	observaciones text NULL,
	id_hoja_historia integer NULL
)
;

CREATE TABLE psicologia.DiagnosticoPsicologia
(
	id_diagnostico_psicologia serial NOT NULL,
	id_hoja_historia integer NULL,
	hipotesis text NULL,
	acuerdo text NULL,
	observaciones text NULL,
	evolucion text NULL
)
;

CREATE TABLE medicina.Examen
(
	id_examen serial NOT NULL,
	id_hoja_historia integer NULL,
	nombre text NULL,
	observacion text NULL,
	id_tipo_examen integer NULL,
	fecha_examen date NULL
)
;

CREATE TABLE odontologia.ExamenDental
(
	id_examen_dental serial NOT NULL,
	id_hoja_historia integer NULL,
	supernumerarios text NULL,
	abrasion text NULL,
	manchas text NULL,
	patologia_pulpar text NULL,
	placa_blanda text NULL,
	placa_calcificada text NULL,
	oclusion text NULL,
	otros text NULL,
	observaciones text NULL
)
;

CREATE TABLE odontologia.ExamenesComplementarios
(
	id_examenes_complementarios serial NOT NULL,
	periapical_inicio bytea NULL,	-- archivo
	periapical_final bytea NULL,
	panoramica_inicio bytea NULL,
	panoramica_final bytea NULL,
	otra_inicio bytea NULL,
	otra_final bytea NULL,
	laboratorio_inicio bytea NULL,
	laboratorio_final bytea NULL,
	id_hoja_historia integer NULL,
	tp text NULL,
	tpt text NULL,
	coagulacion text NULL,
	sangria text NULL,
	otra text NULL
)
;

CREATE TABLE odontologia.ExamenEstomatologico
(
	id_examen_estomatologico serial NOT NULL,
	observaciones text NULL,
	id_hoja_historia integer NULL,
	articulacion_temporo text NULL,
	labios text NULL,
	lengua text NULL,
	paladar text NULL,
	piso_boca text NULL,
	carrillos text NULL,
	glandulas_salivares text NULL,
	maxilares text NULL,
	senos_maxilares text NULL,
	musculos_masticadores text NULL,
	sistema_nervioso text NULL,
	sistema_vascular text NULL,
	sitema_linfatico text NULL,
	sistema_regional text NULL
)
;

CREATE TABLE medicina.HistoriaClinica
(
	id_historia_clinica serial NOT NULL,
	id_tercero integer NULL
)
;

CREATE TABLE medicina.HojaHistoria
(
	id_hoja_historia serial NOT NULL,
	fecha_consulta timestamp without time zone NULL,
	motivo text NULL,
	id_diagnostico integer NULL,
	observacion text NULL,
	evolucion text NULL,
	id_especialidad integer NULL,
	id_profesional integer NULL,
	id_historia_clinica integer NULL,
	id_persona integer NULL
)
;

CREATE TABLE psicologia.Limites
(
	id_limite serial NOT NULL,
	difusos text NULL,
	claros text NULL,
	rigidos text NULL,
	id_hoja_historia integer NULL
)
;

CREATE TABLE odontologia.Odontograma
(
	id_odontograma serial NOT NULL,
	id_hoja_historia integer NULL,
	diagrama text NULL,
	observaciones varchar(50) NULL,
	id_tipo_odontograma integer NULL
)
;

CREATE TABLE medicina.Sistemas
(
	id_sistema serial NOT NULL,
	nombre_sistema text NULL,
	observacion text NULL,
	id_hoja_historia integer NULL
)
;

CREATE TABLE medicina.TipoAntecedente
(
	id_tipo_antecedente serial NOT NULL,
	nombre text NULL,
	descripcion text NULL,
	activo boolean NULL,
	fecha_creacion timestamp without time zone NULL,
	fecha_modificacion timestamp without time zone NULL
)
;

CREATE TABLE psicologia.TipoAntecedentePsicologico
(
	id_tipo_antecedente_psicologico serial NOT NULL,
	nombre text NULL,
	descripcion text NULL,
	activo boolean NULL,
	fecha_creacion timestamp without time zone NULL,
	fecha_modificacion timestamp without time zone NULL
)
;

CREATE TABLE medicina.TipoExamen
(
	id_tipo_examen serial NOT NULL,
	nombre text NULL
)
;

CREATE TABLE odontologia.TipoOdontograma
(
	id_tipo_odontograma serial NOT NULL,
	nombre text NULL,
	descripcion text NULL,
	fecha_creacion text NULL,
	fecha_modificacion text NULL
)
;

CREATE TABLE psicologia.ValoracionInterpersonal
(
	id_valoracion_interpersonal serial NOT NULL,
	id_hoja_historia integer NULL,
	autoridad text NULL,
	pares text NULL,
	pareja text NULL,
	relaciones boolean NULL,
	satisfaccion text NULL,
	proteccion text NULL,
	orientacion text NULL,
	judiciales text NULL,
	economicos text NULL,
	drogas text NULL,
	motivo text NULL
)
;

/* Create Primary Keys, Indexes, Uniques, Checks */

ALTER TABLE acceso_historia.AccesoHistoriaClinica ADD CONSTRAINT PK_AccesoHistoriaClinica
	PRIMARY KEY (id_acceso)
;

CREATE INDEX IXFK_AccesoHistoriaClinica_HistoriaClinica ON acceso_historia.AccesoHistoriaClinica (id_historia_clinica ASC)
;

ALTER TABLE odontologia.Anamnesis ADD CONSTRAINT PK_Anamnesis
	PRIMARY KEY (id_anamnesis)
;

CREATE INDEX IXFK_Anamnesis_HistoriaClinica ON odontologia.Anamnesis (id_historia_clinica ASC)
;

ALTER TABLE medicina.Antecedente ADD CONSTRAINT PK_Antecedente
	PRIMARY KEY (id_antecedente)
;

CREATE INDEX IXFK_Antecedente_HistoriaClinica ON medicina.Antecedente (id_historia_clinica ASC)
;

CREATE INDEX IXFK_Antecedente_TipoAntecedente ON medicina.Antecedente (id_tipo_antecedente ASC)
;

ALTER TABLE psicologia.AntecedentePsicologico ADD CONSTRAINT PK_AntecedentePsicologico
	PRIMARY KEY (id_antecedente)
;

CREATE INDEX IXFK_AntecedentePsicologico_HistoriaClinica ON psicologia.AntecedentePsicologico (id_historia_clinica ASC)
;

CREATE INDEX IXFK_AntecedentePsicologico_TipoAntecedentePsicologico ON psicologia.AntecedentePsicologico (id_tipo_antecedente_psicologico ASC)
;

ALTER TABLE psicologia.ComportamientoConsulta ADD CONSTRAINT PK_ComportamientoConsulta
	PRIMARY KEY (id_comportamiento_consulta)
;

CREATE INDEX IXFK_ComportamientoConsulta_HojaHistoria ON psicologia.ComportamientoConsulta (id_hoja_historia ASC)
;

ALTER TABLE psicologia.ComposicionFamiliar ADD CONSTRAINT PK_ComposicionFamiliar
	PRIMARY KEY (id_composicion_familiar)
;

CREATE INDEX IXFK_ComposicionFamiliar_HojaHistoria ON psicologia.ComposicionFamiliar (id_hoja_historia ASC)
;

ALTER TABLE medicina.ConsultaFisioterapia ADD CONSTRAINT PK_ConsultaFisioterapia
	PRIMARY KEY (id_consulta_fisioterapia)
;

CREATE INDEX IXFK_ConsultaFisioterapia_HojaHistoria ON medicina.ConsultaFisioterapia (id_hoja_historia ASC)
;

ALTER TABLE odontologia.ControlPlaca ADD CONSTRAINT PK_ControlPlaca
	PRIMARY KEY (id_control_placa)
;

CREATE INDEX IXFK_ControlPlaca_HojaHistoria ON odontologia.ControlPlaca (id_hoja_historia ASC)
;

CREATE INDEX IXFK_ControlPlaca_TipoOdontograma ON odontologia.ControlPlaca (id_tipo_odontograma ASC)
;

ALTER TABLE medicina.DiagnosticoMedicina ADD CONSTRAINT PK_Diagnostico
	PRIMARY KEY (id_diagnostico)
;

ALTER TABLE odontologia.DiagnosticoOdontologia ADD CONSTRAINT PK_Diagnostico
	PRIMARY KEY (id_diagnostico)
;

CREATE INDEX IXFK_Diagnostico_HojaHistoria ON odontologia.DiagnosticoOdontologia (id_hoja_historia ASC)
;

ALTER TABLE psicologia.DiagnosticoPsicologia ADD CONSTRAINT PK_Diagnostico
	PRIMARY KEY (id_diagnostico_psicologia)
;

CREATE INDEX IXFK_Diagnostico_HojaHistoria ON psicologia.DiagnosticoPsicologia (id_hoja_historia ASC)
;

ALTER TABLE medicina.Examen ADD CONSTRAINT PK_Examen
	PRIMARY KEY (id_examen)
;

CREATE INDEX IXFK_Examen_HojaHistoria ON medicina.Examen (id_hoja_historia ASC)
;

CREATE INDEX IXFK_Examen_TipoExamen ON medicina.Examen (id_tipo_examen ASC)
;

ALTER TABLE odontologia.ExamenDental ADD CONSTRAINT PK_ExamenDental
	PRIMARY KEY (id_examen_dental)
;

CREATE INDEX IXFK_ExamenDental_HojaHistoria ON odontologia.ExamenDental (id_hoja_historia ASC)
;

ALTER TABLE odontologia.ExamenesComplementarios ADD CONSTRAINT PK_ExamenesComplementarios
	PRIMARY KEY (id_examenes_complementarios)
;

CREATE INDEX IXFK_ExamenesComplementarios_HojaHistoria ON odontologia.ExamenesComplementarios (id_hoja_historia ASC)
;

ALTER TABLE odontologia.ExamenEstomatologico ADD CONSTRAINT PK_ExamenEstomatologico
	PRIMARY KEY (id_examen_estomatologico)
;

CREATE INDEX IXFK_ExamenEstomatologico_HojaHistoria ON odontologia.ExamenEstomatologico (id_hoja_historia ASC)
;

ALTER TABLE medicina.HistoriaClinica ADD CONSTRAINT PK_HistoriaClinica
	PRIMARY KEY (id_historia_clinica)
;

ALTER TABLE medicina.HojaHistoria ADD CONSTRAINT PK_HojaHistoria
	PRIMARY KEY (id_hoja_historia)
;

CREATE INDEX IXFK_HojaHistoria_Diagnostico ON medicina.HojaHistoria (id_diagnostico ASC)
;

CREATE INDEX IXFK_HojaHistoria_HistoriaClinica ON medicina.HojaHistoria (id_historia_clinica ASC)
;

ALTER TABLE psicologia.Limites ADD CONSTRAINT PK_Limites
	PRIMARY KEY (id_limite)
;

CREATE INDEX IXFK_Limites_HojaHistoria ON psicologia.Limites (id_hoja_historia ASC)
;

ALTER TABLE odontologia.Odontograma ADD CONSTRAINT PK_Odontograma
	PRIMARY KEY (id_odontograma)
;

CREATE INDEX IXFK_Odontograma_HojaHistoria ON odontologia.Odontograma (id_hoja_historia ASC)
;

CREATE INDEX IXFK_Odontograma_TipoOdontograma ON odontologia.Odontograma (id_tipo_odontograma ASC)
;

ALTER TABLE medicina.Sistemas ADD CONSTRAINT PK_Sistemas
	PRIMARY KEY (id_sistema)
;

CREATE INDEX IXFK_Sistemas_HojaHistoria ON medicina.Sistemas (id_hoja_historia ASC)
;

ALTER TABLE medicina.TipoAntecedente ADD CONSTRAINT PK_TipoAntecedente
	PRIMARY KEY (id_tipo_antecedente)
;

ALTER TABLE psicologia.TipoAntecedentePsicologico ADD CONSTRAINT PK_TipoAntecedentePsicologico
	PRIMARY KEY (id_tipo_antecedente_psicologico)
;

ALTER TABLE medicina.TipoExamen ADD CONSTRAINT PK_TipoExamen
	PRIMARY KEY (id_tipo_examen)
;

ALTER TABLE odontologia.TipoOdontograma ADD CONSTRAINT PK_TipoOdontograma
	PRIMARY KEY (id_tipo_odontograma)
;

ALTER TABLE psicologia.ValoracionInterpersonal ADD CONSTRAINT PK_ValoracionInterpersonal
	PRIMARY KEY (id_valoracion_interpersonal)
;

CREATE INDEX IXFK_ValoracionInterpersonal_HojaHistoria ON psicologia.ValoracionInterpersonal (id_hoja_historia ASC)
;

/* Create Foreign Key Constraints */

ALTER TABLE acceso_historia.AccesoHistoriaClinica ADD CONSTRAINT FK_AccesoHistoriaClinica_HistoriaClinica
	FOREIGN KEY (id_historia_clinica) REFERENCES medicina.HistoriaClinica (id_historia_clinica) ON DELETE No Action ON UPDATE No Action
;

ALTER TABLE odontologia.Anamnesis ADD CONSTRAINT FK_Anamnesis_HistoriaClinica
	FOREIGN KEY (id_historia_clinica) REFERENCES medicina.HistoriaClinica (id_historia_clinica) ON DELETE No Action ON UPDATE No Action
;

ALTER TABLE medicina.Antecedente ADD CONSTRAINT FK_Antecedente_HistoriaClinica
	FOREIGN KEY (id_historia_clinica) REFERENCES medicina.HistoriaClinica (id_historia_clinica) ON DELETE No Action ON UPDATE No Action
;

ALTER TABLE medicina.Antecedente ADD CONSTRAINT FK_Antecedente_TipoAntecedente
	FOREIGN KEY (id_tipo_antecedente) REFERENCES medicina.TipoAntecedente (id_tipo_antecedente) ON DELETE No Action ON UPDATE No Action
;

ALTER TABLE psicologia.AntecedentePsicologico ADD CONSTRAINT FK_AntecedentePsicologico_HistoriaClinica
	FOREIGN KEY (id_historia_clinica) REFERENCES medicina.HistoriaClinica (id_historia_clinica) ON DELETE No Action ON UPDATE No Action
;

ALTER TABLE psicologia.AntecedentePsicologico ADD CONSTRAINT FK_AntecedentePsicologico_TipoAntecedentePsicologico
	FOREIGN KEY (id_tipo_antecedente_psicologico) REFERENCES psicologia.TipoAntecedentePsicologico (id_tipo_antecedente_psicologico) ON DELETE No Action ON UPDATE No Action
;

ALTER TABLE psicologia.ComportamientoConsulta ADD CONSTRAINT FK_ComportamientoConsulta_HojaHistoria
	FOREIGN KEY (id_hoja_historia) REFERENCES medicina.HojaHistoria (id_hoja_historia) ON DELETE No Action ON UPDATE No Action
;

ALTER TABLE psicologia.ComposicionFamiliar ADD CONSTRAINT FK_ComposicionFamiliar_HojaHistoria
	FOREIGN KEY (id_hoja_historia) REFERENCES medicina.HojaHistoria (id_hoja_historia) ON DELETE No Action ON UPDATE No Action
;

ALTER TABLE medicina.ConsultaFisioterapia ADD CONSTRAINT FK_ConsultaFisioterapia_HojaHistoria
	FOREIGN KEY (id_hoja_historia) REFERENCES medicina.HojaHistoria (id_hoja_historia) ON DELETE No Action ON UPDATE No Action
;

ALTER TABLE odontologia.ControlPlaca ADD CONSTRAINT FK_ControlPlaca_HojaHistoria
	FOREIGN KEY (id_hoja_historia) REFERENCES medicina.HojaHistoria (id_hoja_historia) ON DELETE No Action ON UPDATE No Action
;

ALTER TABLE odontologia.ControlPlaca ADD CONSTRAINT FK_ControlPlaca_TipoOdontograma
	FOREIGN KEY (id_tipo_odontograma) REFERENCES odontologia.TipoOdontograma (id_tipo_odontograma) ON DELETE No Action ON UPDATE No Action
;

ALTER TABLE odontologia.DiagnosticoOdontologia ADD CONSTRAINT FK_Diagnostico_HojaHistoria
	FOREIGN KEY (id_hoja_historia) REFERENCES medicina.HojaHistoria (id_hoja_historia) ON DELETE No Action ON UPDATE No Action
;

ALTER TABLE psicologia.DiagnosticoPsicologia ADD CONSTRAINT FK_Diagnostico_HojaHistoria
	FOREIGN KEY (id_hoja_historia) REFERENCES medicina.HojaHistoria (id_hoja_historia) ON DELETE No Action ON UPDATE No Action
;

ALTER TABLE medicina.Examen ADD CONSTRAINT FK_Examen_HojaHistoria
	FOREIGN KEY (id_hoja_historia) REFERENCES medicina.HojaHistoria (id_hoja_historia) ON DELETE No Action ON UPDATE No Action
;

ALTER TABLE medicina.Examen ADD CONSTRAINT FK_Examen_TipoExamen
	FOREIGN KEY (id_tipo_examen) REFERENCES medicina.TipoExamen (id_tipo_examen) ON DELETE No Action ON UPDATE No Action
;

ALTER TABLE odontologia.ExamenDental ADD CONSTRAINT FK_ExamenDental_HojaHistoria
	FOREIGN KEY (id_hoja_historia) REFERENCES medicina.HojaHistoria (id_hoja_historia) ON DELETE No Action ON UPDATE No Action
;

ALTER TABLE odontologia.ExamenesComplementarios ADD CONSTRAINT FK_ExamenesComplementarios_HojaHistoria
	FOREIGN KEY (id_hoja_historia) REFERENCES medicina.HojaHistoria (id_hoja_historia) ON DELETE No Action ON UPDATE No Action
;

ALTER TABLE odontologia.ExamenEstomatologico ADD CONSTRAINT FK_ExamenEstomatologico_HojaHistoria
	FOREIGN KEY (id_hoja_historia) REFERENCES medicina.HojaHistoria (id_hoja_historia) ON DELETE No Action ON UPDATE No Action
;

ALTER TABLE medicina.HojaHistoria ADD CONSTRAINT FK_HojaHistoria_Diagnostico
	FOREIGN KEY (id_diagnostico) REFERENCES medicina.DiagnosticoMedicina (id_diagnostico) ON DELETE No Action ON UPDATE No Action
;

ALTER TABLE medicina.HojaHistoria ADD CONSTRAINT FK_HojaHistoria_HistoriaClinica
	FOREIGN KEY (id_historia_clinica) REFERENCES medicina.HistoriaClinica (id_historia_clinica) ON DELETE No Action ON UPDATE No Action
;

ALTER TABLE psicologia.Limites ADD CONSTRAINT FK_Limites_HojaHistoria
	FOREIGN KEY (id_hoja_historia) REFERENCES medicina.HojaHistoria (id_hoja_historia) ON DELETE No Action ON UPDATE No Action
;

ALTER TABLE odontologia.Odontograma ADD CONSTRAINT FK_Odontograma_HojaHistoria
	FOREIGN KEY (id_hoja_historia) REFERENCES medicina.HojaHistoria (id_hoja_historia) ON DELETE No Action ON UPDATE No Action
;

ALTER TABLE odontologia.Odontograma ADD CONSTRAINT FK_Odontograma_TipoOdontograma
	FOREIGN KEY (id_tipo_odontograma) REFERENCES odontologia.TipoOdontograma (id_tipo_odontograma) ON DELETE No Action ON UPDATE No Action
;

ALTER TABLE medicina.Sistemas ADD CONSTRAINT FK_Sistemas_HojaHistoria
	FOREIGN KEY (id_hoja_historia) REFERENCES medicina.HojaHistoria (id_hoja_historia) ON DELETE No Action ON UPDATE No Action
;

ALTER TABLE psicologia.ValoracionInterpersonal ADD CONSTRAINT FK_ValoracionInterpersonal_HojaHistoria
	FOREIGN KEY (id_hoja_historia) REFERENCES medicina.HojaHistoria (id_hoja_historia) ON DELETE No Action ON UPDATE No Action
;

/* Create Table Comments, Sequences for Autonumber Columns */

COMMENT ON COLUMN odontologia.ExamenesComplementarios.periapical_inicio
	IS 'archivo'
;
