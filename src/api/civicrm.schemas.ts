// Auto-generated — do not edit manually
// Run generate_civicrm_types.py to regenerate
import { z } from "zod"


// ────────────────────────────────────────────────────────
// ACL
// ────────────────────────────────────────────────────────

export const ACLSchema = z.object({
  // Eindeutige Tabellen-ID
  id: z.number().int().optional(),
  // ACL Name
  name: z.string().nullable().optional(),
  // Ist dieser ACL-Eintrag Erlauben (0) oder Verbieten (1)?
  deny: z.boolean().optional(),
  // Table of the object possessing this ACL entry (Contact, Group, or ACL Group)
  entity_table: z.string(),
  // ID of the object possessing this ACL
  entity_id: z.number().int().nullable().optional(),
  // What operation does this ACL entry control?
  operation: z.string(),
  // The table of the object controlled by this ACL entry
  object_table: z.string().nullable().optional(),
  // The ID of the object controlled by this ACL entry
  object_id: z.number().int().nullable().optional(),
  // If this is a grant/revoke entry, what table are we granting?
  acl_table: z.string().nullable().optional(),
  // ID of the ACL or ACL group being granted/revoked
  acl_id: z.number().int().nullable().optional(),
  // Ist diese Eigenschaft aktiv?
  is_active: z.boolean().optional(),
  // Priorität
  priority: z.number().int().optional(),
})
export type ACL = z.infer<typeof ACLSchema>


// ────────────────────────────────────────────────────────
// ACLEntityRole
// ────────────────────────────────────────────────────────

export const ACLEntityRoleSchema = z.object({
  // Eindeutige Tabellen-ID
  id: z.number().int().optional(),
  // Foreign Key to ACL Role (which is an option value pair and hence an implicit FK)
  acl_role_id: z.number().int(),
  // Table of the object joined to the ACL Role (Contact or Group)
  entity_table: z.string(),
  // ID of the group/contact object being joined
  entity_id: z.number().int(),
  // Ist diese Eigenschaft aktiv?
  is_active: z.boolean().optional(),
})
export type ACLEntityRole = z.infer<typeof ACLEntityRoleSchema>


// ────────────────────────────────────────────────────────
// ActionSchedule
// ────────────────────────────────────────────────────────

export const ActionScheduleSchema = z.object({
  // Action Schedule ID
  id: z.number().int().optional(),
  // Name of the scheduled action
  name: z.string(),
  // Title of the action(reminder)
  title: z.string().nullable().optional(),
  // Empfänger
  recipient: z.string().nullable().optional(),
  // Is this the recipient criteria limited to OR in addition to?
  limit_to: z.number().int().nullable().optional(),
  // Entitäts-Wert
  entity_value: z.string().nullable().optional(),
  // Entitäts-Status
  entity_status: z.string().nullable().optional(),
  // Reminder Intervall
  start_action_offset: z.number().int().nullable().optional(),
  // Zeiteinheiten für Erinnerung (Reminder).
  start_action_unit: z.string().nullable().optional(),
  // Reminder Action
  start_action_condition: z.string().nullable().optional(),
  // Entitäts-Datum
  start_action_date: z.string().nullable().optional(),
  // Wiederholen
  is_repeat: z.boolean().optional(),
  // Time units for repetition of reminder.
  repetition_frequency_unit: z.string().nullable().optional(),
  // Time interval for repeating the reminder.
  repetition_frequency_interval: z.number().int().nullable().optional(),
  // Time units till repetition of reminder.
  end_frequency_unit: z.string().nullable().optional(),
  // Time interval till repeating the reminder.
  end_frequency_interval: z.number().int().nullable().optional(),
  // Reminder Action till repeating the reminder.
  end_action: z.string().nullable().optional(),
  // Entity end date
  end_date: z.string().nullable().optional(),
  // Ist diese Option aktiv?
  is_active: z.boolean().optional(),
  // Contact IDs to which reminder should be sent.
  recipient_manual: z.string().nullable().optional(),
  // listing based on recipient field.
  recipient_listing: z.string().nullable().optional(),
  // Body des Mailings im Text-Format.
  body_text: z.string().nullable().optional(),
  // Body des Mailings im HTML-Format.
  body_html: z.string().nullable().optional(),
  // Inhalt vom SMS Text.
  sms_body_text: z.string().nullable().optional(),
  // Betreff der E-Mail
  subject: z.string().nullable().optional(),
  // Aktivität für diesen Reminder aufzeichnen?
  record_activity: z.boolean().optional(),
  // Name/ID of the mapping to use on this table
  mapping_id: z.string().nullable().optional(),
  // FK zur Gruppe
  group_id: z.number().int().nullable().optional(),
  // FK zur Nachrichtenvorlage.
  msg_template_id: z.number().int().nullable().optional(),
  // FK zur Nachrichtenvorlage.
  sms_template_id: z.number().int().nullable().optional(),
  // Datum, an dem die Erinnerung gesendet wird.
  absolute_date: z.string().nullable().optional(),
  // Name im "Von"-Feld
  from_name: z.string().nullable().optional(),
  // E-Mail-Adresse im "Von"-Feld
  from_email: z.string().nullable().optional(),
  // Sende die Nachricht als E-Mail oder SMS oder beides.
  mode: z.string().nullable().optional(),
  // ID des SMS-Anbieters
  sms_provider_id: z.number().int().nullable().optional(),
  // Used for repeating entity
  used_for: z.string().nullable().optional(),
  // Benutzt für mehrsprachige Installationen
  filter_contact_language: z.string().nullable().optional(),
  // Benutzt für mehrsprachige Installationen
  communication_language: z.string().nullable().optional(),
  // When was the scheduled reminder created.
  created_date: z.string().nullable().optional(),
  // Wann die Erinnerung erstellt oder bearbeitet wurde.
  modified_date: z.string().nullable().optional(),
  // Earliest date to consider start events from.
  effective_start_date: z.string().nullable().optional(),
  // Latest date to consider end events from.
  effective_end_date: z.string().nullable().optional(),
})
export type ActionSchedule = z.infer<typeof ActionScheduleSchema>


// ────────────────────────────────────────────────────────
// Activity
// ────────────────────────────────────────────────────────

export const ActivitySchema = z.object({
  // Unique Other Activity ID
  id: z.number().int().optional(),
  // Artificial FK to original transaction (e.g. contribution) IF it is not an Activity. Entity table is discovered by filtering by the appropriate activity_type_id.
  source_record_id: z.number().int().nullable().optional(),
  // FK to civicrm_option_value.value, that has to be valid, registered activity type.
  activity_type_id: z.number().int().optional(),
  // Der Betreff/Zweck/kurze Beschreibung dieser Aktivität.
  subject: z.string().nullable().optional(),
  // Date and time this activity is scheduled to occur. Formerly named scheduled_date_time.
  activity_date_time: z.string().nullable().optional(),
  // Planned or actual duration of activity expressed in minutes. Conglomerate of former duration_hours and duration_minutes.
  duration: z.number().int().nullable().optional(),
  // Ort der Aktivität (optional, Freitext)
  location: z.string().nullable().optional(),
  // Phone ID of the number called (optional - used if an existing phone number is selected).
  phone_id: z.number().int().nullable().optional(),
  // Phone number in case the number does not exist in the civicrm_phone table.
  phone_number: z.string().nullable().optional(),
  // Einzelheiten zur Aktivität (Notizen etc.)
  details: z.string().nullable().optional(),
  // ID of the status this activity is currently in. Foreign key to civicrm_option_value.
  status_id: z.number().int().nullable().optional(),
  // ID of the priority given to this activity. Foreign key to civicrm_option_value.
  priority_id: z.number().int().nullable().optional(),
  // Parent meeting ID (if this is a follow-up item).
  parent_id: z.number().int().nullable().optional(),
  // Test
  is_test: z.boolean().optional().default(false),
  // Activity Medium, Implicit FK to civicrm_option_value where option_group = encounter_medium.
  medium_id: z.number().int().nullable().optional(),
  // Auto
  is_auto: z.boolean().optional(),
  // FK to Relationship ID
  relationship_id: z.number().int().nullable().optional(),
  // Unused deprecated column.
  is_current_revision: z.boolean().optional(),
  // Unused deprecated column.
  original_id: z.number().int().nullable().optional(),
  // Currently being used to store result id for survey activity, FK to option value.
  result: z.string().nullable().optional(),
  // Aktivität wurde in den Papierkorb verschoben.
  is_deleted: z.boolean().optional().default(false),
  // Assign a specific level of engagement to this activity. Used for tracking constituents in ladder of engagement.
  engagement_level: z.number().int().nullable().optional(),
  // Reihenfolge
  weight: z.number().int().nullable().optional(),
  // Aktivität als Favorit markiert.
  is_star: z.boolean().optional(),
  // Wann die Aktivität erstellt wurde.
  created_date: z.string().nullable().optional(),
  // When was the activity (or closely related entity) was created or modified or deleted.
  modified_date: z.string().nullable().optional(),
  // CiviCase this activity belongs to.
  case_id: z.number().int().nullable().optional(),
  // Contact who created this activity.
  source_contact_id: z.number().int().nullable().optional(),
  // Contacts involved in this activity.
  target_contact_id: z.array(z.unknown()).nullable().optional(),
  // Contacts assigned to this activity.
  assignee_contact_id: z.array(z.unknown()).nullable().optional(),
  // Filter by tags (including child tags)
  tags: z.array(z.unknown()).nullable().optional(),
  // Tiefe in der geschachtelten Hierarchie
  _depth: z.number().int().nullable().optional(),
  // Number of descendents in the nested hierarchy
  _descendents: z.number().int().nullable().optional(),
})
export type Activity = z.infer<typeof ActivitySchema>


// ────────────────────────────────────────────────────────
// ActivityContact
// ────────────────────────────────────────────────────────

export const ActivityContactSchema = z.object({
  // Aktivität Kontakt-ID
  id: z.number().int().optional(),
  // Foreign key to the activity for this record.
  activity_id: z.number().int(),
  // Foreign key to the contact for this record.
  contact_id: z.number().int(),
  // Determines the contact's role in the activity (source, target, or assignee).
  record_type_id: z.number().int().nullable().optional(),
})
export type ActivityContact = z.infer<typeof ActivityContactSchema>


// ────────────────────────────────────────────────────────
// Address
// ────────────────────────────────────────────────────────

export const AddressSchema = z.object({
  // Eindeutige Adress-ID
  id: z.number().int().optional(),
  // FK zu Kontakt ID
  contact_id: z.number().int().nullable().optional(),
  // Zu welcher Adresskategrie diese Adresse gehört.
  location_type_id: z.number().int().nullable().optional(),
  // Ist das die Hauptadresse.
  is_primary: z.boolean().optional(),
  // Ist das die Rechnungsadresse.
  is_billing: z.boolean().optional(),
  // Concatenation of all routable street address components (prefix, street number, street name, suffix, unit number OR P.O. Box). Apps should be able to determine physical location with this data (for mapping, mail delivery, etc.).
  street_address: z.string().nullable().optional(),
  // Numeric portion of address number on the street, e.g. For 112A Main St, the street_number = 112.
  street_number: z.number().int().nullable().optional(),
  // Non-numeric portion of address number on the street, e.g. For 112A Main St, the street_number_suffix = A
  street_number_suffix: z.string().nullable().optional(),
  // Richtungs-Präfix, z.B. SE Main Str., SE ist der Präfix.
  street_number_predirectional: z.string().nullable().optional(),
  // konkreter Straßenname, ausgenommen St, Dr, Rd, Ave; z. B. ist für "112 Main St" der Straßenname = Main
  street_name: z.string().nullable().optional(),
  // Str, Pl. Bv, etc.
  street_type: z.string().nullable().optional(),
  // Richtungs-Suffix, z.B. Main Str. S, S ist Suffix.
  street_number_postdirectional: z.string().nullable().optional(),
  // Secondary unit designator, e.g. Apt 3 or Unit # 14, or Bldg 1200
  street_unit: z.string().nullable().optional(),
  // Zusätzliche Adressinformation, Zeile 1
  supplemental_address_1: z.string().nullable().optional(),
  // Zusätzliche Adressinformation, Zeile 2
  supplemental_address_2: z.string().nullable().optional(),
  // Zusätzliche Adressinformation, Zeile 3
  supplemental_address_3: z.string().nullable().optional(),
  // Name der Stadt, Ort oder Dorf.
  city: z.string().nullable().optional(),
  // Zu welchem Landkreis/Bezirk diese Adresse gehört.
  county_id: z.number().int().nullable().optional(),
  // Zu welchem Bundesland / Provinz diese Adresse gehört.
  state_province_id: z.number().int().nullable().optional(),
  // Speichert den Suffix, wie der +4 Teil im USPS-System.
  postal_code_suffix: z.string().nullable().optional(),
  // Store both US (zip5) AND international postal codes. App is responsible for country/region appropriate validation.
  postal_code: z.string().nullable().optional(),
  // USPS Massenmail Code.
  usps_adc: z.string().nullable().optional(),
  // Zu welchem Land diese Adresse gehört.
  country_id: z.number().int().nullable().optional(),
  // Breitengrad
  geo_code_1: z.number().nullable().optional(),
  // Längengrad
  geo_code_2: z.number().nullable().optional(),
  // Ist das ein manuell eingegebener Geocode
  manual_geo_code: z.boolean().optional(),
  // Timezone expressed as a UTC offset - e.g. United States CST would be written as "UTC-6".
  timezone: z.string().nullable().optional(),
  // Adressenname
  name: z.string().nullable().optional(),
  // FK to Address ID
  master_id: z.number().int().nullable().optional(),
  // Address is within a given distance to a location
  proximity: z.boolean().nullable().optional(),
})
export type Address = z.infer<typeof AddressSchema>


// ────────────────────────────────────────────────────────
// Afform
// ────────────────────────────────────────────────────────

export const AfformSchema = z.object({
  // Name
  name: z.string().nullable().optional(),
  // Typ
  type: z.string().nullable().optional().default("form"),
  // Angular module dependencies; calculated at runtime
  requires: z.array(z.unknown()).nullable().optional(),
  // Block used for this entity type
  entity_type: z.string().nullable().optional(),
  // Used for blocks that join a sub-entity (e.g. Emails for a Contact)
  join_entity: z.string().nullable().optional(),
  // Titel
  title: z.string().nullable().optional(),
  // Beschreibung
  description: z.string().nullable().optional(),
  // Placement
  placement: z.array(z.unknown()).nullable().optional(),
  // E.g. contact_type, case_type, event_type, etc.
  placement_filters: z.array(z.unknown()).nullable().optional(),
  // Placement Order
  placement_weight: z.number().int().nullable().optional(),
  // Tags
  tags: z.array(z.unknown()).nullable().optional(),
  // Icon shown in the placement
  icon: z.string().nullable().optional(),
  // Page Route
  server_route: z.string().nullable().optional(),
  // Is Public
  is_public: z.boolean().nullable().optional().default(false),
  // Berechtigung
  permission: z.array(z.unknown()).nullable().optional(),
  // Rechte-Operator
  permission_operator: z.string().nullable().optional().default("AND"),
  // Post-Submit Page
  redirect: z.string().nullable().optional(),
  // Allow Submissions
  submit_enabled: z.boolean().nullable().optional().default(true),
  // Max Submissions (total)
  submit_limit: z.number().int().nullable().optional(),
  // Max Submissions (per user)
  submit_limit_per_user: z.number().int().nullable().optional(),
  // Keep a log of the date, time, user, and items saved by each form submission.
  create_submission: z.boolean().nullable().optional(),
  // Verify submission before processing
  manual_processing: z.boolean().nullable().optional(),
  // Allow verification by email
  allow_verification_by_email: z.boolean().nullable().optional(),
  // Email Template
  email_confirmation_template_id: z.number().int().nullable().optional(),
  // For authenticated users, form will auto-save periodically.
  autosave_draft: z.boolean().nullable().optional(),
  // Insert into navigation menu {parent: string, label: string, weight: int}
  navigation: z.array(z.unknown()).nullable().optional(),
  // HTML form layout; format is controlled by layoutFormat param
  layout: z.array(z.unknown()).nullable().optional(),
  // Date Modified
  modified_date: z.string().nullable().optional(),
  // Confirmation Type
  confirmation_type: z.string().nullable().optional().default("redirect_to_url"),
  // Bestätigungsnachricht
  confirmation_message: z.string().nullable().optional(),
  // Erstellt von Kontakt-ID
  created_id: z.number().int().nullable().optional(),
  // Locale
  locale: z.string().nullable().optional(),
  // Name of generated Angular module (CamelCase)
  module_name: z.string().nullable().optional(),
  // Html tag name to invoke this form (dash-case)
  directive_name: z.string().nullable().optional(),
  // Number of submission records for this form
  submission_count: z.number().int().nullable().optional(),
  // Number of submission records for the current user
  user_submission_count: z.number().int().nullable().optional(),
  // Date & time of last form submission
  submission_date: z.string().nullable().optional(),
  // Based on settings and current submission count, is the form open for submissions
  submit_currently_open: z.boolean().nullable().optional(),
  // Whether a local copy is saved on site
  has_local: z.boolean().nullable().optional(),
  // Is provided by an extension
  has_base: z.boolean().nullable().optional(),
  // Name of extension which provides this form
  base_module: z.string().nullable().optional(),
  // Embedded search displays, formatted like ["search-name.display-name"]
  search_displays: z.array(z.unknown()).nullable().optional(),
})
export type Afform = z.infer<typeof AfformSchema>


// ────────────────────────────────────────────────────────
// AfformBehavior
// ────────────────────────────────────────────────────────

export const AfformBehaviorSchema = z.object({
  // Unique identifier in dashed-format, name of entity attribute for selected mode
  key: z.string().nullable().optional(),
  // Array of attributes added to the entity by this behavior, keyed by attribute name
  attributes: z.array(z.unknown()).nullable().optional(),
  // Localized title displayed on admin screen
  title: z.string().nullable().optional(),
  // Optional localized description displayed on admin screen
  description: z.string().nullable().optional(),
  // Optional template for configuring the behavior in the AfformGuiEditor
  template: z.string().nullable().optional(),
  // Afform entities this behavior supports
  entities: z.array(z.unknown()).nullable().optional(),
  // Nested array of supported behavior modes, keyed by entity name
  modes: z.array(z.unknown()).nullable().optional(),
  // If set then mode will not be de-selectable
  default_mode: z.string().nullable().optional(),
})
export type AfformBehavior = z.infer<typeof AfformBehaviorSchema>


// ────────────────────────────────────────────────────────
// AfformSubmission
// ────────────────────────────────────────────────────────

export const AfformSubmissionSchema = z.object({
  // Unique Submission ID
  id: z.number().int().optional(),
  // User Contact ID
  contact_id: z.number().int().nullable().optional(),
  // Name of submitted afform
  afform_name: z.string().nullable().optional(),
  // IDs of saved entities
  data: z.string().nullable().optional(),
  // Submission Date/Time
  submission_date: z.string().nullable().optional(),
  // fk to Afform Submission Status options in civicrm_option_values
  status_id: z.number().int().optional(),
})
export type AfformSubmission = z.infer<typeof AfformSubmissionSchema>


// ────────────────────────────────────────────────────────
// Batch
// ────────────────────────────────────────────────────────

export const BatchSchema = z.object({
  // Eindeutige Adress-ID
  id: z.number().int().optional(),
  // Variable name/programmatic handle for this batch.
  name: z.string().nullable().optional(),
  // Friendly Name.
  title: z.string().nullable().optional(),
  // Description of this batch set.
  description: z.string().nullable().optional(),
  // FK zu Kontakt ID
  created_id: z.number().int().nullable().optional(),
  // When was this item created
  created_date: z.string().optional(),
  // FK zu Kontakt ID
  modified_id: z.number().int().nullable().optional(),
  // When was this item modified
  modified_date: z.string().optional(),
  // FK to Saved Search ID
  saved_search_id: z.number().int().nullable().optional(),
  // fk to Batch Status options in civicrm_option_values
  status_id: z.number().int(),
  // fk to Batch Type options in civicrm_option_values
  type_id: z.number().int().nullable().optional(),
  // fk to Batch mode options in civicrm_option_values
  mode_id: z.number().int().nullable().optional(),
  // Total amount for this batch.
  total: z.number().nullable().optional(),
  // Number of items in a batch.
  item_count: z.number().int().nullable().optional(),
  // fk to Payment Instrument options in civicrm_option_values
  payment_instrument_id: z.number().int().nullable().optional(),
  // Batch Exported Date
  exported_date: z.string().nullable().optional(),
  // cache entered data
  data: z.string().nullable().optional(),
})
export type Batch = z.infer<typeof BatchSchema>


// ────────────────────────────────────────────────────────
// BouncePattern
// ────────────────────────────────────────────────────────

export const BouncePatternSchema = z.object({
  // Bounce (Rückweisung) Regelmäßigkeit ID
  id: z.number().int().optional(),
  // Typ der Zurückweisung
  bounce_type_id: z.number().int(),
  // A regexp to match a message to a bounce type
  pattern: z.string().nullable().optional(),
})
export type BouncePattern = z.infer<typeof BouncePatternSchema>


// ────────────────────────────────────────────────────────
// BounceType
// ────────────────────────────────────────────────────────

export const BounceTypeSchema = z.object({
  // Bounce (Rückweisung) Typ ID
  id: z.number().int().optional(),
  // Typ der Zurückweisung
  name: z.string(),
  // Eine Beschreibung des Bounce-Typs
  description: z.string().nullable().optional(),
  // Number of bounces of this type required before the email address is put on bounce hold
  hold_threshold: z.number().int(),
})
export type BounceType = z.infer<typeof BounceTypeSchema>


// ────────────────────────────────────────────────────────
// Case
// ────────────────────────────────────────────────────────

export const CaseSchema = z.object({
  // Eindeutige Fall-ID
  id: z.number().int().optional(),
  // FK to civicrm_case_type.id
  case_type_id: z.number().int(),
  // Kurzbezeichnung des Falls.
  subject: z.string().nullable().optional(),
  // Datum an dem der aktuelle Fall beginnt.
  start_date: z.string().nullable().optional(),
  // Datum an dem der aktuelle Fall endet.
  end_date: z.string().nullable().optional(),
  // Details von Open Case. Wird nur in der Erweiterung CiviCase benutzt.
  details: z.string().nullable().optional(),
  // ID des Fallstatus
  status_id: z.number().int(),
  // Fall befindet sich im Papierkorb
  is_deleted: z.boolean().optional().default(false),
  // Wann der Fall eröffnet wurde.
  created_date: z.string().nullable().optional(),
  // Wann der Fall (oder die geschlossene, verbundene Entität) erstellt oder geändert oder gelöscht wurde.
  modified_date: z.string().nullable().optional(),
  // Filter by tags (including child tags)
  tags: z.array(z.unknown()).nullable().optional(),
})
export type Case = z.infer<typeof CaseSchema>


// ────────────────────────────────────────────────────────
// CaseActivity
// ────────────────────────────────────────────────────────

export const CaseActivitySchema = z.object({
  // Eindeutige Fall-Aktivitäten-Verbindungs ID
  id: z.number().int().optional(),
  // Fall ID der Fall-Aktivität-Verbindung.
  case_id: z.number().int(),
  // Aktivitäts ID der Fall-Aktivität-Verbindung.
  activity_id: z.number().int(),
})
export type CaseActivity = z.infer<typeof CaseActivitySchema>


// ────────────────────────────────────────────────────────
// CaseContact
// ────────────────────────────────────────────────────────

export const CaseContactSchema = z.object({
  // Eindeutige Fall-Kontakt-Verbindungs ID
  id: z.number().int().optional(),
  // Fall ID der Fall-Kontakt-Verbindung.
  case_id: z.number().int(),
  // Kontakt-ID des Kontakts zu dem der Fall gehört.
  contact_id: z.number().int(),
})
export type CaseContact = z.infer<typeof CaseContactSchema>


// ────────────────────────────────────────────────────────
// CaseType
// ────────────────────────────────────────────────────────

export const CaseTypeSchema = z.object({
  // Autoinkrementierte Typ-ID
  id: z.number().int().optional(),
  // Maschinenname für Falltyp
  name: z.string(),
  // Name des Falltyps in natürlicher Sprache
  title: z.string(),
  // Beschreibung des Falltyps
  description: z.string().nullable().optional(),
  // Ist dieser Falltyp aktiviert?
  is_active: z.boolean().optional(),
  // Ist dieser Falltyp ein vordefinierter Systemtyp?
  is_reserved: z.boolean().optional(),
  // Reihenfolge der Falltypen
  weight: z.number().int().optional(),
  // XML-Definition des Falltyps
  definition: z.string().nullable().optional(),
  // Is provided by an extension
  has_base: z.boolean().nullable().optional(),
  // Name of extension which provides this package
  base_module: z.string().nullable().optional(),
  // When the managed entity was changed from its original settings
  local_modified_date: z.string().nullable().optional(),
})
export type CaseType = z.infer<typeof CaseTypeSchema>


// ────────────────────────────────────────────────────────
// Contact
// ────────────────────────────────────────────────────────

export const ContactSchema = z.object({
  // Eindeutige Kontakt-ID
  id: z.number().int().optional(),
  // Kontaktart.
  contact_type: z.string().nullable().optional(),
  // Unique trusted external ID (generally from a legacy app/datasource). Particularly useful for deduping operations.
  external_identifier: z.string().nullable().optional(),
  // Formatierter Name, der das bevorzugte Format für Anzeige/Druck/andere Ausgaben angibt.
  display_name: z.string().nullable().optional(),
  // Organisationsname
  organization_name: z.string().nullable().optional(),
  // May be used to over-ride contact view and edit templates.
  contact_sub_type: z.string().nullable().optional(),
  // Vorname
  first_name: z.string().nullable().optional(),
  // zweiter Vorname
  middle_name: z.string().nullable().optional(),
  // Nachname
  last_name: z.string().nullable().optional(),
  // Keine E-Mails senden
  do_not_email: z.boolean().optional(),
  // Nicht anrufen
  do_not_phone: z.boolean().optional(),
  // Nicht anschreiben
  do_not_mail: z.boolean().optional(),
  // Keine SMS senden
  do_not_sms: z.boolean().optional(),
  // Nicht weitergeben
  do_not_trade: z.boolean().optional(),
  // Hat sich dieser Kontakt von allen Massen-E-Mails der Organisation bzw. der Webseitendomain abgemeldet, sogenanntes Opt-Out?
  is_opt_out: z.boolean().optional(),
  // May be used for SSN, EIN/TIN, Household ID (census) or other applicable unique legal/government ID.
  legal_identifier: z.string().nullable().optional(),
  // Name used for sorting different contact types
  sort_name: z.string().nullable().optional(),
  // Pseudonym.
  nick_name: z.string().nullable().optional(),
  // Gesetzlicher Name.
  legal_name: z.string().nullable().optional(),
  // optionale URL für bevorzugtes Bild (Foto, Logo, etc.), welches für diesen Kontakt angezeigt wird.
  image_URL: z.string().nullable().optional(),
  // Was ist die bevorzugte Kommunikationsart.
  preferred_communication_method: z.string().nullable().optional(),
  // Which language is preferred for communication. FK to languages in civicrm_option_value.
  preferred_language: z.string().nullable().optional(),
  // Schlüssel zur Anfragenverifizierung bezogen auf diesen Kontakt.
  hash: z.string().nullable().optional(),
  // API-Schlüssel zur Anfragenverifizierung bezogen auf diesen Kontakt.
  api_key: z.string().nullable().optional(),
  // woher der Kontakt stammt, z.B. Import, Eintrag vom Spendenmodul...
  source: z.string().nullable().optional(),
  // Prefix or Title for name (Ms, Mr...). FK to prefix ID
  prefix_id: z.number().int().nullable().optional(),
  // Suffix for name (Jr, Sr...). FK to suffix ID
  suffix_id: z.number().int().nullable().optional(),
  // Formeller (akademisch oder ähnlich) Titel vor dem Namen. (Prof., Dr. etc.)
  formal_title: z.string().nullable().optional(),
  // Communication style (e.g. formal vs. familiar) to use with this contact. FK to communication styles in civicrm_option_value.
  communication_style_id: z.number().int().nullable().optional(),
  // FK to civicrm_option_value.value, that has to be valid registered Email Greeting.
  email_greeting_id: z.number().int().nullable().optional(),
  // Benutzerdefinierte E-Mail-Grußformel
  email_greeting_custom: z.string().nullable().optional(),
  // Cache Email Greeting.
  email_greeting_display: z.string().nullable().optional(),
  // FK to civicrm_option_value.value, that has to be valid registered Postal Greeting.
  postal_greeting_id: z.number().int().nullable().optional(),
  // Benutzerdefinierte Brief-Grußformel.
  postal_greeting_custom: z.string().nullable().optional(),
  // Cache Postal greeting.
  postal_greeting_display: z.string().nullable().optional(),
  // FK to civicrm_option_value.value, that has to be valid registered Addressee.
  addressee_id: z.number().int().nullable().optional(),
  // Benutzerdefinierter Adressat
  addressee_custom: z.string().nullable().optional(),
  // Cache Addressee.
  addressee_display: z.string().nullable().optional(),
  // Funktion
  job_title: z.string().nullable().optional(),
  // FK to gender ID
  gender_id: z.number().int().nullable().optional(),
  // Geburtsdatum
  birth_date: z.string().nullable().optional(),
  // Verstorben / Geschlossen
  is_deceased: z.boolean().optional(),
  // Datum Todestag / Existiert nicht mehr seit
  deceased_date: z.string().nullable().optional(),
  // Name des Haushalts
  household_name: z.string().nullable().optional(),
  // Optional FK to Primary Contact for this household.
  primary_contact_id: z.number().int().nullable().optional(),
  // Standard Industry Classification Code.
  sic_code: z.string().nullable().optional(),
  // the OpenID (or OpenID-style http://username.domain/) unique identifier for this contact mainly used for logging in to CiviCRM
  user_unique_id: z.string().nullable().optional(),
  // OPTIONAL FK zu civicrm_contact record.
  employer_id: z.number().int().nullable().optional(),
  // Kontakt ist im Papierkorb
  is_deleted: z.boolean().optional().default(false),
  // Wann der Kontakt erstellt wurde.
  created_date: z.string().nullable().optional(),
  // When was the contact (or closely related entity) was created or modified or deleted.
  modified_date: z.string().nullable().optional(),
  // Deprecated setting for text vs html mailings
  preferred_mail_format: z.string().nullable().optional(),
  // Primary Address ID
  address_primary: z.number().int().nullable().optional(),
  // Billing Address ID
  address_billing: z.number().int().nullable().optional(),
  // Primary Email ID
  email_primary: z.number().int().nullable().optional(),
  // Billing Email ID
  email_billing: z.number().int().nullable().optional(),
  // Primary Phone ID
  phone_primary: z.number().int().nullable().optional(),
  // Billing Phone ID
  phone_billing: z.number().int().nullable().optional(),
  // Primary IM ID
  im_primary: z.number().int().nullable().optional(),
  // Billing IM ID
  im_billing: z.number().int().nullable().optional(),
  // Groups (or sub-groups of groups) to which this contact belongs
  groups: z.array(z.unknown()).nullable().optional(),
  // Age of individual (in years)
  age_years: z.number().int().nullable().optional(),
  // Number of days until next birthday
  next_birthday: z.number().int().nullable().optional(),
  // Filter by tags (including child tags)
  tags: z.array(z.unknown()).nullable().optional(),
})
export type Contact = z.infer<typeof ContactSchema>


// ────────────────────────────────────────────────────────
// ContactType
// ────────────────────────────────────────────────────────

export const ContactTypeSchema = z.object({
  // Kontaktart ID
  id: z.number().int().optional(),
  // Interner Name der Kontaktart (oder Unterart)
  name: z.string(),
  // localized Name of Contact Type.
  label: z.string().nullable().optional(),
  // localized Optional verbose description of the type.
  description: z.string().nullable().optional(),
  // Bild-URL falls vorhanden
  image_URL: z.string().nullable().optional(),
  // crm-i icon Klasse, die die Kontaktart repräsentiert
  icon: z.string().nullable().optional(),
  // Optional FK to parent contact type.
  parent_id: z.number().int().nullable().optional(),
  // Ist dieser Eintrag aktiv?
  is_active: z.boolean().optional(),
  // Ist diese Kontaktart ein vordefinierter Systemtyp?
  is_reserved: z.boolean().optional(),
  // Tiefe in der geschachtelten Hierarchie
  _depth: z.number().int().nullable().optional(),
  // Number of descendents in the nested hierarchy
  _descendents: z.number().int().nullable().optional(),
  // Is provided by an extension
  has_base: z.boolean().nullable().optional(),
  // Name of extension which provides this package
  base_module: z.string().nullable().optional(),
  // When the managed entity was changed from its original settings
  local_modified_date: z.string().nullable().optional(),
})
export type ContactType = z.infer<typeof ContactTypeSchema>


// ────────────────────────────────────────────────────────
// Contribution
// ────────────────────────────────────────────────────────

export const ContributionSchema = z.object({
  // Zuwendungs-ID
  id: z.number().int().optional(),
  // FK zu Kontakt ID
  contact_id: z.number().int(),
  // FK to Financial Type for (total_amount - non_deductible_amount).
  financial_type_id: z.number().int().nullable().optional(),
  // The Contribution Page which triggered this contribution
  contribution_page_id: z.number().int().nullable().optional(),
  // FK to Payment Instrument
  payment_instrument_id: z.number().int().nullable().optional(),
  // Zuwendungsdatum
  receive_date: z.string().nullable().optional(),
  // Portion of total amount which is NOT tax deductible. Equal to total_amount for non-deductible financial types.
  non_deductible_amount: z.number().nullable().optional(),
  // Total amount of this contribution. Use market value for non-monetary gifts.
  total_amount: z.number(),
  // aktuelle Gebühr Zahlungsprozessor falls bekannt - kann 0 sein.
  fee_amount: z.number().nullable().optional(),
  // actual funds transfer amount. total less fees. if processor does not report actual fee during transaction, this is set to total_amount.
  net_amount: z.number().nullable().optional(),
  // unique transaction id. may be processor id, bank id + trans id, or account number + check number... depending on payment_method
  trxn_id: z.string().nullable().optional(),
  // unique invoice id, system generated or passed in
  invoice_id: z.string().nullable().optional(),
  // Human readable invoice number
  invoice_number: z.string().nullable().optional(),
  // 3 Zeichen String, Wert aus den Konfig-Einstellungen oder aus der Benutzereingabe.
  currency: z.string().nullable().optional(),
  // when was gift cancelled
  cancel_date: z.string().nullable().optional(),
  // Storno- / Rückzahlungsgrund
  cancel_reason: z.string().nullable().optional(),
  // when (if) receipt was sent. populated automatically for online donations w/ automatic receipting
  receipt_date: z.string().nullable().optional(),
  // when (if) was donor thanked
  thankyou_date: z.string().nullable().optional(),
  // Origin of this Contribution.
  source: z.string().nullable().optional(),
  // Amount Label
  amount_level: z.string().nullable().optional(),
  // Conditional foreign key to civicrm_contribution_recur id. Each contribution made in connection with a recurring contribution carries a foreign key to the recurring contribution record. This assumes we can track these processor initiated events.
  contribution_recur_id: z.number().int().nullable().optional(),
  // Test Mode
  is_test: z.boolean().optional().default(false),
  // Ist "Später zahlen"
  is_pay_later: z.boolean().optional(),
  // Contribution Status ID
  contribution_status_id: z.number().int().nullable().optional(),
  // Conditional foreign key to civicrm_address.id. We insert an address record for each contribution when we have associated billing name and address data.
  address_id: z.number().int().nullable().optional(),
  // Schecknummer
  check_number: z.string().nullable().optional(),
  // unique credit note id, system generated or passed in
  creditnote_id: z.string().nullable().optional(),
  // Total tax amount of this contribution.
  tax_amount: z.number().optional(),
  // Stores the date when revenue should be recognized.
  revenue_recognition_date: z.string().nullable().optional(),
  // Shows this is a template for recurring contributions.
  is_template: z.boolean().optional().default(false),
  // When was the contribution created.
  created_date: z.string().nullable().optional(),
  // When was the contribution created or modified or deleted.
  modified_date: z.string().nullable().optional(),
  // Bezahlter Betrag
  paid_amount: z.number().nullable().optional(),
  // Guthaben
  balance_amount: z.number().nullable().optional(),
  // Betrag ohne Steuer
  tax_exclusive_amount: z.number().nullable().optional(),
})
export type Contribution = z.infer<typeof ContributionSchema>


// ────────────────────────────────────────────────────────
// ContributionPage
// ────────────────────────────────────────────────────────

export const ContributionPageSchema = z.object({
  // Zuwendungs-ID
  id: z.number().int().optional(),
  // Contribution Page title. For top of page display
  title: z.string(),
  // Contribution Page Public title
  frontend_title: z.string().optional(),
  // Unique name for identifying contribution page
  name: z.string(),
  // Text and html allowed. Displayed below title.
  intro_text: z.string().nullable().optional(),
  // default financial type assigned to contributions submitted via this page, e.g. Contribution, Campaign Contribution
  financial_type_id: z.number().int().nullable().optional(),
  // Payment Processors configured for this contribution Page
  payment_processor: z.string().nullable().optional(),
  // if true - processing logic must reject transaction at confirmation stage if pay method != credit card
  is_credit_card_only: z.boolean().optional(),
  // if true - allows real-time monetary transactions otherwise non-monetary transactions
  is_monetary: z.boolean().optional(),
  // if true - allows recurring contributions, valid only for PayPal_Standard
  is_recur: z.boolean().optional(),
  // if FALSE, the confirm page in contribution pages gets skipped
  is_confirm_enabled: z.boolean().optional(),
  // Supported recurring frequency units.
  recur_frequency_unit: z.string().nullable().optional(),
  // if true - supports recurring intervals
  is_recur_interval: z.boolean().optional(),
  // if true - asks user for recurring installments
  is_recur_installments: z.boolean().optional(),
  // if true - user is able to adjust payment start date
  adjust_recur_start_date: z.boolean().optional(),
  // falls aktiv - erlaubt dem Benutzer, die Zahlung später direkt an zu senden
  is_pay_later: z.boolean().optional(),
  // Der Text, der dem Nutzer im Hauptformular angezeigt wird
  pay_later_text: z.string().nullable().optional(),
  // Der Beleg, der anstatt des normalen Belegtextes an den Benutzer geschickt wird
  pay_later_receipt: z.string().nullable().optional(),
  // is partial payment enabled for this online contribution page
  is_partial_payment: z.boolean().nullable().optional(),
  // Bezeichnung für den Erstbetrag bei Teilzahlung
  initial_amount_label: z.string().nullable().optional(),
  // Hilfetext zum Erstbetrag bei Teilzahlung
  initial_amount_help_text: z.string().nullable().optional(),
  // Minimalster Erstbetrag für Teilzahlung
  min_initial_amount: z.number().nullable().optional(),
  // if TRUE, page will include an input text field where user can enter their own amount
  is_allow_other_amount: z.boolean().optional(),
  // FK zu civicrm_option_value.
  default_amount_id: z.number().int().nullable().optional(),
  // if other amounts allowed, user can configure minimum allowed.
  min_amount: z.number().nullable().optional(),
  // if other amounts allowed, user can configure maximum allowed.
  max_amount: z.number().nullable().optional(),
  // The target goal for this page, allows people to build a goal meter
  goal_amount: z.number().nullable().optional(),
  // Title for Thank-you page (header title tag, and display at the top of the page).
  thankyou_title: z.string().nullable().optional(),
  // Text und HTML erlaubt. Wird über dem Ergebnis der Erfolgsseite angezeigt.
  thankyou_text: z.string().nullable().optional(),
  // Text and html allowed. displayed at the bottom of the success page. Common usage is to include link(s) to other pages such as tell-a-friend, etc.
  thankyou_footer: z.string().nullable().optional(),
  // if TRUE, receipt is automatically emailed to contact on success
  is_email_receipt: z.boolean().optional(),
  // FROM email name used for receipts generated by contributions to this contribution page.
  receipt_from_name: z.string().nullable().optional(),
  // FROM email address used for receipts generated by contributions to this contribution page.
  receipt_from_email: z.string().nullable().optional(),
  // comma-separated list of email addresses to cc each time a receipt is sent
  cc_receipt: z.string().nullable().optional(),
  // comma-separated list of email addresses to bcc each time a receipt is sent
  bcc_receipt: z.string().nullable().optional(),
  // text to include above standard receipt info on receipt email. emails are text-only, so do not allow html for now
  receipt_text: z.string().nullable().optional(),
  // Is this page active?
  is_active: z.boolean().optional(),
  // Text and html allowed. Displayed at the bottom of the first page of the contribution wizard.
  footer_text: z.string().nullable().optional(),
  // Ist diese Eigenschaft aktiv?
  amount_block_is_active: z.boolean().optional(),
  // Date and time that this page starts.
  start_date: z.string().nullable().optional(),
  // Date and time that this page ends. May be NULL if no defined end date/time
  end_date: z.string().nullable().optional(),
  // FK to civicrm_contact, who created this contribution page
  created_id: z.number().int().nullable().optional(),
  // Date and time that contribution page was created.
  created_date: z.string().optional(),
  // 3 Zeichen String, Wert aus den Konfig-Einstellungen oder aus der Benutzereingabe.
  currency: z.string().nullable().optional(),
  // Can people share the contribution page through social media?
  is_share: z.boolean().optional(),
  // if true - billing block is required for online contribution page
  is_billing_required: z.boolean().optional(),
})
export type ContributionPage = z.infer<typeof ContributionPageSchema>


// ────────────────────────────────────────────────────────
// ContributionProduct
// ────────────────────────────────────────────────────────

export const ContributionProductSchema = z.object({
  // Contribution Product ID
  id: z.number().int().optional(),
  // Product ID
  product_id: z.number().int(),
  // Zuwendungs-ID
  contribution_id: z.number().int(),
  // Option value selected if applicable - e.g. color, size etc.
  product_option: z.string().nullable().optional(),
  // Anzahl
  quantity: z.number().int().nullable().optional(),
  // Optional. Can be used to record the date this product was fulfilled or shipped.
  fulfilled_date: z.string().nullable().optional(),
  // Actual start date for a time-delimited premium (subscription, service or membership)
  start_date: z.string().nullable().optional(),
  // Actual end date for a time-delimited premium (subscription, service or membership)
  end_date: z.string().nullable().optional(),
  // Premium comment
  comment: z.string().nullable().optional(),
  // FK to Financial Type(for membership price sets only).
  financial_type_id: z.number().int().nullable().optional(),
})
export type ContributionProduct = z.infer<typeof ContributionProductSchema>


// ────────────────────────────────────────────────────────
// ContributionRecur
// ────────────────────────────────────────────────────────

export const ContributionRecurSchema = z.object({
  // Contribution Recur ID
  id: z.number().int().optional(),
  // Foreign key to civicrm_contact.id.
  contact_id: z.number().int(),
  // Amount to be collected (including any sales tax) by payment processor each recurrence.
  amount: z.number(),
  // 3 Zeichen String, Wert aus den Konfig-Einstellungen oder aus der Benutzereingabe.
  currency: z.string().nullable().optional(),
  // Time units for recurrence of payment.
  frequency_unit: z.string().nullable().optional(),
  // Number of time units for recurrence of payment.
  frequency_interval: z.number().int().optional(),
  // Gesamtzahl der Zahlungen. Auf 0 setzen, wenn diese auf unbestimmte Zeit fortgeführt werden sollen, d.h. ohne festes Enddatum.
  installments: z.number().int().nullable().optional(),
  // The date the first scheduled recurring contribution occurs.
  start_date: z.string().optional(),
  // When this recurring contribution record was created.
  create_date: z.string().optional(),
  // Last updated date for this record. mostly the last time a payment was received
  modified_date: z.string().optional(),
  // Date this recurring contribution was cancelled by contributor- if we can get access to it
  cancel_date: z.string().nullable().optional(),
  // Free text field for a reason for cancelling
  cancel_reason: z.string().nullable().optional(),
  // Date this recurring contribution finished successfully
  end_date: z.string().nullable().optional(),
  // Possibly needed to store a unique identifier for this recurring payment order - if this is available from the processor??
  processor_id: z.string().nullable().optional(),
  // Optionally used to store a link to a payment token used for this recurring contribution.
  payment_token_id: z.number().int().nullable().optional(),
  // unique transaction id (deprecated - use processor_id)
  trxn_id: z.string().nullable().optional(),
  // unique invoice id, system generated or passed in
  invoice_id: z.string().nullable().optional(),
  // Status
  contribution_status_id: z.number().int().nullable().optional(),
  // Test
  is_test: z.boolean().optional().default(false),
  // Day in the period when the payment should be charged e.g. 1st of month, 15th etc.
  cycle_day: z.number().int().optional(),
  // Next scheduled date
  next_sched_contribution_date: z.string().nullable().optional(),
  // Number of failed charge attempts since last success. Business rule could be set to deactivate on more than x failures.
  failure_count: z.number().int().nullable().optional(),
  // Date to retry failed attempt
  failure_retry_date: z.string().nullable().optional(),
  // Some systems allow contributor to set a number of installments - but then auto-renew the subscription or commitment if they do not cancel.
  auto_renew: z.boolean().optional(),
  // Foreign key to civicrm_payment_processor.id
  payment_processor_id: z.number().int().nullable().optional(),
  // FK to Financial Type
  financial_type_id: z.number().int().nullable().optional(),
  // FK to Payment Instrument
  payment_instrument_id: z.number().int().nullable().optional(),
  // if TRUE, receipt is automatically emailed to contact on each successful payment
  is_email_receipt: z.boolean().optional(),
})
export type ContributionRecur = z.infer<typeof ContributionRecurSchema>


// ────────────────────────────────────────────────────────
// ContributionSoft
// ────────────────────────────────────────────────────────

export const ContributionSoftSchema = z.object({
  // Soft Credit ID
  id: z.number().int().optional(),
  // FK to contribution table.
  contribution_id: z.number().int(),
  // FK zu Kontakt ID
  contact_id: z.number().int(),
  // Amount of this soft credit.
  amount: z.number(),
  // 3 Zeichen String, Wert aus den Konfig-Einstellungen oder aus der Benutzereingabe.
  currency: z.string().nullable().optional(),
  // FK to civicrm_pcp.id
  pcp_id: z.number().int().nullable().optional(),
  // Soft Contribution Display on PCP
  pcp_display_in_roll: z.boolean().optional(),
  // Soft Contribution PCP Nickname
  pcp_roll_nickname: z.string().nullable().optional(),
  // Soft Contribution PCP Note
  pcp_personal_note: z.string().nullable().optional(),
  // Soft Credit Type ID.Implicit FK to civicrm_option_value where option_group = soft_credit_type.
  soft_credit_type_id: z.number().int().nullable().optional(),
})
export type ContributionSoft = z.infer<typeof ContributionSoftSchema>


// ────────────────────────────────────────────────────────
// Country
// ────────────────────────────────────────────────────────

export const CountrySchema = z.object({
  // Länder-ID
  id: z.number().int().optional(),
  // Ländername
  name: z.string().nullable().optional(),
  // ISO Code
  iso_code: z.string().nullable().optional(),
  // Nationale Vorwahl, die benutzt werden muss um IN dieses Land anzurufen.
  country_code: z.string().nullable().optional(),
  // Foreign key to civicrm_address_format.id.
  address_format_id: z.number().int().nullable().optional(),
  // Internationale Vorwahl für Direktranrufe von einem Land ZU einem anderen Land
  idd_prefix: z.string().nullable().optional(),
  // Access prefix to call within a country to a different area
  ndd_prefix: z.string().nullable().optional(),
  // Foreign key to civicrm_worldregion.id.
  region_id: z.number().int(),
  // Soll Bundesland/Provinz als Abkürzung für Kontakte aus diesem Land angezeigt werden?
  is_province_abbreviated: z.boolean().optional(),
  // Ist dieses Land aktiv?
  is_active: z.boolean().optional(),
})
export type Country = z.infer<typeof CountrySchema>


// ────────────────────────────────────────────────────────
// County
// ────────────────────────────────────────────────────────

export const CountySchema = z.object({
  // Bezirks-ID
  id: z.number().int().optional(),
  // Name des Bezirks / Landkreises
  name: z.string().nullable().optional(),
  // 2-4 Zeichen Abkürzung des Landes
  abbreviation: z.string().nullable().optional(),
  // ID vom Bundesland/Provinz, zu dem der Landkreis gehört
  state_province_id: z.number().int(),
  // Ist dieser Bezirk / Landkreis aktiv?
  is_active: z.boolean().optional(),
})
export type County = z.infer<typeof CountySchema>


// ────────────────────────────────────────────────────────
// CustomField
// ────────────────────────────────────────────────────────

export const CustomFieldSchema = z.object({
  // Eindeutige ID Benutzerdefiniertes Feld
  id: z.number().int().optional(),
  // FK zu civicrm_custom_group.
  custom_group_id: z.number().int(),
  // Variablenname/programmatic handle für dieses Feld
  name: z.string(),
  // Text for form field label (also friendly name for administering this custom property).
  label: z.string(),
  // Controls location of data storage in extended_data table.
  data_type: z.string(),
  // HTML types plus several built-in extended types.
  html_type: z.string(),
  // Use form_options.is_default for field_types which use options.
  default_value: z.string().nullable().optional(),
  // Ist ein Wert für diese Eigenschaft notwendig.
  is_required: z.boolean().optional(),
  // Fügt einen Datenbankindex hinzu, der die Suche in diesem Feld erheblich beschleunigt. Allerdings kann dies mehr Speicherplatz erfordern und das System verlangsamen, wenn die Daten häufig aktualisiert werden.
  is_searchable: z.boolean().optional(),
  // Is this property range searchable.
  is_search_range: z.boolean().optional(),
  // Controls field display order within an extended property group.
  weight: z.number().int().optional(),
  // Description and/or help text to display before this field.
  help_pre: z.string().nullable().optional(),
  // Description and/or help text to display after this field.
  help_post: z.string().nullable().optional(),
  // Store collection of type-appropriate attributes, e.g. textarea needs rows/cols attributes
  attributes: z.string().nullable().optional(),
  // Ist diese Eigenschaft aktiv?
  is_active: z.boolean().nullable().optional(),
  // Ob dieses Feld durch PHP (über einen custom hook) gesetzt wird. Es wird nicht von CiviCRM aktualisiert.
  is_view: z.boolean().optional(),
  // number of options per line for checkbox and radio
  options_per_line: z.number().int().nullable().optional(),
  // Feldlänge, wenn alphanummerisch
  text_length: z.number().int().nullable().optional(),
  // Date may be up to start_date_years years prior to the current date.
  start_date_years: z.number().int().nullable().optional(),
  // Date may be up to end_date_years years after the current date.
  end_date_years: z.number().int().nullable().optional(),
  // Datumsformat für benutzerdefiniertes Datum
  date_format: z.string().nullable().optional(),
  // Uhrzeitformat für benutzerdefiniertes Datum
  time_format: z.number().int().nullable().optional(),
  // Number of columns in Note Field
  note_columns: z.number().int().nullable().optional(),
  // Number of rows in Note Field
  note_rows: z.number().int().nullable().optional(),
  // Name of the column that holds the values for this field.
  column_name: z.string().nullable().optional(),
  // For elements with options, the option group id that is used
  option_group_id: z.number().int().nullable().optional(),
  // Serialisierungsmethode - ein nicht-null Wert ist Indikator für ein Mehrere-Wert-Feld.
  serialize: z.number().int().optional(),
  // Stores Contact Get API params contact reference custom fields. May be used for other filters in the future.
  filter: z.string().nullable().optional(),
  // Should the multi-record custom field values be displayed in tab table listing
  in_selector: z.boolean().optional(),
  // Name of entity being referenced.
  fk_entity: z.string().nullable().optional(),
  // Behavior if referenced entity is deleted.
  fk_entity_on_delete: z.string().optional(),
  // Is provided by an extension
  has_base: z.boolean().nullable().optional(),
  // Name of extension which provides this package
  base_module: z.string().nullable().optional(),
  // When the managed entity was changed from its original settings
  local_modified_date: z.string().nullable().optional(),
})
export type CustomField = z.infer<typeof CustomFieldSchema>


// ────────────────────────────────────────────────────────
// CustomGroup
// ────────────────────────────────────────────────────────

export const CustomGroupSchema = z.object({
  // Unique Custom Group ID
  id: z.number().int().optional(),
  // Variable name/programmatic handle for this group.
  name: z.string(),
  // Friendly Name.
  title: z.string(),
  // Type of object this group extends (can add other options later e.g. contact_address, etc.).
  extends: z.string().optional(),
  // FK to civicrm_option_value.value (for option group custom_data_type)
  extends_entity_column_id: z.number().int().nullable().optional(),
  // linking custom group for dynamic object
  extends_entity_column_value: z.string().nullable().optional(),
  // Visual relationship between this form and its parent.
  style: z.string().optional(),
  // Will this group be in collapsed or expanded mode on initial display ?
  collapse_display: z.boolean().optional(),
  // Beschreibung und/oder Hilfetext, der vor diesem Feld auf dem Formular angezeigt wird.
  help_pre: z.string().nullable().optional(),
  // Description and/or help text to display after fields in form.
  help_post: z.string().nullable().optional(),
  // Controls display order when multiple extended property groups are setup for the same class.
  weight: z.number().int().optional(),
  // Ist diese Eigenschaft aktiv?
  is_active: z.boolean().optional(),
  // Name of the table that holds the values for this group.
  table_name: z.string().nullable().optional(),
  // Enthält die Gruppe mehrfache Werte?
  is_multiple: z.boolean().optional(),
  // Unused deprecated column.
  min_multiple: z.number().int().nullable().optional(),
  // maximum number of multiple records, if 0 - no max
  max_multiple: z.number().int().nullable().optional(),
  // Will this group be in collapsed or expanded mode on advanced search display ?
  collapse_adv_display: z.boolean().optional(),
  // FK to civicrm_contact, who created this custom group
  created_id: z.number().int().nullable().optional(),
  // Datum und Uhrzeit, als diese benutzerdefinierte Gruppe erstellt wurde.
  created_date: z.string().nullable().optional(),
  // Ist es eine reservierte Benutzerdefinierte Gruppe?
  is_reserved: z.boolean().optional(),
  // Ist diese Eigenschaft öffentlich?
  is_public: z.boolean().optional(),
  // crm-i icon class
  icon: z.string().nullable().optional(),
  // Is provided by an extension
  has_base: z.boolean().nullable().optional(),
  // Name of extension which provides this package
  base_module: z.string().nullable().optional(),
  // When the managed entity was changed from its original settings
  local_modified_date: z.string().nullable().optional(),
})
export type CustomGroup = z.infer<typeof CustomGroupSchema>


// ────────────────────────────────────────────────────────
// Dashboard
// ────────────────────────────────────────────────────────

export const DashboardSchema = z.object({
  // Dashlet-ID
  id: z.number().int().optional(),
  // Domain für Übersichtsseite
  domain_id: z.number().int(),
  // Interner Name des Dashlets.
  name: z.string().nullable().optional(),
  // Dashlet Titel
  label: z.string().nullable().optional(),
  // url in case of external dashlet
  url: z.string().nullable().optional(),
  // Berechtigungen für das Dashlet
  permission: z.string().nullable().optional(),
  // Rechte-Operator
  permission_operator: z.string().nullable().optional(),
  // Vollbild URL für Dashlet
  fullscreen_url: z.string().nullable().optional(),
  // Ist dieses Dashlet aktiv?
  is_active: z.boolean().optional(),
  // Ist dieses Dashlet reserviert?
  is_reserved: z.boolean().optional(),
  // Anzahl Minuten, für die Dashlet-Inhalte im Browser localStorage zwischengespeichert werden.
  cache_minutes: z.number().int().optional(),
  // Element name of angular directive to invoke (lowercase hyphenated format)
  directive: z.string().nullable().optional(),
  // Is provided by an extension
  has_base: z.boolean().nullable().optional(),
  // Name of extension which provides this package
  base_module: z.string().nullable().optional(),
  // When the managed entity was changed from its original settings
  local_modified_date: z.string().nullable().optional(),
})
export type Dashboard = z.infer<typeof DashboardSchema>


// ────────────────────────────────────────────────────────
// DashboardContact
// ────────────────────────────────────────────────────────

export const DashboardContactSchema = z.object({
  // Dashboard Kontakt-ID
  id: z.number().int().optional(),
  // Dashboard-ID
  dashboard_id: z.number().int(),
  // CiviCRM-ID
  contact_id: z.number().int(),
  // Spaltennr. für dieses Widget
  column_no: z.number().int().nullable().optional(),
  // Ist das Widget aktiv?
  is_active: z.boolean().optional(),
  // Sortierung der Widgets
  weight: z.number().int().nullable().optional(),
})
export type DashboardContact = z.infer<typeof DashboardContactSchema>


// ────────────────────────────────────────────────────────
// DedupeException
// ────────────────────────────────────────────────────────

export const DedupeExceptionSchema = z.object({
  // Unique dedupe exception id
  id: z.number().int().optional(),
  // FK zu Kontakt ID
  contact_id1: z.number().int(),
  // FK zu Kontakt ID
  contact_id2: z.number().int(),
})
export type DedupeException = z.infer<typeof DedupeExceptionSchema>


// ────────────────────────────────────────────────────────
// DedupeRule
// ────────────────────────────────────────────────────────

export const DedupeRuleSchema = z.object({
  // Unique dedupe rule id
  id: z.number().int().optional(),
  // The id of the rule group this rule belongs to
  dedupe_rule_group_id: z.number().int(),
  // The name of the table this rule is about
  rule_table: z.string(),
  // Der Name des Felds in der Tabelle, welche als rule_table referenziert ist
  rule_field: z.string(),
  // The length of the matching substring
  rule_length: z.number().int().nullable().optional(),
  // Die Gewichtung der Regel
  rule_weight: z.number().int(),
  // Is provided by an extension
  has_base: z.boolean().nullable().optional(),
  // Name of extension which provides this package
  base_module: z.string().nullable().optional(),
  // When the managed entity was changed from its original settings
  local_modified_date: z.string().nullable().optional(),
})
export type DedupeRule = z.infer<typeof DedupeRuleSchema>


// ────────────────────────────────────────────────────────
// DedupeRuleGroup
// ────────────────────────────────────────────────────────

export const DedupeRuleGroupSchema = z.object({
  // Unique dedupe rule group id
  id: z.number().int().optional(),
  // Die Kontaktart auf welche diese Gruppe zielt
  contact_type: z.string().nullable().optional(),
  // The weight threshold the sum of the rule weights has to cross to consider two contacts the same
  threshold: z.number().int(),
  // Whether the rule should be used for cases where usage is Unsupervised, Supervised OR General(programatically)
  used: z.string(),
  // Eindeutiger Name der Regelgruppe
  name: z.string().nullable().optional(),
  // Label der Regelgruppe
  title: z.string().nullable().optional(),
  // Is this a reserved rule - a rule group that has been optimized and cannot be changed by the admin
  is_reserved: z.boolean().optional(),
  // Is provided by an extension
  has_base: z.boolean().nullable().optional(),
  // Name of extension which provides this package
  base_module: z.string().nullable().optional(),
  // When the managed entity was changed from its original settings
  local_modified_date: z.string().nullable().optional(),
})
export type DedupeRuleGroup = z.infer<typeof DedupeRuleGroupSchema>


// ────────────────────────────────────────────────────────
// Discount
// ────────────────────────────────────────────────────────

export const DiscountSchema = z.object({
  // Primary Key
  id: z.number().int().optional(),
  // physical tablename for entity being joined to discount, e.g. civicrm_event
  entity_table: z.string(),
  // FK to entity table specified in entity_table column.
  entity_id: z.number().int(),
  // FK zu civicrm_price_set
  price_set_id: z.number().int(),
  // Datum, ab wann der Rabatt gilt.
  start_date: z.string().nullable().optional(),
  // Datum, ab wann der Rabatt abläuft.
  end_date: z.string().nullable().optional(),
})
export type Discount = z.infer<typeof DiscountSchema>


// ────────────────────────────────────────────────────────
// Domain
// ────────────────────────────────────────────────────────

export const DomainSchema = z.object({
  // Domain ID
  id: z.number().int().optional(),
  // Name der Domain / Organisation
  name: z.string().nullable().optional(),
  // Beschreibung der Domain.
  description: z.string().nullable().optional(),
  // Die CiviCRM-Version, die auf dieser Instanz läuft
  version: z.string().nullable().optional(),
  // FK to Contact ID. This is specifically not an FK to avoid circular constraints
  contact_id: z.number().int().nullable().optional(),
  // list of locales supported by the current db state (NULL for single-lang install)
  locales: z.string().nullable().optional(),
  // Locale specific string overrides
  locale_custom_strings: z.string().nullable().optional(),
  // Is this the current active domain
  is_active: z.boolean().nullable().optional(),
})
export type Domain = z.infer<typeof DomainSchema>


// ────────────────────────────────────────────────────────
// Email
// ────────────────────────────────────────────────────────

export const EmailSchema = z.object({
  // Eindeutige E-Mail-ID
  id: z.number().int().optional(),
  // FK zu Kontakt ID
  contact_id: z.number().int().nullable().optional(),
  // Zu welchem Adresskategorie diese E-Mail gehört.
  location_type_id: z.number().int().nullable().optional(),
  // E-Mail-Adresse
  email: z.string().nullable().optional(),
  // Ist dies die primäre E-Mail-Adresse des Kontakts
  is_primary: z.boolean().optional(),
  // Für Rechnungen benutzt?
  is_billing: z.boolean().optional(),
  // Implicit FK to civicrm_option_value where option_group = email_on_hold.
  on_hold: z.number().int().optional(),
  // Ist diese Adresse für Massenversand?
  is_bulkmail: z.boolean().optional(),
  // When the address went on bounce hold
  hold_date: z.string().nullable().optional(),
  // When the address bounce status was last reset
  reset_date: z.string().nullable().optional(),
  // Text-formatierte Signatur für diese E-Mail.
  signature_text: z.string().nullable().optional(),
  // HTML-formatierte Signatur für diese E-Mail
  signature_html: z.string().nullable().optional(),
})
export type Email = z.infer<typeof EmailSchema>


// ────────────────────────────────────────────────────────
// Entity
// ────────────────────────────────────────────────────────

export const EntitySchema = z.object({
  // Entity name
  name: z.string().nullable().optional(),
  // Localized title (singular)
  title: z.string().nullable().optional(),
  // Localized title (plural)
  title_plural: z.string().nullable().optional(),
  // Base class for this entity
  type: z.array(z.unknown()).nullable().optional(),
  // Description from docblock
  description: z.string().nullable().optional(),
  // Comments from docblock
  comment: z.string().nullable().optional(),
  // crm-i icon class associated with this entity
  icon: z.string().nullable().optional(),
  // Class name for dao-based entities
  dao: z.string().nullable().optional(),
  // Name of sql table, if applicable
  table_name: z.string().nullable().optional(),
  // Name of sql database, if different from CiviCRM
  database_name: z.string().nullable().optional(),
  // Name of unique identifier field(s) (e.g. [id])
  primary_key: z.array(z.unknown()).nullable().optional(),
  // Field to show when displaying a record
  label_field: z.string().nullable().optional(),
  // Fields to show in search context
  search_fields: z.array(z.unknown()).nullable().optional(),
  // Field(s) which contain the icon for a record, listed in order of precedence
  icon_field: z.array(z.unknown()).nullable().optional(),
  // Default column to sort results
  order_by: z.string().nullable().optional(),
  // Field linking a hierarchical entity to its parent
  parent_field: z.string().nullable().optional(),
  // How should this entity be presented in search UIs
  searchable: z.string().nullable().optional(),
  // System paths for accessing this entity
  paths: z.array(z.unknown()).nullable().optional(),
  // Any @see annotations from docblock
  see: z.array(z.unknown()).nullable().optional(),
  // Version this API entity was added
  since: z.string().nullable().optional(),
  // PHP class name
  class: z.string().nullable().optional(),
  // Arguments needed by php action factory functions (used when multiple entities share a class, e.g. CustomValue).
  class_args: z.array(z.unknown()).nullable().optional(),
  // Constant values which will be force-set when reading/writing this entity (e.g. [contact_type => Individual])
  where: z.array(z.unknown()).nullable().optional(),
  // Connecting fields for EntityBridge types
  bridge: z.array(z.unknown()).nullable().optional(),
  // When joining entities in the UI, which fields should be presented by default in the ON clause
  ui_join_filters: z.array(z.unknown()).nullable().optional(),
  // Combination of fields used for unique matching
  match_fields: z.array(z.unknown()).nullable().optional(),
  // For sortable entities, what field groupings are used to order by weight
  group_weights_by: z.array(z.unknown()).nullable().optional(),
})
export type Entity = z.infer<typeof EntitySchema>


// ────────────────────────────────────────────────────────
// EntityBatch
// ────────────────────────────────────────────────────────

export const EntityBatchSchema = z.object({
  // Primary Key
  id: z.number().int().optional(),
  // physical tablename for entity being joined to batch, e.g. civicrm_contact
  entity_table: z.string().nullable().optional(),
  // FK to entity table specified in entity_table column.
  entity_id: z.number().int(),
  // FK zu civicrm_batch
  batch_id: z.number().int(),
})
export type EntityBatch = z.infer<typeof EntityBatchSchema>


// ────────────────────────────────────────────────────────
// EntityFile
// ────────────────────────────────────────────────────────

export const EntityFileSchema = z.object({
  // Primary Key
  id: z.number().int().optional(),
  // physical tablename for entity being joined to file, e.g. civicrm_contact
  entity_table: z.string(),
  // FK to entity table specified in entity_table column.
  entity_id: z.number().int(),
  // FK to civicrm_file
  file_id: z.number().int(),
})
export type EntityFile = z.infer<typeof EntityFileSchema>


// ────────────────────────────────────────────────────────
// EntityFinancialAccount
// ────────────────────────────────────────────────────────

export const EntityFinancialAccountSchema = z.object({
  // ID
  id: z.number().int().optional(),
  // Links to an entity_table like civicrm_financial_type
  entity_table: z.string(),
  // Links to an id in the entity_table, such as vid in civicrm_financial_type
  entity_id: z.number().int(),
  // FK to a new civicrm_option_value (account_relationship)
  account_relationship: z.number().int(),
  // FK to the financial_account_id
  financial_account_id: z.number().int(),
})
export type EntityFinancialAccount = z.infer<typeof EntityFinancialAccountSchema>


// ────────────────────────────────────────────────────────
// EntityFinancialTrxn
// ────────────────────────────────────────────────────────

export const EntityFinancialTrxnSchema = z.object({
  // ID
  id: z.number().int().optional(),
  // May contain civicrm_financial_item, civicrm_contribution, civicrm_financial_trxn, civicrm_grant, etc
  entity_table: z.string(),
  // Entitäts-ID
  entity_id: z.number().int(),
  // Financial Transaction ID
  financial_trxn_id: z.number().int().nullable().optional(),
  // allocated amount of transaction to this entity
  amount: z.number(),
})
export type EntityFinancialTrxn = z.infer<typeof EntityFinancialTrxnSchema>


// ────────────────────────────────────────────────────────
// EntityTag
// ────────────────────────────────────────────────────────

export const EntityTagSchema = z.object({
  // Primary Key
  id: z.number().int().optional(),
  // physical tablename for entity being joined to file, e.g. civicrm_contact
  entity_table: z.string().nullable().optional(),
  // FK to entity table specified in entity_table column.
  entity_id: z.number().int(),
  // FK zu civicrm_tag
  tag_id: z.number().int(),
})
export type EntityTag = z.infer<typeof EntityTagSchema>


// ────────────────────────────────────────────────────────
// Event
// ────────────────────────────────────────────────────────

export const EventSchema = z.object({
  // Veranstaltung
  id: z.number().int().optional(),
  // Titel der Veranstaltung (z.B. Herbst-Spendengala)
  title: z.string().nullable().optional(),
  // Kurzbeschreibung der Veranstaltung. Eingabe von Reintext oder HTML möglich. Wird im Anmeldeformular angezeigt und kann auf anderen Seiten im CMS verwendet werden, die eine Veranstaltungszusammenfassung benötigen.
  summary: z.string().nullable().optional(),
  // Vollständige Veranstaltungsbeschreibung. Text und HTML erlaubt. Wird in den eingebauten Bereichen zur Veranstaltungsinformation angezeigt.
  description: z.string().nullable().optional(),
  // Event Type ID.Implicit FK to civicrm_option_value where option_group = event_type.
  event_type_id: z.number().int().nullable().optional(),
  // Sollen wir die Teilnehmerliste anzeigen? (Implicit FK to civicrm_option_value where option_group = participant_listing).
  participant_listing_id: z.number().int().nullable().optional(),
  // Öffentliche Veranstaltungen werden in den iCal-Feed eingebunden. Der Zugang zu nicht-öffentlichen, veranstaltungsbezogenen Informationen könnte von ACL-Regeln eingeschränkt sein. 
  is_public: z.boolean().optional(),
  // Datum und Uhrzeit, an dem die Veranstaltung beginnt.
  start_date: z.string().nullable().optional(),
  // Datum und Uhrzeit, an dem die Veranstaltung endet. Kann leer bleiben, wenn kein genaues Veranstaltungsende bekannt ist.
  end_date: z.string().nullable().optional(),
  // Falls ja, wird ein Registrierungslink auf der Seite mit den Veranstaltungsinformationen eingefügt.
  is_online_registration: z.boolean().optional(),
  // Text für den Link zum Veranstaltungsanmeldeformular, welcher auf der Seite der Veranstaltungsinformationen angezeigt wird, falls Online-Anmeldungen erlaubt sind.
  registration_link_text: z.string().nullable().optional(),
  // Startdatum und Uhrzeit der Online-Anmeldung.
  registration_start_date: z.string().nullable().optional(),
  // Startdatum und Uhrzeit der Online-Anmeldung.
  registration_end_date: z.string().nullable().optional(),
  // Maximale Teilnehmer:innenanzahl. Wenn das Limit erreicht wurde, erscheint eine entsprechende Nachricht. Das Feld leer lassen für unbegrenzte Teilnehmer:innen.
  max_participants: z.number().int().nullable().optional(),
  // Nachricht, die auf der Veranstaltungsinformationsseite und ANSTELLE des Veranstaltungsregistrierungsformulars angezeigt werden soll, sobald die maximale Teilnehmerzahl angemeldet ist. Kann E-Mail-Adresse/Informationen zur Aufnahme in eine Warteliste usw. enthalten. Text und HTML sind erlaubt. 
  event_full_text: z.string().nullable().optional(),
  // Falls ja, müssen ein oder mehrere Gebührenbeträge festgelegt und ein Zahlungsprozessor für die Online-Anmeldung konfiguriert sein. 
  is_monetary: z.boolean().optional(),
  // Zuwendungsart dieser Veranstaltung, die kostenpflichtigen Veranstaltunganmeldungen zugewiesen ist. Nur erforderlich, wenn die Veranstaltung kostenpflichtig ist.
  financial_type_id: z.number().int().nullable().optional(),
  // Für diese Veranstaltung konfigurierter Zahlungsprozessor (falls kostenpflichtig)
  payment_processor: z.string().nullable().optional(),
  // Füge einen Kartenblock auf der Veranstaltungsseite ein, wenn Geocodierung aktiviert ist und ein Mapping-Provider angegeben ist?
  is_map: z.boolean().optional(),
  // Ist die Veranstaltung aktiv oder inaktiv/abgebrochen?
  is_active: z.boolean().optional(),
  // Label der Gebühr
  fee_label: z.string().nullable().optional(),
  // Wenn ja, zeige den Veranstaltungsort.
  is_show_location: z.boolean().optional(),
  // FK to Location Block ID
  loc_block_id: z.number().int().nullable().optional(),
  // Teilnehmendenrolle ID. Impliziter FK zu civicrm_option_value where option_group = participant_role.
  default_role_id: z.number().int().nullable().optional(),
  // Einführungsnachricht für die Veranstaltungsanmeldeseite. Text und HTML erlaubt. Wird über dem Anmeldeformular angezeigt.
  intro_text: z.string().nullable().optional(),
  // Texte im Fußbereich der Anmeldeseite. Nur-Text und HTML sind erlaubt. Der Text wird am unteren Ende der Anmeldeseite eingeblendet.
  footer_text: z.string().nullable().optional(),
  // Titel der Bestätigungsseite.
  confirm_title: z.string().nullable().optional(),
  // Einführungsnachricht für die Veranstaltungsanmeldeseite. Text und HTML erlaubt. Wird über dem Anmeldeformular angezeigt.
  confirm_text: z.string().nullable().optional(),
  // Texte im Fußbereich der Anmeldeseite. Nur-Text und HTML sind erlaubt. Der Text wird am unteren Ende der Anmeldeseite eingeblendet.
  confirm_footer_text: z.string().nullable().optional(),
  // Wenn ja, wird automatisch eine Bestätigung an den erfolgreich angemeldeten Kontakt gesendet.  
  is_email_confirm: z.boolean().optional(),
  // Text, der in der Bescheinigungsmail oberhalb der Standardbescheinigung steht. Nur Text, bislang kein HTML
  confirm_email_text: z.string().nullable().optional(),
  // Anzeigename der E-Mail-Absenderadresse für die Anmeldebestätigung
  confirm_from_name: z.string().nullable().optional(),
  // E-Mail-Absenderadresse für die Anmeldebestätigung
  confirm_from_email: z.string().nullable().optional(),
  // komma-getrennte Liste von E-Mail-Adressen, an die jedes Mal eine Bestätigung in Kopie (CC) gesendet wird
  cc_confirm: z.string().nullable().optional(),
  // komma-getrennte Liste von E-Mail-Adressen, an die jedes Mal eine Bestätigung in Blindkopie (BCC) gesendet wird
  bcc_confirm: z.string().nullable().optional(),
  // FK zu civicrm_option_value.
  default_fee_id: z.number().int().nullable().optional(),
  // FK zu civicrm_option_value.
  default_discount_fee_id: z.number().int().nullable().optional(),
  // Titel der Dankeschön-Seite.
  thankyou_title: z.string().nullable().optional(),
  // Danke-Nachricht
  thankyou_text: z.string().nullable().optional(),
  // Text im Fußbereich.
  thankyou_footer_text: z.string().nullable().optional(),
  // falls aktiv - erlaubt dem Benutzer, die Zahlung später direkt an zu senden
  is_pay_later: z.boolean().optional(),
  // Der Text, der dem Nutzer im Hauptformular angezeigt wird
  pay_later_text: z.string().nullable().optional(),
  // Der Beleg, der anstatt des normalen Belegtextes an den Benutzer geschickt wird
  pay_later_receipt: z.string().nullable().optional(),
  // sind Teilzahlungen für die Veranstaltung möglich
  is_partial_payment: z.boolean().optional(),
  // Bezeichnung für den Erstbetrag bei Teilzahlung
  initial_amount_label: z.string().nullable().optional(),
  // Hilfetext zum Erstbetrag bei Teilzahlung
  initial_amount_help_text: z.string().nullable().optional(),
  // Minimalster Erstbetrag für Teilzahlung
  min_initial_amount: z.number().nullable().optional(),
  // Falls "ja", ist es einem Benutzer erlaubt, mehrere Teilnehmer:innen anzumelden
  is_multiple_registrations: z.boolean().optional(),
  // Maximale Anzahl von zusätzlichen Teilnehmern, die mit einer einzigen Anmeldung registriert werden können
  max_additional_participants: z.number().int().nullable().optional(),
  // Falls "ja", können mehrere Anmeldungen von derselben E-Mail-Adresse erfolgen.
  allow_same_participant_emails: z.boolean().optional(),
  // Ob die Veranstaltung eine Warteliste unterstützt.
  has_waitlist: z.boolean().optional(),
  // Ob Teilnehmer genehmigt werden müssen, bevor sie ihre Anmeldung abschließen können.
  requires_approval: z.boolean().optional(),
  // Lässt ausstehende, aber unbestätigte Anmeldungen nach dieser Zeit ablaufen. 
  expiration_time: z.number().int().nullable().optional(),
  // Stornierung oder Übertragung durch User zulassen?
  allow_selfcancelxfer: z.boolean().optional(),
  // Anzahl der Stunden vor dem Veranstaltungstart, innerhalb derer eine Stornierung oder Übertragung durch den User möglich ist. 
  selfcancelxfer_time: z.number().int().optional(),
  // Angezeigter Text, wenn die Veranstaltung zwar ausgebucht ist, aber Anmeldungen über eine Warteliste möglich sind.  
  waitlist_text: z.string().nullable().optional(),
  // Angezeigter Text, wenn eine Genehmigung erforderlich ist, um die Anmeldung für eine Veranstaltung abzuschließen.
  approval_req_text: z.string().nullable().optional(),
  // ob die Veranstaltung eine Vorlage hat
  is_template: z.boolean().optional().default(false),
  // Titel der Veranstaltungsvorlage
  template_title: z.string().nullable().optional(),
  // FK to civicrm_contact, wer die Veranstaltung angelegt hat
  created_id: z.number().int().nullable().optional(),
  // Datum und Uhrzeit, an dem die Veranstaltung erstellt wurde.
  created_date: z.string().nullable().optional(),
  // 3 Zeichen String, Wert aus den Konfig-Einstellungen oder aus der Benutzereingabe.
  currency: z.string().nullable().optional(),
  // Darf diese Veranstaltung auf Social-Media-Kanälen geteilt werden?
  is_share: z.boolean().optional(),
  // falls nicht, wird die Bestätigungsseite übersprungen
  is_confirm_enabled: z.boolean().optional(),
  // Implicit FK to civicrm_event: parent event
  parent_event_id: z.number().int().nullable().optional(),
  // Muss in die Event cart Erweiterung verschoben werden... Subevent slot label. Implicit FK to civicrm_option_value where option_group = conference_slot.
  slot_label_id: z.number().int().nullable().optional(),
  // Dubletten-Regel, die für Kontakte bei der Anmeldung verwendet wird
  dedupe_rule_group_id: z.number().int().nullable().optional(),
  // falls ja, dann wird ein Abrechnungsblock für diese Veranstaltung benötigt
  is_billing_required: z.boolean().optional(),
  // Falls ja, dann werden Kalender-Links für diese Veranstaltung angezeigt.
  is_show_calendar_links: z.boolean().optional(),
  // Is active with a non-past end-date
  is_current: z.boolean().nullable().optional(),
  // Maximale Teilnehmende minus angemeldete Teilnehmende
  remaining_participants: z.number().int().nullable().optional(),
})
export type Event = z.infer<typeof EventSchema>


// ────────────────────────────────────────────────────────
// ExampleData
// ────────────────────────────────────────────────────────

export const ExampleDataSchema = z.object({
  // Example Name
  name: z.string().nullable().optional(),
  // Example Title
  title: z.string().nullable().optional(),
  // Workflow Name
  workflow: z.string().nullable().optional(),
  // If the example is loaded from a file, this is the location.
  file: z.string().nullable().optional(),
  // Tags
  tags: z.array(z.unknown()).nullable().optional(),
  // Example data
  data: z.string().nullable().optional(),
  // Test assertions
  asserts: z.string().nullable().optional(),
})
export type ExampleData = z.infer<typeof ExampleDataSchema>


// ────────────────────────────────────────────────────────
// Extension
// ────────────────────────────────────────────────────────

export const ExtensionSchema = z.object({
  // Long, unique extension identifier
  key: z.string().nullable().optional(),
  // Short, unique extension identifier
  file: z.string().nullable().optional(),
  // User-facing extension title
  label: z.string().nullable().optional(),
  // Additional information about the extension
  description: z.string().nullable().optional(),
  // Current version number (string)
  version: z.string().nullable().optional(),
  // Tags which characterize the extension's purpose or functionality
  tags: z.array(z.unknown()).nullable().optional(),
  // Absolute file path
  path: z.string().nullable().optional(),
  // Release date
  releaseDate: z.string().nullable().optional(),
  // CiviCRM compatibility
  compatibility: z.string().nullable().optional(),
  // Development stage
  develStage: z.string().nullable().optional(),
  // URLs for extension page, documentation, licensing and support
  urls: z.array(z.unknown()).nullable().optional(),
  // Authors
  authors: z.array(z.unknown()).nullable().optional(),
  // License
  license: z.string().nullable().optional(),
  // Comments
  comments: z.string().nullable().optional(),
  // Extension enabled/disabled/uninstalled status
  status: z.string().nullable().optional(),
})
export type Extension = z.infer<typeof ExtensionSchema>


// ────────────────────────────────────────────────────────
// File
// ────────────────────────────────────────────────────────

export const FileSchema = z.object({
  // Eindeutige ID
  id: z.number().int().optional(),
  // Dateityp (z.B. Niederschrift, Einkommenssteuererstattung, usw.). FK to civicrm_option_value.
  file_type_id: z.number().int().nullable().optional(),
  // mime type of the document
  mime_type: z.string().nullable().optional(),
  // Location of file on disk relative to $config.customFileUploadDir
  uri: z.string().nullable().optional(),
  // Unused deprecated column.
  document: z.unknown().nullable().optional(),
  // Zusätzlicher Beschreibungstext für diesen Anhang (optional).
  description: z.string().nullable().optional(),
  // Datum und Uhrzeit als dieser Anhang hochgeladen oder auf den Server gespeichert wurde.
  upload_date: z.string().optional(),
  // FK to civicrm_contact, who uploaded this file
  created_id: z.number().int().nullable().optional(),
  // Filter by tags (including child tags)
  tags: z.array(z.unknown()).nullable().optional(),
  // Name of uploaded file
  file_name: z.string().nullable().optional(),
  // Url at which this file can be downloaded
  url: z.string().nullable().optional(),
  // Icon associated with this filetype
  icon: z.string().nullable().optional(),
  // Is this a recognized image type file
  is_image: z.boolean().nullable().optional(),
  // Contents of file
  content: z.string().nullable().optional(),
})
export type File = z.infer<typeof FileSchema>


// ────────────────────────────────────────────────────────
// FinancialAccount
// ────────────────────────────────────────────────────────

export const FinancialAccountSchema = z.object({
  // ID
  id: z.number().int().optional(),
  // Financial Account Name.
  name: z.string().optional(),
  // User-facing financial account label
  label: z.string().optional(),
  // FK to Contact ID that is responsible for the funds in this account
  contact_id: z.number().int().nullable().optional(),
  // pseudo FK into civicrm_option_value.
  financial_account_type_id: z.number().int().optional(),
  // Optional value for mapping monies owed and received to accounting system codes.
  accounting_code: z.string().nullable().optional(),
  // Optional value for mapping account types to accounting system account categories (QuickBooks Account Type Codes for example).
  account_type_code: z.string().nullable().optional(),
  // Financial Type Description.
  description: z.string().nullable().optional(),
  // Parent ID in account hierarchy
  parent_id: z.number().int().nullable().optional(),
  // Is this a header account which does not allow transactions to be posted against it directly, but only to its sub-accounts?
  is_header_account: z.boolean().optional(),
  // Is this account tax-deductible?
  is_deductible: z.boolean().optional(),
  // Is this account for taxes?
  is_tax: z.boolean().optional(),
  // The percentage of the total_amount that is due for this tax.
  tax_rate: z.number().nullable().optional(),
  // Is this a predefined system object?
  is_reserved: z.boolean().optional(),
  // Ist diese Eigenschaft aktiv?
  is_active: z.boolean().optional(),
  // Is this account the default one (or default tax one) for its financial_account_type?
  is_default: z.boolean().optional(),
})
export type FinancialAccount = z.infer<typeof FinancialAccountSchema>


// ────────────────────────────────────────────────────────
// FinancialItem
// ────────────────────────────────────────────────────────

export const FinancialItemSchema = z.object({
  // Financial Item ID
  id: z.number().int().optional(),
  // Date and time the item was created
  created_date: z.string().optional(),
  // Date and time of the source transaction
  transaction_date: z.string(),
  // FK to Contact ID of contact the item is from
  contact_id: z.number().int(),
  // Human readable description of this item, to ease display without lookup of source item.
  description: z.string().nullable().optional(),
  // Total amount of this item
  amount: z.number().optional(),
  // Currency for the amount
  currency: z.string().nullable().optional(),
  // FK to civicrm_financial_account
  financial_account_id: z.number().int().nullable().optional(),
  // Payment status: test, paid, part_paid, unpaid (if empty assume unpaid)
  status_id: z.number().int().nullable().optional(),
  // May contain civicrm_line_item, civicrm_financial_trxn etc
  entity_table: z.string().nullable().optional(),
  // The specific source item that is responsible for the creation of this financial_item
  entity_id: z.number().int().nullable().optional(),
})
export type FinancialItem = z.infer<typeof FinancialItemSchema>


// ────────────────────────────────────────────────────────
// FinancialTrxn
// ────────────────────────────────────────────────────────

export const FinancialTrxnSchema = z.object({
  // Financial Transaction ID
  id: z.number().int().optional(),
  // FK to financial_account table.
  from_financial_account_id: z.number().int().nullable().optional(),
  // FK to financial_financial_account table.
  to_financial_account_id: z.number().int().nullable().optional(),
  // date transaction occurred
  trxn_date: z.string().nullable().optional(),
  // amount of transaction
  total_amount: z.number(),
  // aktuelle Gebühr Zahlungsprozessor falls bekannt - kann 0 sein.
  fee_amount: z.number().nullable().optional(),
  // actual funds transfer amount. total less fees. if processor does not report actual fee during transaction, this is set to total_amount.
  net_amount: z.number().nullable().optional(),
  // 3 Zeichen String, Wert aus den Konfig-Einstellungen oder aus der Benutzereingabe.
  currency: z.string().nullable().optional(),
  // Is this entry either a payment or a reversal of a payment?
  is_payment: z.boolean().optional(),
  // Transaction id supplied by external processor. This may not be unique.
  trxn_id: z.string().nullable().optional(),
  // processor result code
  trxn_result_code: z.string().nullable().optional(),
  // pseudo FK to civicrm_option_value of contribution_status_id option_group
  status_id: z.number().int().nullable().optional(),
  // Payment Processor for this financial transaction
  payment_processor_id: z.number().int().nullable().optional(),
  // FK to payment_instrument option group values
  payment_instrument_id: z.number().int().nullable().optional(),
  // FK to accept_creditcard option group values
  card_type_id: z.number().int().nullable().optional(),
  // Check number
  check_number: z.string().nullable().optional(),
  // Last 4 digits of credit card
  pan_truncation: z.string().nullable().optional(),
  // Payment Processor external order reference
  order_reference: z.string().nullable().optional(),
})
export type FinancialTrxn = z.infer<typeof FinancialTrxnSchema>


// ────────────────────────────────────────────────────────
// FinancialType
// ────────────────────────────────────────────────────────

export const FinancialTypeSchema = z.object({
  // ID of original financial_type so you can search this table by the financial_type.id and then select the relevant version based on the timestamp
  id: z.number().int().optional(),
  // Financial Type Name.
  name: z.string().optional(),
  // User-facing financial type label
  label: z.string().optional(),
  // Financial Type Description.
  description: z.string().nullable().optional(),
  // Is this financial type tax-deductible? If TRUE, contributions of this type may be fully OR partially deductible - non-deductible amount is stored in the Contribution record.
  is_deductible: z.boolean().optional(),
  // Is this a predefined system object?
  is_reserved: z.boolean().optional(),
  // Ist diese Eigenschaft aktiv?
  is_active: z.boolean().optional(),
})
export type FinancialType = z.infer<typeof FinancialTypeSchema>


// ────────────────────────────────────────────────────────
// Group
// ────────────────────────────────────────────────────────

export const GroupSchema = z.object({
  // Gruppen-ID
  id: z.number().int().optional(),
  // Interner Name der Gruppe.
  name: z.string(),
  // Name der Gruppe.
  title: z.string().optional(),
  // Optional verbose description of the group.
  description: z.string().nullable().optional(),
  // Module or process which created this group.
  source: z.string().nullable().optional(),
  // FK to saved search table.
  saved_search_id: z.number().int().nullable().optional(),
  // Ist diese Gruppe aktiv?
  is_active: z.boolean().optional(),
  // In welchen Kontext(en) ist dieses Feld sichtbar?
  visibility: z.string().nullable().optional(),
  // the sql where clause if a saved search acl
  where_clause: z.string().nullable().optional(),
  // the tables to be included in a select data
  select_tables: z.string().nullable().optional(),
  // the tables to be included in the count statement
  where_tables: z.string().nullable().optional(),
  // FK zu Gruppentyp
  group_type: z.string().nullable().optional(),
  // Date when we created the cache for a smart group
  cache_date: z.string().nullable().optional(),
  // Seconds taken to fill smart group cache
  cache_fill_took: z.number().nullable().optional(),
  // Unused deprecated column.
  refresh_date: z.string().nullable().optional(),
  // Liste der Elterngruppen
  parents: z.string().nullable().optional(),
  // Liste der Untergruppen (berechnet)
  children: z.string().nullable().optional(),
  // Ist diese Gruppe ausgeblendet?
  is_hidden: z.boolean().optional(),
  // Gruppe ist reserviert
  is_reserved: z.boolean().optional(),
  // FK zur contact Tabelle.
  created_id: z.number().int().nullable().optional(),
  // FK zur contact Tabelle.
  modified_id: z.number().int().nullable().optional(),
  // Alternativer öffentlicher Titel dieser Gruppe
  frontend_title: z.string().optional(),
  // Alternative öffentliche Beschreibung der Gruppe
  frontend_description: z.string().nullable().optional(),
  // Number of contacts in group
  contact_count: z.number().int().nullable().optional(),
  // Is the smart group cache expired
  cache_expired: z.boolean().nullable().optional(),
  // Tiefe in der geschachtelten Hierarchie
  _depth: z.number().int().nullable().optional(),
  // Number of descendents in the nested hierarchy
  _descendents: z.number().int().nullable().optional(),
  // Is provided by an extension
  has_base: z.boolean().nullable().optional(),
  // Name of extension which provides this package
  base_module: z.string().nullable().optional(),
  // When the managed entity was changed from its original settings
  local_modified_date: z.string().nullable().optional(),
})
export type Group = z.infer<typeof GroupSchema>


// ────────────────────────────────────────────────────────
// GroupContact
// ────────────────────────────────────────────────────────

export const GroupContactSchema = z.object({
  // Primary Key
  id: z.number().int().optional(),
  // FK to civicrm_group
  group_id: z.number().int(),
  // FK zu civicrm_contact
  contact_id: z.number().int(),
  // status of contact relative to membership in group
  status: z.string().nullable().optional(),
  // Optional location to associate with this membership
  location_id: z.number().int().nullable().optional(),
  // Optional email to associate with this membership
  email_id: z.number().int().nullable().optional(),
})
export type GroupContact = z.infer<typeof GroupContactSchema>


// ────────────────────────────────────────────────────────
// GroupNesting
// ────────────────────────────────────────────────────────

export const GroupNestingSchema = z.object({
  // Beziehungs ID
  id: z.number().int().optional(),
  // ID der Untergruppe
  child_group_id: z.number().int(),
  // ID der übergeordneten "Eltern"-Gruppe
  parent_group_id: z.number().int(),
})
export type GroupNesting = z.infer<typeof GroupNestingSchema>


// ────────────────────────────────────────────────────────
// GroupOrganization
// ────────────────────────────────────────────────────────

export const GroupOrganizationSchema = z.object({
  // Beziehungs ID
  id: z.number().int().optional(),
  // ID der Gruppe
  group_id: z.number().int(),
  // ID des Organisationskontakts
  organization_id: z.number().int(),
})
export type GroupOrganization = z.infer<typeof GroupOrganizationSchema>


// ────────────────────────────────────────────────────────
// GroupSubscription
// ────────────────────────────────────────────────────────

export const GroupSubscriptionSchema = z.object({
  // CiviCRM-ID
  contact_id: z.number().int(),
  // Kontakten in dieser Gruppe werden die Berechtigungen der Administrator-Rolle zugewiesen.
  Administrators: z.boolean().nullable().optional().default(false),
  // Presseverteiler-Freiburg
  Presseverteiler_Frei_2: z.boolean().nullable().optional().default(false),
  // test
  _3: z.boolean().nullable().optional().default(false),
  // Contacts in this group are listed with their phone number and email when viewing case. You also can send copies of case activities to these contacts.
  Case_Resources: z.boolean().nullable().optional().default(false),
})
export type GroupSubscription = z.infer<typeof GroupSubscriptionSchema>


// ────────────────────────────────────────────────────────
// Household
// ────────────────────────────────────────────────────────

export const HouseholdSchema = z.object({
  // Eindeutige Kontakt-ID
  id: z.number().int().optional(),
  // Unique trusted external ID (generally from a legacy app/datasource). Particularly useful for deduping operations.
  external_identifier: z.string().nullable().optional(),
  // Formatierter Name, der das bevorzugte Format für Anzeige/Druck/andere Ausgaben angibt.
  display_name: z.string().nullable().optional(),
  // May be used to over-ride contact view and edit templates.
  contact_sub_type: z.string().nullable().optional(),
  // Keine E-Mails senden
  do_not_email: z.boolean().optional(),
  // Nicht anrufen
  do_not_phone: z.boolean().optional(),
  // Nicht anschreiben
  do_not_mail: z.boolean().optional(),
  // Keine SMS senden
  do_not_sms: z.boolean().optional(),
  // Nicht weitergeben
  do_not_trade: z.boolean().optional(),
  // Hat sich dieser Kontakt von allen Massen-E-Mails der Organisation bzw. der Webseitendomain abgemeldet, sogenanntes Opt-Out?
  is_opt_out: z.boolean().optional(),
  // May be used for SSN, EIN/TIN, Household ID (census) or other applicable unique legal/government ID.
  legal_identifier: z.string().nullable().optional(),
  // Name used for sorting different contact types
  sort_name: z.string().nullable().optional(),
  // Pseudonym.
  nick_name: z.string().nullable().optional(),
  // optionale URL für bevorzugtes Bild (Foto, Logo, etc.), welches für diesen Kontakt angezeigt wird.
  image_URL: z.string().nullable().optional(),
  // Was ist die bevorzugte Kommunikationsart.
  preferred_communication_method: z.string().nullable().optional(),
  // Which language is preferred for communication. FK to languages in civicrm_option_value.
  preferred_language: z.string().nullable().optional(),
  // Schlüssel zur Anfragenverifizierung bezogen auf diesen Kontakt.
  hash: z.string().nullable().optional(),
  // API-Schlüssel zur Anfragenverifizierung bezogen auf diesen Kontakt.
  api_key: z.string().nullable().optional(),
  // woher der Kontakt stammt, z.B. Import, Eintrag vom Spendenmodul...
  source: z.string().nullable().optional(),
  // Communication style (e.g. formal vs. familiar) to use with this contact. FK to communication styles in civicrm_option_value.
  communication_style_id: z.number().int().nullable().optional(),
  // FK to civicrm_option_value.value, that has to be valid registered Email Greeting.
  email_greeting_id: z.number().int().nullable().optional(),
  // Benutzerdefinierte E-Mail-Grußformel
  email_greeting_custom: z.string().nullable().optional(),
  // Cache Email Greeting.
  email_greeting_display: z.string().nullable().optional(),
  // FK to civicrm_option_value.value, that has to be valid registered Postal Greeting.
  postal_greeting_id: z.number().int().nullable().optional(),
  // Benutzerdefinierte Brief-Grußformel.
  postal_greeting_custom: z.string().nullable().optional(),
  // Cache Postal greeting.
  postal_greeting_display: z.string().nullable().optional(),
  // FK to civicrm_option_value.value, that has to be valid registered Addressee.
  addressee_id: z.number().int().nullable().optional(),
  // Benutzerdefinierter Adressat
  addressee_custom: z.string().nullable().optional(),
  // Cache Addressee.
  addressee_display: z.string().nullable().optional(),
  // Is Closed
  is_deceased: z.boolean().optional(),
  // Date closed or disbanded
  deceased_date: z.string().nullable().optional(),
  // Name des Haushalts
  household_name: z.string().nullable().optional(),
  // Optional FK to Primary Contact for this household.
  primary_contact_id: z.number().int().nullable().optional(),
  // the OpenID (or OpenID-style http://username.domain/) unique identifier for this contact mainly used for logging in to CiviCRM
  user_unique_id: z.string().nullable().optional(),
  // Kontakt ist im Papierkorb
  is_deleted: z.boolean().optional().default(false),
  // Wann der Kontakt erstellt wurde.
  created_date: z.string().nullable().optional(),
  // When was the contact (or closely related entity) was created or modified or deleted.
  modified_date: z.string().nullable().optional(),
  // Deprecated setting for text vs html mailings
  preferred_mail_format: z.string().nullable().optional(),
  // Primary Address ID
  address_primary: z.number().int().nullable().optional(),
  // Billing Address ID
  address_billing: z.number().int().nullable().optional(),
  // Primary Email ID
  email_primary: z.number().int().nullable().optional(),
  // Billing Email ID
  email_billing: z.number().int().nullable().optional(),
  // Primary Phone ID
  phone_primary: z.number().int().nullable().optional(),
  // Billing Phone ID
  phone_billing: z.number().int().nullable().optional(),
  // Primary IM ID
  im_primary: z.number().int().nullable().optional(),
  // Billing IM ID
  im_billing: z.number().int().nullable().optional(),
  // Groups (or sub-groups of groups) to which this contact belongs
  groups: z.array(z.unknown()).nullable().optional(),
  // Filter by tags (including child tags)
  tags: z.array(z.unknown()).nullable().optional(),
})
export type Household = z.infer<typeof HouseholdSchema>


// ────────────────────────────────────────────────────────
// IM
// ────────────────────────────────────────────────────────

export const IMSchema = z.object({
  // Eindeutige Sofortnachrichtendienst-ID
  id: z.number().int().optional(),
  // FK zu Kontakt ID
  contact_id: z.number().int().nullable().optional(),
  // Zu welchem Adresskategorie diese E-Mail gehört.
  location_type_id: z.number().int().nullable().optional(),
  // Benutzername des Sofortnachrichtendienstes
  name: z.string().nullable().optional(),
  // Zu welchem Sofortnachrichtendienstanbieter dieser Anzeigename gehört.
  provider_id: z.number().int().nullable().optional(),
  // Ist das der Sofortnachrichtendienst-Haupteintrag für diesen Kontakt und Ort?
  is_primary: z.boolean().optional(),
  // Für Rechnungen benutzt?
  is_billing: z.boolean().optional(),
})
export type IM = z.infer<typeof IMSchema>


// ────────────────────────────────────────────────────────
// Individual
// ────────────────────────────────────────────────────────

export const IndividualSchema = z.object({
  // Eindeutige Kontakt-ID
  id: z.number().int().optional(),
  // Unique trusted external ID (generally from a legacy app/datasource). Particularly useful for deduping operations.
  external_identifier: z.string().nullable().optional(),
  // Formatierter Name, der das bevorzugte Format für Anzeige/Druck/andere Ausgaben angibt.
  display_name: z.string().nullable().optional(),
  // May be used to over-ride contact view and edit templates.
  contact_sub_type: z.string().nullable().optional(),
  // Vorname
  first_name: z.string().nullable().optional(),
  // zweiter Vorname
  middle_name: z.string().nullable().optional(),
  // Nachname
  last_name: z.string().nullable().optional(),
  // Keine E-Mails senden
  do_not_email: z.boolean().optional(),
  // Nicht anrufen
  do_not_phone: z.boolean().optional(),
  // Nicht anschreiben
  do_not_mail: z.boolean().optional(),
  // Keine SMS senden
  do_not_sms: z.boolean().optional(),
  // Nicht weitergeben
  do_not_trade: z.boolean().optional(),
  // Hat sich dieser Kontakt von allen Massen-E-Mails der Organisation bzw. der Webseitendomain abgemeldet, sogenanntes Opt-Out?
  is_opt_out: z.boolean().optional(),
  // May be used for SSN, EIN/TIN, Household ID (census) or other applicable unique legal/government ID.
  legal_identifier: z.string().nullable().optional(),
  // Name used for sorting different contact types
  sort_name: z.string().nullable().optional(),
  // Pseudonym.
  nick_name: z.string().nullable().optional(),
  // optionale URL für bevorzugtes Bild (Foto, Logo, etc.), welches für diesen Kontakt angezeigt wird.
  image_URL: z.string().nullable().optional(),
  // Was ist die bevorzugte Kommunikationsart.
  preferred_communication_method: z.string().nullable().optional(),
  // Which language is preferred for communication. FK to languages in civicrm_option_value.
  preferred_language: z.string().nullable().optional(),
  // Schlüssel zur Anfragenverifizierung bezogen auf diesen Kontakt.
  hash: z.string().nullable().optional(),
  // API-Schlüssel zur Anfragenverifizierung bezogen auf diesen Kontakt.
  api_key: z.string().nullable().optional(),
  // woher der Kontakt stammt, z.B. Import, Eintrag vom Spendenmodul...
  source: z.string().nullable().optional(),
  // Prefix or Title for name (Ms, Mr...). FK to prefix ID
  prefix_id: z.number().int().nullable().optional(),
  // Suffix for name (Jr, Sr...). FK to suffix ID
  suffix_id: z.number().int().nullable().optional(),
  // Formeller (akademisch oder ähnlich) Titel vor dem Namen. (Prof., Dr. etc.)
  formal_title: z.string().nullable().optional(),
  // Communication style (e.g. formal vs. familiar) to use with this contact. FK to communication styles in civicrm_option_value.
  communication_style_id: z.number().int().nullable().optional(),
  // FK to civicrm_option_value.value, that has to be valid registered Email Greeting.
  email_greeting_id: z.number().int().nullable().optional(),
  // Benutzerdefinierte E-Mail-Grußformel
  email_greeting_custom: z.string().nullable().optional(),
  // Cache Email Greeting.
  email_greeting_display: z.string().nullable().optional(),
  // FK to civicrm_option_value.value, that has to be valid registered Postal Greeting.
  postal_greeting_id: z.number().int().nullable().optional(),
  // Benutzerdefinierte Brief-Grußformel.
  postal_greeting_custom: z.string().nullable().optional(),
  // Cache Postal greeting.
  postal_greeting_display: z.string().nullable().optional(),
  // FK to civicrm_option_value.value, that has to be valid registered Addressee.
  addressee_id: z.number().int().nullable().optional(),
  // Benutzerdefinierter Adressat
  addressee_custom: z.string().nullable().optional(),
  // Cache Addressee.
  addressee_display: z.string().nullable().optional(),
  // Funktion
  job_title: z.string().nullable().optional(),
  // FK to gender ID
  gender_id: z.number().int().nullable().optional(),
  // Geburtsdatum
  birth_date: z.string().nullable().optional(),
  // Ist verstorben
  is_deceased: z.boolean().optional(),
  // Date deceased
  deceased_date: z.string().nullable().optional(),
  // the OpenID (or OpenID-style http://username.domain/) unique identifier for this contact mainly used for logging in to CiviCRM
  user_unique_id: z.string().nullable().optional(),
  // OPTIONAL FK zu civicrm_contact record.
  employer_id: z.number().int().nullable().optional(),
  // Kontakt ist im Papierkorb
  is_deleted: z.boolean().optional().default(false),
  // Wann der Kontakt erstellt wurde.
  created_date: z.string().nullable().optional(),
  // When was the contact (or closely related entity) was created or modified or deleted.
  modified_date: z.string().nullable().optional(),
  // Deprecated setting for text vs html mailings
  preferred_mail_format: z.string().nullable().optional(),
  // Primary Address ID
  address_primary: z.number().int().nullable().optional(),
  // Billing Address ID
  address_billing: z.number().int().nullable().optional(),
  // Primary Email ID
  email_primary: z.number().int().nullable().optional(),
  // Billing Email ID
  email_billing: z.number().int().nullable().optional(),
  // Primary Phone ID
  phone_primary: z.number().int().nullable().optional(),
  // Billing Phone ID
  phone_billing: z.number().int().nullable().optional(),
  // Primary IM ID
  im_primary: z.number().int().nullable().optional(),
  // Billing IM ID
  im_billing: z.number().int().nullable().optional(),
  // Groups (or sub-groups of groups) to which this contact belongs
  groups: z.array(z.unknown()).nullable().optional(),
  // Age of individual (in years)
  age_years: z.number().int().nullable().optional(),
  // Number of days until next birthday
  next_birthday: z.number().int().nullable().optional(),
  // Filter by tags (including child tags)
  tags: z.array(z.unknown()).nullable().optional(),
})
export type Individual = z.infer<typeof IndividualSchema>


// ────────────────────────────────────────────────────────
// Job
// ────────────────────────────────────────────────────────

export const JobSchema = z.object({
  // Job ID
  id: z.number().int().optional(),
  // Which Domain is this scheduled job for
  domain_id: z.number().int(),
  // Lauffrequenz geplante Aufgabe
  run_frequency: z.string().nullable().optional(),
  // Wann dieser Cron Eintrag zuletzt gelaufen ist
  last_run: z.string().nullable().optional(),
  // Wann dieser Cron Eintrag zuletzt beendet wurde
  last_run_end: z.string().nullable().optional(),
  // When is this cron entry scheduled to run
  scheduled_run_date: z.string().nullable().optional(),
  // Titel der Aufgabe
  name: z.string().nullable().optional(),
  // Beschreibung dieser Aufgabe
  description: z.string().nullable().optional(),
  // Entity of the job api call
  api_entity: z.string().nullable().optional(),
  // Action of the job api call
  api_action: z.string().nullable().optional(),
  // Liste mit Parametern für den Befehl
  parameters: z.string().nullable().optional(),
  // Ist diese Aufgabe aktiv?
  is_active: z.boolean().optional(),
  // Is provided by an extension
  has_base: z.boolean().nullable().optional(),
  // Name of extension which provides this package
  base_module: z.string().nullable().optional(),
  // When the managed entity was changed from its original settings
  local_modified_date: z.string().nullable().optional(),
})
export type Job = z.infer<typeof JobSchema>


// ────────────────────────────────────────────────────────
// JobLog
// ────────────────────────────────────────────────────────

export const JobLogSchema = z.object({
  // Job log entry ID
  id: z.number().int().optional(),
  // Which Domain is this scheduled job for
  domain_id: z.number().int(),
  // Log entry date
  run_time: z.string().optional(),
  // Pointer to job id
  job_id: z.number().int().nullable().optional(),
  // Titel der Aufgabe
  name: z.string().nullable().optional(),
  // Full path to file containing job script
  command: z.string().nullable().optional(),
  // Title line of log entry
  description: z.string().nullable().optional(),
  // Potential extended data for specific job run (e.g. tracebacks).
  data: z.string().nullable().optional(),
})
export type JobLog = z.infer<typeof JobLogSchema>


// ────────────────────────────────────────────────────────
// LineItem
// ────────────────────────────────────────────────────────

export const LineItemSchema = z.object({
  // Line Item
  id: z.number().int().optional(),
  // May contain civicrm_contribution, civicrm_participant or civicrm_membership
  entity_table: z.string(),
  // entry in table
  entity_id: z.number().int(),
  // FK to civicrm_contribution
  contribution_id: z.number().int().nullable().optional(),
  // FK to civicrm_price_field
  price_field_id: z.number().int().nullable().optional(),
  // descriptive label for item - from price_field_value.label
  label: z.string().nullable().optional(),
  // How many items ordered
  qty: z.number(),
  // price of each item
  unit_price: z.number(),
  // qty * unit_price
  line_total: z.number(),
  // Participant count for field
  participant_count: z.number().int().nullable().optional(),
  // FK to civicrm_price_field_value
  price_field_value_id: z.number().int().nullable().optional(),
  // FK to Financial Type.
  financial_type_id: z.number().int().nullable().optional(),
  // Portion of total amount which is NOT tax deductible.
  non_deductible_amount: z.number().optional(),
  // tax of each item
  tax_amount: z.number().optional(),
  // Number of terms for this membership (only supported in Order->Payment flow). If the field is NULL it means unknown and it will be assumed to be 1 during payment.create if entity_table is civicrm_membership
  membership_num_terms: z.number().int().nullable().optional(),
})
export type LineItem = z.infer<typeof LineItemSchema>


// ────────────────────────────────────────────────────────
// LocBlock
// ────────────────────────────────────────────────────────

export const LocBlockSchema = z.object({
  // Eindeutige ID
  id: z.number().int().optional(),
  // Adress-ID
  address_id: z.number().int().nullable().optional(),
  // E-Mail-ID
  email_id: z.number().int().nullable().optional(),
  // Telefon-ID
  phone_id: z.number().int().nullable().optional(),
  // Sofortnachrichtendienst-ID
  im_id: z.number().int().nullable().optional(),
  // Adresse 2 ID
  address_2_id: z.number().int().nullable().optional(),
  // E-Mail 2 ID
  email_2_id: z.number().int().nullable().optional(),
  // Telefon 2 ID
  phone_2_id: z.number().int().nullable().optional(),
  // Sofortnachrichtendienst 2 ID
  im_2_id: z.number().int().nullable().optional(),
})
export type LocBlock = z.infer<typeof LocBlockSchema>


// ────────────────────────────────────────────────────────
// LocationType
// ────────────────────────────────────────────────────────

export const LocationTypeSchema = z.object({
  // Adresskategorie-ID
  id: z.number().int().optional(),
  // Adresskategorien-Name
  name: z.string(),
  // Adresskategorien-Anzeigename
  display_name: z.string(),
  // vCard Adresskategorien-Name.
  vcard_name: z.string().nullable().optional(),
  // Adresskategorie Beschreibung
  description: z.string().nullable().optional(),
  // Ist diese Adresskategorie ein vordefinierter Systemtyp?
  is_reserved: z.boolean().optional(),
  // Ist diese Eigenschaft aktiv?
  is_active: z.boolean().optional(),
  // Ist diese Adresskategorie der Standard?
  is_default: z.boolean().optional(),
})
export type LocationType = z.infer<typeof LocationTypeSchema>


// ────────────────────────────────────────────────────────
// Log
// ────────────────────────────────────────────────────────

export const LogSchema = z.object({
  // Log-ID
  id: z.number().int().optional(),
  // Name of table where item being referenced is stored.
  entity_table: z.string(),
  // Foreign key to the referenced item.
  entity_id: z.number().int(),
  // Updates does to this object if any.
  data: z.string().nullable().optional(),
  // FK to Contact ID of person under whose credentials this data modification was made.
  modified_id: z.number().int().nullable().optional(),
  // When was the referenced entity created or modified or deleted.
  modified_date: z.string().nullable().optional(),
})
export type Log = z.infer<typeof LogSchema>


// ────────────────────────────────────────────────────────
// MailSettings
// ────────────────────────────────────────────────────────

export const MailSettingsSchema = z.object({
  // Primary Key
  id: z.number().int().optional(),
  // Which Domain is this match entry for
  domain_id: z.number().int(),
  // Name dieser Gruppe von Einstellungen
  name: z.string().nullable().optional(),
  // whether this is the default set of settings for this domain
  is_default: z.boolean().optional(),
  // Domain der E-Mail-Adresse (der Teil hinter dem @)
  domain: z.string().nullable().optional(),
  // optional local part (like civimail+ for addresses like civimail+s.1.2@example.com)
  localpart: z.string().nullable().optional(),
  // contents of the Return-Path header
  return_path: z.string().nullable().optional(),
  // name of the protocol to use for polling (like IMAP, POP3 or Maildir)
  protocol: z.string().nullable().optional(),
  // Server, der zum Abstimmen benutzt wird
  server: z.string().nullable().optional(),
  // Port, der zum Abstimmen benutzt wird
  port: z.number().int().nullable().optional(),
  // Benutzername, der zum Abstimmen benutzt wird
  username: z.string().nullable().optional(),
  // Passwort, was zum Abstimmen benutzt wird
  password: z.string().nullable().optional(),
  // ob SSL benutzt wird oder nicht
  is_ssl: z.boolean().optional(),
  // folder to poll from when using IMAP, path to poll from when using Maildir, etc.
  source: z.string().nullable().optional(),
  // Name des Status, der benutzt wird um die E-Mail zu einer Aktivität umzuwandeln.
  activity_status: z.string().nullable().optional(),
  // Enabling this option will have CiviCRM skip any emails that do not have the Case ID or Case Hash so that the system will only process emails that can be placed on case records. Any emails that are not processed will be moved to the ignored folder.
  is_non_case_email_skipped: z.boolean().optional(),
  // If this option is enabled, CiviCRM will not create new contacts when filing emails.
  is_contact_creation_disabled_if_no_match: z.boolean().optional(),
  // Ignored for bounce processing, only for email-to-activity
  is_active: z.boolean().optional(),
  // Implicit FK to civicrm_option_value where option_group = activity_type
  activity_type_id: z.number().int().nullable().optional(),
  // Foreign key to the Campaign.
  campaign_id: z.number().int().nullable().optional(),
  // Which email recipient to add as the activity source (from, to, cc, bcc).
  activity_source: z.string().nullable().optional(),
  // Which email recipients to add as the activity targets (from, to, cc, bcc).
  activity_targets: z.string().nullable().optional(),
  // Which email recipients to add as the activity assignees (from, to, cc, bcc).
  activity_assignees: z.string().nullable().optional(),
})
export type MailSettings = z.infer<typeof MailSettingsSchema>


// ────────────────────────────────────────────────────────
// Mailing
// ────────────────────────────────────────────────────────

export const MailingSchema = z.object({
  // Newsletter ID
  id: z.number().int().optional(),
  // Für welche Domain das Rundschreiben ist
  domain_id: z.number().int().nullable().optional(),
  // FK für die Header-Komponente
  header_id: z.number().int().nullable().optional(),
  // FK für die Footer-Komponente
  footer_id: z.number().int().nullable().optional(),
  // FK to the auto-responder component.
  reply_id: z.number().int().nullable().optional(),
  // FK to the unsubscribe component.
  unsubscribe_id: z.number().int().nullable().optional(),
  // Newsletter wiederanmelden
  resubscribe_id: z.number().int().nullable().optional(),
  // FK to the opt-out component.
  optout_id: z.number().int().nullable().optional(),
  // Name des Rundschreibens.
  name: z.string().nullable().optional(),
  // differentiate between standalone mailings, A/B tests, and A/B final-winner
  mailing_type: z.string().nullable().optional(),
  // Das erste experimentelle Rundschreiben (Kondition "A")
  from_name: z.string().nullable().optional(),
  // From Email of mailing
  from_email: z.string().nullable().optional(),
  // Reply-To Email of mailing
  replyto_email: z.string().nullable().optional(),
  // Die Sprache/das Verarbeitungssystem, das für E-Mail-Vorlagen verwendet wird.
  template_type: z.string().optional(),
  // Advanced options used by the email templating system. (JSON encoded)
  template_options: z.string().nullable().optional(),
  // Betreff der E-Mail
  subject: z.string().nullable().optional(),
  // Body des Mailings im Text-Format.
  body_text: z.string().nullable().optional(),
  // Body des Mailings im HTML-Format.
  body_html: z.string().nullable().optional(),
  // Sollen die URL-Klickraten für dieses Rundschreiben verfolgt werden?
  url_tracking: z.boolean().optional(),
  // Sollen die Antworten an den Autor zurückgesendet werden?
  forward_replies: z.boolean().optional(),
  // Soll der Autoresponder aktiviert werden?
  auto_responder: z.boolean().optional(),
  // Soll nachverfolgt werden, wann die Empfänger dieses Rundschreiben öffnen/lesen?
  open_tracking: z.boolean().optional(),
  // Ist mindestens ein Auftrag im Zusammenhang mit diesem Rundschreiben abgeschlossen?
  is_completed: z.boolean().optional(),
  // FK zur Nachrichtenvorlage.
  msg_template_id: z.number().int().nullable().optional(),
  // Overwrite the VERP address in Reply-To
  override_verp: z.boolean().optional(),
  // FK to Contact ID who first created this mailing
  created_id: z.number().int().nullable().optional(),
  // Tag und Uhrzeit als dieses Mailing erstellt wurde.
  created_date: z.string().nullable().optional(),
  // Wann das Rundschreiben (oder eine eng verwandte Entität) erstellt oder geändert oder gelöscht wurde.
  modified_date: z.string().nullable().optional(),
  // FK to Contact ID who scheduled this mailing
  scheduled_id: z.number().int().nullable().optional(),
  // Tag und Uhrzeit für die dieses Mailing geplant wurde.
  scheduled_date: z.string().nullable().optional(),
  // When the mailing started to go out
  start_date: z.string().nullable().optional(),
  // When the mailing finished going out.
  end_date: z.string().nullable().optional(),
  // Rundschreibenstatus
  status: z.string().nullable().optional(),
  // FK to Contact ID who approved this mailing
  approver_id: z.number().int().nullable().optional(),
  // Tag und Uhrzeit als dieses Rundschreiben genehmigt wurde.
  approval_date: z.string().nullable().optional(),
  // Der Status dieses Rundschreibens. Werte: kein Status, genehmigt, abgelehnt
  approval_status_id: z.number().int().nullable().optional(),
  // Notiz zu dieser Entscheidung.
  approval_note: z.string().nullable().optional(),
  // Ist dieses Rundschreiben archiviert?
  is_archived: z.boolean().optional(),
  // In welchem(n) Kontext(en) ist der Inhalt des Rundschreibens sichtbar (Online-Ansicht)
  visibility: z.string().nullable().optional(),
  // Doppelte E-Mail-Adressen entfernen?
  dedupe_email: z.boolean().optional(),
  // ID des SMS-Anbieters
  sms_provider_id: z.number().int().nullable().optional(),
  // Key for validating requests related to this mailing.
  hash: z.string().nullable().optional(),
  // With email_selection_method, determines which email address to use
  location_type_id: z.number().int().nullable().optional(),
  // With location_type_id, determine how to choose the email address to use.
  email_selection_method: z.string().nullable().optional(),
  // Language of the content of the mailing. Useful for tokens.
  language: z.string().nullable().optional(),
  // One Click Unsubscribe mode either unsubscribe or opt-out
  unsubscribe_mode: z.string().optional(),
  // Total emails sent
  stats_intended_recipients: z.number().int().nullable().optional(),
  // Total emails delivered minus bounces
  stats_successful: z.number().int().nullable().optional(),
  // Total tracked mailing opens
  stats_opens_total: z.number().int().nullable().optional(),
  // Total unique tracked mailing opens
  stats_opens_unique: z.number().int().nullable().optional(),
  // Total mailing clicks
  stats_clicks_total: z.number().int().nullable().optional(),
  // Total unique mailing clicks
  stats_clicks_unique: z.number().int().nullable().optional(),
  // Total mailing bounces
  stats_bounces: z.number().int().nullable().optional(),
  // Total mailing unsubscribes
  stats_unsubscribes: z.number().int().nullable().optional(),
  // Total mailing opt outs
  stats_optouts: z.number().int().nullable().optional(),
  // Total contacts who opted out or unsubscribed from a mailing
  stats_optouts_and_unsubscribes: z.number().int().nullable().optional(),
  // Total mailing replies
  stats_replies: z.number().int().nullable().optional(),
})
export type Mailing = z.infer<typeof MailingSchema>


// ────────────────────────────────────────────────────────
// MailingComponent
// ────────────────────────────────────────────────────────

export const MailingComponentSchema = z.object({
  // Newsletter Elemente Nr.
  id: z.number().int().optional(),
  // Der Name dieser Komponente
  name: z.string().nullable().optional(),
  // Type of Component.
  component_type: z.string().nullable().optional(),
  // Thema, Betreff
  subject: z.string().nullable().optional(),
  // Body of the component in html format.
  body_html: z.string().nullable().optional(),
  // Body of the component in text format.
  body_text: z.string().nullable().optional(),
  // Is this the default component for this component_type?
  is_default: z.boolean().optional(),
  // Ist diese Eigenschaft aktiv?
  is_active: z.boolean().optional(),
})
export type MailingComponent = z.infer<typeof MailingComponentSchema>


// ────────────────────────────────────────────────────────
// MailingEventBounce
// ────────────────────────────────────────────────────────

export const MailingEventBounceSchema = z.object({
  // Bounce-ID
  id: z.number().int().optional(),
  // FK to EventQueue
  event_queue_id: z.number().int(),
  // Was für ein Bounce-Typ war es?
  bounce_type_id: z.number().int().nullable().optional(),
  // The reason the email bounced.
  bounce_reason: z.string().nullable().optional(),
  // When this bounce event occurred.
  time_stamp: z.string().optional(),
})
export type MailingEventBounce = z.infer<typeof MailingEventBounceSchema>


// ────────────────────────────────────────────────────────
// MailingEventConfirm
// ────────────────────────────────────────────────────────

export const MailingEventConfirmSchema = z.object({
  // Newsletter Bestätigungs-ID
  id: z.number().int().optional(),
  // FK to civicrm_mailing_event_subscribe
  event_subscribe_id: z.number().int(),
  // Wann die Bestätigung durchgeführt wurde.
  time_stamp: z.string().optional(),
})
export type MailingEventConfirm = z.infer<typeof MailingEventConfirmSchema>


// ────────────────────────────────────────────────────────
// MailingEventDelivered
// ────────────────────────────────────────────────────────

export const MailingEventDeliveredSchema = z.object({
  // Zugestellt ID
  id: z.number().int().optional(),
  // FK to EventQueue
  event_queue_id: z.number().int(),
  // When this delivery event occurred.
  time_stamp: z.string().optional(),
})
export type MailingEventDelivered = z.infer<typeof MailingEventDeliveredSchema>


// ────────────────────────────────────────────────────────
// MailingEventOpened
// ────────────────────────────────────────────────────────

export const MailingEventOpenedSchema = z.object({
  // Mailing Opened ID
  id: z.number().int().optional(),
  // FK to EventQueue
  event_queue_id: z.number().int(),
  // When this open event occurred.
  time_stamp: z.string().optional(),
})
export type MailingEventOpened = z.infer<typeof MailingEventOpenedSchema>


// ────────────────────────────────────────────────────────
// MailingEventQueue
// ────────────────────────────────────────────────────────

export const MailingEventQueueSchema = z.object({
  // Newsletter Veranstaltungsreihe Nr.
  id: z.number().int().optional(),
  // Mailing Job
  job_id: z.number().int().nullable().optional(),
  // Related mailing. Used for reporting on mailing success, if present.
  mailing_id: z.number().int().nullable().optional(),
  // Test
  is_test: z.boolean().optional().default(false),
  // FK to Email
  email_id: z.number().int().nullable().optional(),
  // FK zu Kontakt
  contact_id: z.number().int(),
  // Security hash
  hash: z.string(),
  // FK to Phone
  phone_id: z.number().int().nullable().optional(),
})
export type MailingEventQueue = z.infer<typeof MailingEventQueueSchema>


// ────────────────────────────────────────────────────────
// MailingEventReply
// ────────────────────────────────────────────────────────

export const MailingEventReplySchema = z.object({
  // Reply-ID
  id: z.number().int().optional(),
  // FK to EventQueue
  event_queue_id: z.number().int(),
  // When this reply event occurred.
  time_stamp: z.string().optional(),
})
export type MailingEventReply = z.infer<typeof MailingEventReplySchema>


// ────────────────────────────────────────────────────────
// MailingEventSubscribe
// ────────────────────────────────────────────────────────

export const MailingEventSubscribeSchema = z.object({
  // Newsletter Anmeldung Nr.
  id: z.number().int().optional(),
  // FK zur Gruppe
  group_id: z.number().int(),
  // FK zu Kontakt
  contact_id: z.number().int(),
  // Security hash
  hash: z.string(),
  // When this subscription event occurred.
  time_stamp: z.string().optional(),
})
export type MailingEventSubscribe = z.infer<typeof MailingEventSubscribeSchema>


// ────────────────────────────────────────────────────────
// MailingEventTrackableURLOpen
// ────────────────────────────────────────────────────────

export const MailingEventTrackableURLOpenSchema = z.object({
  // Trackable URL Open ID
  id: z.number().int().optional(),
  // FK to EventQueue
  event_queue_id: z.number().int(),
  // FK zu TrackableURL
  trackable_url_id: z.number().int(),
  // When this trackable URL open occurred.
  time_stamp: z.string().optional(),
})
export type MailingEventTrackableURLOpen = z.infer<typeof MailingEventTrackableURLOpenSchema>


// ────────────────────────────────────────────────────────
// MailingEventUnsubscribe
// ────────────────────────────────────────────────────────

export const MailingEventUnsubscribeSchema = z.object({
  // Abmeldung Nr.
  id: z.number().int().optional(),
  // FK to EventQueue
  event_queue_id: z.number().int(),
  // Abbestellung auf org- oder Gruppen-Bereich
  org_unsubscribe: z.boolean(),
  // When this delivery event occurred.
  time_stamp: z.string().optional(),
})
export type MailingEventUnsubscribe = z.infer<typeof MailingEventUnsubscribeSchema>


// ────────────────────────────────────────────────────────
// MailingGroup
// ────────────────────────────────────────────────────────

export const MailingGroupSchema = z.object({
  // Newsletter Gruppe Nr.
  id: z.number().int().optional(),
  // Die ID eines früheren Rundschreibens zum Ein-/Ausschließen von Empfängern.
  mailing_id: z.number().int(),
  // Sind die Gruppenmitglieder ein- oder ausgeschlossen?
  group_type: z.string().nullable().optional(),
  // Name of table where item being referenced is stored.
  entity_table: z.string(),
  // Foreign key to the referenced item.
  entity_id: z.number().int(),
  // The filtering search. custom search id or -1 for civicrm api search
  search_id: z.number().int().nullable().optional(),
  // The arguments to be sent to the search function
  search_args: z.string().nullable().optional(),
})
export type MailingGroup = z.infer<typeof MailingGroupSchema>


// ────────────────────────────────────────────────────────
// MailingJob
// ────────────────────────────────────────────────────────

export const MailingJobSchema = z.object({
  // Newsletter Bearbeitung Nr.
  id: z.number().int().optional(),
  // The ID of the mailing this Job will send.
  mailing_id: z.number().int(),
  // date on which this job was scheduled.
  scheduled_date: z.string().nullable().optional(),
  // Die Aufgabe wurde begonnen am:
  start_date: z.string().nullable().optional(),
  // Die Aufgabe wurde beendet am:
  end_date: z.string().nullable().optional(),
  // Der Status dieser Aufgabe
  status: z.string().nullable().optional(),
  // Ist diese Aufgabe für eine Testmail?
  is_test: z.boolean().optional().default(false),
  // Type of mailling job: null | child
  job_type: z.string().nullable().optional(),
  // Parent job id
  parent_id: z.number().int().nullable().optional(),
  // Offset of the child job
  job_offset: z.number().int().nullable().optional(),
  // Queue size limit for each child job
  job_limit: z.number().int().nullable().optional(),
})
export type MailingJob = z.infer<typeof MailingJobSchema>


// ────────────────────────────────────────────────────────
// MailingTrackableURL
// ────────────────────────────────────────────────────────

export const MailingTrackableURLSchema = z.object({
  // Verfolgbare URL ID
  id: z.number().int().optional(),
  // URL, die getrackt werden soll.
  url: z.string(),
  // FK to the mailing
  mailing_id: z.number().int(),
})
export type MailingTrackableURL = z.infer<typeof MailingTrackableURLSchema>


// ────────────────────────────────────────────────────────
// Managed
// ────────────────────────────────────────────────────────

export const ManagedSchema = z.object({
  // Surrogate Key
  id: z.number().int().optional(),
  // Name of the module which declared this object (soft FK to civicrm_extension.full_name)
  module: z.string(),
  // Symbolic name used by the module to identify the object
  name: z.string(),
  // API-Entitätstyp
  entity_type: z.string(),
  // Soft foreign key to the referenced item.
  entity_id: z.number().int().nullable().optional(),
  // Configuration of the managed-entity when last stored
  checksum: z.string().nullable().optional(),
  // Policy on when to cleanup entity (always, never, unused)
  cleanup: z.string().optional(),
  // When the managed entity was changed from its original settings.
  entity_modified_date: z.string().nullable().optional(),
})
export type Managed = z.infer<typeof ManagedSchema>


// ────────────────────────────────────────────────────────
// Mapping
// ────────────────────────────────────────────────────────

export const MappingSchema = z.object({
  // Mapping ID
  id: z.number().int().optional(),
  // Eindeutiger Name der Feldzuordnung
  name: z.string(),
  // Beschreibung der Feldzuordnung.
  description: z.string().nullable().optional(),
  // Art der Feldzuordnung
  mapping_type_id: z.number().int().nullable().optional(),
})
export type Mapping = z.infer<typeof MappingSchema>


// ────────────────────────────────────────────────────────
// MappingField
// ────────────────────────────────────────────────────────

export const MappingFieldSchema = z.object({
  // Feldzuordnung-ID
  id: z.number().int().optional(),
  // Feldzuordnung zu der dieses Feld gehört
  mapping_id: z.number().int(),
  // Feldzuordnung Schlüssel
  name: z.string().nullable().optional(),
  // Kontaktart in Feldzuordnung
  contact_type: z.string().nullable().optional(),
  // Column number for mapping set
  column_number: z.number().int(),
  // Adresskategorie für diese Zuordnung / Mapping, falls benötigt
  location_type_id: z.number().int().nullable().optional(),
  // Zu welchem Telefontyp diese Rufnummer gehört.
  phone_type_id: z.number().int().nullable().optional(),
  // Zu welchem Sofortnachrichtendienst dieser Name gehört.
  im_provider_id: z.number().int().nullable().optional(),
  // Zu welchem Webseitenart diese Seite gehört
  website_type_id: z.number().int().nullable().optional(),
  // Beziehungstyp, falls erforderlich
  relationship_type_id: z.number().int().nullable().optional(),
  // Richtung der Beziehung
  relationship_direction: z.string().nullable().optional(),
  // Used to group mapping_field records into related sets (e.g. for criteria sets in search builder mappings).
  grouping: z.number().int().nullable().optional(),
  // SQL WHERE operator for search-builder mapping fields (search criteria).
  operator: z.string().nullable().optional(),
  // SQL WHERE value for search-builder mapping fields.
  value: z.string().nullable().optional(),
})
export type MappingField = z.infer<typeof MappingFieldSchema>


// ────────────────────────────────────────────────────────
// Membership
// ────────────────────────────────────────────────────────

export const MembershipSchema = z.object({
  // Mitgliedschafts ID
  id: z.number().int().optional(),
  // FK zu Kontakt ID
  contact_id: z.number().int(),
  // FK zum Mitgliedstyp
  membership_type_id: z.number().int(),
  // Zeitpunkt der initialen Mitgliedschaft
  join_date: z.string().nullable().optional(),
  // Zeitpunkt, an dem die derzeitige (seitdem ununterbrochene) Mitgliedschaft begann. 
  start_date: z.string().nullable().optional(),
  // Aktuelle Mitgliedschaft Ablaufdatum
  end_date: z.string().nullable().optional(),
  // Mitgliedschaft Bezugsquelle
  source: z.string().nullable().optional(),
  // FK zu Mitgliedsstatus
  status_id: z.number().int(),
  // Admin users may set a manual status which overrides the calculated status. When this flag is TRUE, automated status update scripts should NOT modify status for the record.
  is_override: z.boolean().optional(),
  // Dann wird das Enddatum des Mitgliedschaftsstatus überschrieben, wenn "Überschreiben bis zum gewählten Datum" ausgewählt ist.
  status_override_end_date: z.string().nullable().optional(),
  // Optionaler FK zur Eltern-Mitgliedschaft.
  owner_membership_id: z.number().int().nullable().optional(),
  // Maximale Anzahl "vererbbarer" Mitgliedschaften (membership_type override).
  max_related: z.number().int().nullable().optional(),
  // Test
  is_test: z.boolean().optional().default(false),
  // Ist "Später zahlen"
  is_pay_later: z.boolean().optional(),
  // Conditional foreign key to civicrm_contribution_recur id. Each membership in connection with a recurring contribution carries a foreign key to the recurring contribution record. This assumes we can track these processor initiated events.
  contribution_recur_id: z.number().int().nullable().optional(),
  // Ist dies eine Primärmitgliedschaft?
  is_primary_member: z.boolean().nullable().optional(),
})
export type Membership = z.infer<typeof MembershipSchema>


// ────────────────────────────────────────────────────────
// MembershipBlock
// ────────────────────────────────────────────────────────

export const MembershipBlockSchema = z.object({
  // Mitgliedschafts ID
  id: z.number().int().optional(),
  // Name des Mitgliedschafts-Status
  entity_table: z.string().nullable().optional(),
  // FK zu civicrm_contribution_page.id
  entity_id: z.number().int(),
  // Membership types to be exposed by this block
  membership_types: z.string().nullable().optional(),
  // Optionaler Fremdschlüssel zur Mitgliedsart
  membership_type_default: z.number().int().nullable().optional(),
  // Anzeige des Mindestbeitrags
  display_min_fee: z.boolean().optional(),
  // Sollen Mitgliedschaftstransaktionen separat verarbeitet werden
  is_separate_payment: z.boolean().optional(),
  // Titel zur Anzeige oberhalb des Blocks
  new_title: z.string().nullable().optional(),
  // Text zur Anzeige unterhalb des Titels
  new_text: z.string().nullable().optional(),
  // Titel für Verlängerung
  renewal_title: z.string().nullable().optional(),
  // Anzeigetext für Verlängerung der Mitgliedschaft
  renewal_text: z.string().nullable().optional(),
  // Ist die Eintragung einer Mitgliedschaft optional
  is_required: z.boolean().optional(),
  // Ist dieser membership_block aktiviert
  is_active: z.boolean().optional(),
})
export type MembershipBlock = z.infer<typeof MembershipBlockSchema>


// ────────────────────────────────────────────────────────
// MembershipLog
// ────────────────────────────────────────────────────────

export const MembershipLogSchema = z.object({
  // Mitgliedschaftslog-ID
  id: z.number().int().optional(),
  // FK zur Mitgliedschaftstabelle
  membership_id: z.number().int(),
  // Durch diese Aktion wird ein neuer Status zur Mitgliedschaft zugeweisen. FK zu Mitgliedsstatus
  status_id: z.number().int(),
  // New membership period start date
  start_date: z.string().nullable().optional(),
  // Ablaufdatum der neuen Mitgliedschaft
  end_date: z.string().nullable().optional(),
  // FK to Contact ID of person under whose credentials this data modification was made.
  modified_id: z.number().int().nullable().optional(),
  // Datum, wann die Veränderung der Mitgliedschaft geloggt wurde.
  modified_date: z.string().nullable().optional(),
  // FK zur Mitgliedschaftsart
  membership_type_id: z.number().int().nullable().optional(),
  // Maximale Anzahl der zugehörigen Mitgliedschaften.
  max_related: z.number().int().nullable().optional(),
})
export type MembershipLog = z.infer<typeof MembershipLogSchema>


// ────────────────────────────────────────────────────────
// MembershipStatus
// ────────────────────────────────────────────────────────

export const MembershipStatusSchema = z.object({
  // Mitgliedschafts ID
  id: z.number().int().optional(),
  // Name des Mitgliedschafts-Status
  name: z.string(),
  // Name für den Mitgliedschafts-Status
  label: z.string().nullable().optional(),
  // Ereignis, wann dieser Status beginnt.
  start_event: z.string().nullable().optional(),
  // Einheit, die zur Anpassung des Startereignisses benutzt wird.
  start_event_adjust_unit: z.string().nullable().optional(),
  // Status range begins this many units from start_event.
  start_event_adjust_interval: z.number().int().nullable().optional(),
  // Ereignis, wann dieser Status endet.
  end_event: z.string().nullable().optional(),
  // Einheit, die zur Anpassung des Endereignisses benutzt wird.
  end_event_adjust_unit: z.string().nullable().optional(),
  // Status range ends this many units from end_event.
  end_event_adjust_interval: z.number().int().nullable().optional(),
  // Does this status aggregate to current members (e.g. New, Renewed, Grace might all be TRUE... while Unrenewed, Lapsed, Inactive would be FALSE).
  is_current_member: z.boolean().optional(),
  // Ist dieser Status nur für Administration / manuelle Zuordnung.
  is_admin: z.boolean().optional(),
  // Reihenfolge
  weight: z.number().int().nullable().optional(),
  // Assign this status to a membership record if no other status match is found.
  is_default: z.boolean().optional(),
  // Ist dieser membership_status aktiviert.
  is_active: z.boolean().optional(),
  // Ist dieser membership_status reserviert.
  is_reserved: z.boolean().optional(),
  // Is provided by an extension
  has_base: z.boolean().nullable().optional(),
  // Name of extension which provides this package
  base_module: z.string().nullable().optional(),
  // When the managed entity was changed from its original settings
  local_modified_date: z.string().nullable().optional(),
  // Ist das der Status für neue Mitglieder
  is_new: z.boolean().nullable().optional(),
})
export type MembershipStatus = z.infer<typeof MembershipStatusSchema>


// ────────────────────────────────────────────────────────
// MembershipType
// ────────────────────────────────────────────────────────

export const MembershipTypeSchema = z.object({
  // Mitgliedschafts ID
  id: z.number().int().optional(),
  // Which Domain is this match entry for
  domain_id: z.number().int(),
  // Mitgliedsart Name
  name: z.string(),
  // Title of Membership Type when shown to CiviCRM administrators.
  title: z.string().optional(),
  // Title of Membership Type when shown on public pages etc.
  frontend_title: z.string().optional(),
  // Beschreibung der Mitgliedstypen
  description: z.string().nullable().optional(),
  // Owner organization for this membership type. FK to Contact ID
  member_of_contact_id: z.number().int(),
  // If membership is paid by a contribution - what financial type should be used. FK to civicrm_financial_type.id
  financial_type_id: z.number().int(),
  // Mindestbeitrag für diesen Mitgliedschaft (0 für kostenlosen Mitgliedschaften).
  minimum_fee: z.number().nullable().optional(),
  // Einheit, in der die Dauer des Mitgieldschafts-Typs angegegen wird.
  duration_unit: z.string(),
  // Anzahl der Zeiteinheiten der Mitgliedschaftsperiode (z.B. 1 Jahr, 12 Monate).
  duration_interval: z.number().int().nullable().optional(),
  // Laufende Mitgliedschaften starten am Tag der Anmeldung. Feste Mitgliedschaften starten am fixed_period_start_day.
  period_type: z.string(),
  // For fixed period memberships, month and day (mmdd) on which subscription/membership will start. Period start is back-dated unless after rollover day.
  fixed_period_start_day: z.number().int().nullable().optional(),
  // For fixed period memberships, signups after this day (mmdd) rollover to next period.
  fixed_period_rollover_day: z.number().int().nullable().optional(),
  // FK zu Relationship Type ID
  relationship_type_id: z.string().nullable().optional(),
  // Richtung der Beziehung
  relationship_direction: z.string().nullable().optional(),
  // Maximale Anzahl der zugehörigen Mitgliedschaften.
  max_related: z.number().int().nullable().optional(),
  // Sichtbar
  visibility: z.string().nullable().optional(),
  // Reihenfolge
  weight: z.number().int().nullable().optional(),
  // Receipt Text for membership signup
  receipt_text_signup: z.string().nullable().optional(),
  // Receipt Text for membership renewal
  receipt_text_renewal: z.string().nullable().optional(),
  // 0 = keine automatische Verlängerung; 1 = wählbar, aber nicht Pflicht; 2 = automatische Verlängerung ist verpflichtend
  auto_renew: z.number().int().nullable().optional(),
  // Ist diese Mitgliedsart membership_type aktiviert
  is_active: z.boolean().nullable().optional(),
  // Is provided by an extension
  has_base: z.boolean().nullable().optional(),
  // Name of extension which provides this package
  base_module: z.string().nullable().optional(),
  // When the managed entity was changed from its original settings
  local_modified_date: z.string().nullable().optional(),
})
export type MembershipType = z.infer<typeof MembershipTypeSchema>


// ────────────────────────────────────────────────────────
// MessageTemplate
// ────────────────────────────────────────────────────────

export const MessageTemplateSchema = z.object({
  // Nachrichtenvorlagen-ID
  id: z.number().int().optional(),
  // Aussagekräftiger Titel für diese Nachricht
  msg_title: z.string().nullable().optional(),
  // Betreff für E-Mail-Nachricht.
  msg_subject: z.string().nullable().optional(),
  // Textformatierte Nachricht
  msg_text: z.string().nullable().optional(),
  // HTML-formatierte Nachricht
  msg_html: z.string().nullable().optional(),
  // ist Aktiv
  is_active: z.boolean().optional(),
  // a pseudo-FK to civicrm_option_value
  workflow_id: z.number().int().nullable().optional(),
  // Nachrichtenvorlage Workflow Name
  workflow_name: z.string().nullable().optional(),
  // is this the default message template for the workflow referenced by workflow_id?
  is_default: z.boolean().optional(),
  // is this the reserved message template which we ship for the workflow referenced by workflow_id?
  is_reserved: z.boolean().optional(),
  // Ob diese Nachrichtenvorlage für SMS benutzt wird?
  is_sms: z.boolean().optional(),
  // a pseudo-FK to civicrm_option_value containing PDF Page Format.
  pdf_format_id: z.number().int().nullable().optional(),
  // MessageID that this could revert to
  master_id: z.number().int().nullable().optional(),
})
export type MessageTemplate = z.infer<typeof MessageTemplateSchema>


// ────────────────────────────────────────────────────────
// MosaicoTemplate
// ────────────────────────────────────────────────────────

export const MosaicoTemplateSchema = z.object({
  // Unique Template ID
  id: z.number().int().optional(),
  // Titel
  title: z.string().nullable().optional(),
  // Name of the Mosaico base template (e.g. versafix-1)
  base: z.string().nullable().optional(),
  // Fully renderd HTML
  html: z.string().nullable().optional(),
  // Mosaico metadata (JSON)
  metadata: z.string().nullable().optional(),
  // Mosaico content (JSON)
  content: z.string().nullable().optional(),
  // FK to civicrm_msg_template.
  msg_tpl_id: z.number().int().nullable().optional(),
  // ID of the category this mailing template is currently belongs. Foreign key to civicrm_option_value.
  category_id: z.number().int().nullable().optional(),
  // Domain ID this message template belongs to.
  domain_id: z.number().int().nullable().optional(),
})
export type MosaicoTemplate = z.infer<typeof MosaicoTemplateSchema>


// ────────────────────────────────────────────────────────
// Navigation
// ────────────────────────────────────────────────────────

export const NavigationSchema = z.object({
  // Navigations-ID
  id: z.number().int().optional(),
  // Which Domain is this navigation item for
  domain_id: z.number().int().optional().default("current_domain"),
  // Navigationstitel
  label: z.string().nullable().optional(),
  // Interner Name
  name: z.string().nullable().optional(),
  // URL bei einem benutzerdefiniertem Link
  url: z.string().nullable().optional(),
  // CSS-Klassenname für ein Icon
  icon: z.string().nullable().optional(),
  // Berechtigung(en) benötigt für Zugriff auf Menüeintrag
  permission: z.string().nullable().optional(),
  // Operator to use if item has more than one permission
  permission_operator: z.string().nullable().optional(),
  // Elterneintrag, wird zum Gruppieren benutzt
  parent_id: z.number().int().nullable().optional(),
  // Ist dieser Navigationseintrag aktiv?
  is_active: z.boolean().optional(),
  // Platziere einen Trenner entweder vor oder hinter dem Menüeintrag.
  has_separator: z.number().int().nullable().optional(),
  // Ordering of the navigation items in various blocks.
  weight: z.number().int().optional(),
  // Tiefe in der geschachtelten Hierarchie
  _depth: z.number().int().nullable().optional(),
  // Number of descendents in the nested hierarchy
  _descendents: z.number().int().nullable().optional(),
  // Is provided by an extension
  has_base: z.boolean().nullable().optional(),
  // Name of extension which provides this package
  base_module: z.string().nullable().optional(),
  // When the managed entity was changed from its original settings
  local_modified_date: z.string().nullable().optional(),
})
export type Navigation = z.infer<typeof NavigationSchema>


// ────────────────────────────────────────────────────────
// Note
// ────────────────────────────────────────────────────────

export const NoteSchema = z.object({
  // Notiz-ID
  id: z.number().int().optional(),
  // Name of table where item being referenced is stored.
  entity_table: z.string(),
  // Foreign key to the referenced item.
  entity_id: z.number().int(),
  // Notiz und/oder Kommentar
  note: z.string().nullable().optional(),
  // FK to Contact ID creator
  contact_id: z.number().int().nullable().optional(),
  // Datum zur Notiz
  note_date: z.string().optional(),
  // Wann die Notiz erstellt wurde.
  created_date: z.string().optional(),
  // Wann wurde diese Notiz zuletzt bearbeitet
  modified_date: z.string().optional(),
  // Betreff der Notiz
  subject: z.string().nullable().optional(),
  // Foreign Key to Note Privacy Level (which is an option value pair and hence an implicit FK)
  privacy: z.string().optional(),
  // Tiefe in der geschachtelten Hierarchie
  _depth: z.number().int().nullable().optional(),
  // Number of descendents in the nested hierarchy
  _descendents: z.number().int().nullable().optional(),
})
export type Note = z.infer<typeof NoteSchema>


// ────────────────────────────────────────────────────────
// OpenID
// ────────────────────────────────────────────────────────

export const OpenIDSchema = z.object({
  // Unique OpenID ID
  id: z.number().int().optional(),
  // FK zu Kontakt ID
  contact_id: z.number().int().nullable().optional(),
  // Zu welchem Adresskategorie diese E-Mail gehört.
  location_type_id: z.number().int().nullable().optional(),
  // the OpenID (or OpenID-style http://username.domain/) unique identifier for this contact mainly used for logging in to CiviCRM
  openid: z.string().nullable().optional(),
  // Ob oder ob nicht dieser Benutzer erlaubt ist zum Login
  allowed_to_login: z.boolean().optional(),
  // Ist das der E-Mail-Haupteintrag für diesen Kontakt und Ort?
  is_primary: z.boolean().optional(),
})
export type OpenID = z.infer<typeof OpenIDSchema>


// ────────────────────────────────────────────────────────
// OptionGroup
// ────────────────────────────────────────────────────────

export const OptionGroupSchema = z.object({
  // ID der Optionsgruppe
  id: z.number().int().optional(),
  // Option group name. Used as selection key by class properties which lookup options in civicrm_option_value.
  name: z.string(),
  // Optionsgruppentitel
  title: z.string().nullable().optional(),
  // Optionsgruppenbeschreibung
  description: z.string().nullable().optional(),
  // Type of data stored by this option group.
  data_type: z.string().nullable().optional(),
  // Is this a predefined system option group (i.e. it can not be deleted)?
  is_reserved: z.boolean().optional(),
  // Ist die Optionengruppe aktiv?
  is_active: z.boolean().optional(),
  // A lock to remove the ability to add new options via the UI.
  is_locked: z.boolean().optional(),
  // Which optional columns from the option_value table are in use by this group.
  option_value_fields: z.string().nullable().optional(),
  // Is provided by an extension
  has_base: z.boolean().nullable().optional(),
  // Name of extension which provides this package
  base_module: z.string().nullable().optional(),
  // When the managed entity was changed from its original settings
  local_modified_date: z.string().nullable().optional(),
})
export type OptionGroup = z.infer<typeof OptionGroupSchema>


// ────────────────────────────────────────────────────────
// OptionValue
// ────────────────────────────────────────────────────────

export const OptionValueSchema = z.object({
  // Option-ID
  id: z.number().int().optional(),
  // Gruppe, zu welcher diese Option gehört.
  option_group_id: z.number().int(),
  // Option string as displayed to users - e.g. the label in an HTML OPTION tag.
  label: z.string(),
  // The actual value stored (as a foreign key) in the data record. Functions which need lookup option_value.title should use civicrm_option_value.option_group_id plus civicrm_option_value.value as the key.
  value: z.string(),
  // Stores a fixed (non-translated) name for this option value. Lookup functions should use the name as the key for the option value row.
  name: z.string().nullable().optional(),
  // Use to sort and/or set display properties for sub-set(s) of options within an option group. EXAMPLE: Use for college_interest field, to differentiate partners from non-partners.
  grouping: z.string().nullable().optional(),
  // Bitwise logic can be used to create subsets of options within an option_group for different uses.
  filter: z.number().int().nullable().optional(),
  // Ist das die Standardoption für die Gruppe?
  is_default: z.boolean().nullable().optional(),
  // Controls display sort order.
  weight: z.number().int(),
  // Optionale Beschreibung.
  description: z.string().nullable().optional(),
  // Is this row simply a display header? Expected usage is to render these as OPTGROUP tags within a SELECT field list of options?
  is_optgroup: z.boolean().nullable().optional(),
  // Is this a predefined system object?
  is_reserved: z.boolean().nullable().optional(),
  // Ist diese Option aktiv?
  is_active: z.boolean().nullable().optional(),
  // Component that this option value belongs/caters to.
  component_id: z.number().int().nullable().optional(),
  // Unused deprecated column.
  domain_id: z.number().int().nullable().optional(),
  // Optionssichtbarkeit
  visibility_id: z.number().int().nullable().optional(),
  // crm-i icon class
  icon: z.string().nullable().optional(),
  // Hex Farbwerte, z.B. #ffffff
  color: z.string().nullable().optional(),
  // Is provided by an extension
  has_base: z.boolean().nullable().optional(),
  // Name of extension which provides this package
  base_module: z.string().nullable().optional(),
  // When the managed entity was changed from its original settings
  local_modified_date: z.string().nullable().optional(),
})
export type OptionValue = z.infer<typeof OptionValueSchema>


// ────────────────────────────────────────────────────────
// Organization
// ────────────────────────────────────────────────────────

export const OrganizationSchema = z.object({
  // Eindeutige Kontakt-ID
  id: z.number().int().optional(),
  // Unique trusted external ID (generally from a legacy app/datasource). Particularly useful for deduping operations.
  external_identifier: z.string().nullable().optional(),
  // Formatierter Name, der das bevorzugte Format für Anzeige/Druck/andere Ausgaben angibt.
  display_name: z.string().nullable().optional(),
  // Organisationsname
  organization_name: z.string().nullable().optional(),
  // May be used to over-ride contact view and edit templates.
  contact_sub_type: z.string().nullable().optional(),
  // Keine E-Mails senden
  do_not_email: z.boolean().optional(),
  // Nicht anrufen
  do_not_phone: z.boolean().optional(),
  // Nicht anschreiben
  do_not_mail: z.boolean().optional(),
  // Keine SMS senden
  do_not_sms: z.boolean().optional(),
  // Nicht weitergeben
  do_not_trade: z.boolean().optional(),
  // Hat sich dieser Kontakt von allen Massen-E-Mails der Organisation bzw. der Webseitendomain abgemeldet, sogenanntes Opt-Out?
  is_opt_out: z.boolean().optional(),
  // May be used for SSN, EIN/TIN, Household ID (census) or other applicable unique legal/government ID.
  legal_identifier: z.string().nullable().optional(),
  // Name used for sorting different contact types
  sort_name: z.string().nullable().optional(),
  // Pseudonym.
  nick_name: z.string().nullable().optional(),
  // Gesetzlicher Name.
  legal_name: z.string().nullable().optional(),
  // optionale URL für bevorzugtes Bild (Foto, Logo, etc.), welches für diesen Kontakt angezeigt wird.
  image_URL: z.string().nullable().optional(),
  // Was ist die bevorzugte Kommunikationsart.
  preferred_communication_method: z.string().nullable().optional(),
  // Which language is preferred for communication. FK to languages in civicrm_option_value.
  preferred_language: z.string().nullable().optional(),
  // Schlüssel zur Anfragenverifizierung bezogen auf diesen Kontakt.
  hash: z.string().nullable().optional(),
  // API-Schlüssel zur Anfragenverifizierung bezogen auf diesen Kontakt.
  api_key: z.string().nullable().optional(),
  // woher der Kontakt stammt, z.B. Import, Eintrag vom Spendenmodul...
  source: z.string().nullable().optional(),
  // Communication style (e.g. formal vs. familiar) to use with this contact. FK to communication styles in civicrm_option_value.
  communication_style_id: z.number().int().nullable().optional(),
  // FK to civicrm_option_value.value, that has to be valid registered Email Greeting.
  email_greeting_id: z.number().int().nullable().optional(),
  // Benutzerdefinierte E-Mail-Grußformel
  email_greeting_custom: z.string().nullable().optional(),
  // Cache Email Greeting.
  email_greeting_display: z.string().nullable().optional(),
  // FK to civicrm_option_value.value, that has to be valid registered Postal Greeting.
  postal_greeting_id: z.number().int().nullable().optional(),
  // Benutzerdefinierte Brief-Grußformel.
  postal_greeting_custom: z.string().nullable().optional(),
  // Cache Postal greeting.
  postal_greeting_display: z.string().nullable().optional(),
  // FK to civicrm_option_value.value, that has to be valid registered Addressee.
  addressee_id: z.number().int().nullable().optional(),
  // Benutzerdefinierter Adressat
  addressee_custom: z.string().nullable().optional(),
  // Cache Addressee.
  addressee_display: z.string().nullable().optional(),
  // Is Closed
  is_deceased: z.boolean().optional(),
  // Date closed or disbanded
  deceased_date: z.string().nullable().optional(),
  // Standard Industry Classification Code.
  sic_code: z.string().nullable().optional(),
  // the OpenID (or OpenID-style http://username.domain/) unique identifier for this contact mainly used for logging in to CiviCRM
  user_unique_id: z.string().nullable().optional(),
  // Kontakt ist im Papierkorb
  is_deleted: z.boolean().optional().default(false),
  // Wann der Kontakt erstellt wurde.
  created_date: z.string().nullable().optional(),
  // When was the contact (or closely related entity) was created or modified or deleted.
  modified_date: z.string().nullable().optional(),
  // Deprecated setting for text vs html mailings
  preferred_mail_format: z.string().nullable().optional(),
  // Primary Address ID
  address_primary: z.number().int().nullable().optional(),
  // Billing Address ID
  address_billing: z.number().int().nullable().optional(),
  // Primary Email ID
  email_primary: z.number().int().nullable().optional(),
  // Billing Email ID
  email_billing: z.number().int().nullable().optional(),
  // Primary Phone ID
  phone_primary: z.number().int().nullable().optional(),
  // Billing Phone ID
  phone_billing: z.number().int().nullable().optional(),
  // Primary IM ID
  im_primary: z.number().int().nullable().optional(),
  // Billing IM ID
  im_billing: z.number().int().nullable().optional(),
  // Groups (or sub-groups of groups) to which this contact belongs
  groups: z.array(z.unknown()).nullable().optional(),
  // Filter by tags (including child tags)
  tags: z.array(z.unknown()).nullable().optional(),
})
export type Organization = z.infer<typeof OrganizationSchema>


// ────────────────────────────────────────────────────────
// PCP
// ────────────────────────────────────────────────────────

export const PCPSchema = z.object({
  // Persönliche Kampagnen-Seiten ID
  id: z.number().int().optional(),
  // FK zu Kontakt ID
  contact_id: z.number().int(),
  // Status der persönlichen Spendenkampagnen-Seite
  status_id: z.number().int(),
  // Persönliche Kampagnenseiten-Titel
  title: z.string().nullable().optional(),
  // Einleitungstext
  intro_text: z.string().nullable().optional(),
  // Seitentext
  page_text: z.string().nullable().optional(),
  // Unterstützer-Linktext 
  donate_link_text: z.string().nullable().optional(),
  // Die Zuwendungs- oder Veranstaltungsseite, welche diese persönliche Kampagnenseite asugelöst hat
  page_id: z.number().int(),
  // The type of PCP this is: contribute or event
  page_type: z.string().nullable().optional(),
  // The pcp block that this pcp page was created from
  pcp_block_id: z.number().int(),
  // Benutze Thermometer?
  is_thermometer: z.number().int().nullable().optional(),
  // Zeige Ehrentafel?
  is_honor_roll: z.number().int().nullable().optional(),
  // Zielbetrag dieser Persönlichen Kampagnenseite
  goal_amount: z.number().nullable().optional(),
  // 3 Zeichen String, Wert aus den Konfig-Einstellungen oder aus der Benutzereingabe.
  currency: z.string().nullable().optional(),
  // Ist Persönliche Kampagnen-Seite aktiv?
  is_active: z.boolean().optional(),
  // Benachrichtigen per E-Mail an den Seiteneigentümer, wenn jemand auf der Seite spendet.
  is_notify: z.boolean().optional(),
})
export type PCP = z.infer<typeof PCPSchema>


// ────────────────────────────────────────────────────────
// PCPBlock
// ────────────────────────────────────────────────────────

export const PCPBlockSchema = z.object({
  // PCP block ID
  id: z.number().int().optional(),
  // Entitäts-Tabelle
  entity_table: z.string().nullable().optional(),
  // FK to civicrm_contribution_page.id OR civicrm_event.id
  entity_id: z.number().int(),
  // The type of entity that this pcp targets
  target_entity_type: z.string().optional(),
  // The entity that this pcp targets
  target_entity_id: z.number().int(),
  // FK to civicrm_uf_group.id. Does Personal Campaign Page require manual activation by administrator? (is inactive by default after setup)?
  supporter_profile_id: z.number().int().nullable().optional(),
  // FK to civicrm_option_group with name = PCP owner notifications
  owner_notify_id: z.number().int().nullable().optional(),
  // Does Personal Campaign Page require manual activation by administrator? (is inactive by default after setup)?
  is_approval_needed: z.boolean().optional(),
  // Does Personal Campaign Page allow using tell a friend?
  is_tellfriend_enabled: z.boolean().optional(),
  // Maximum recipient fields allowed in tell a friend
  tellfriend_limit: z.number().int().nullable().optional(),
  // Linktext für PCP.
  link_text: z.string().nullable().optional(),
  // Ist Persönliche Kampagnen-Seiten Block aktiv?
  is_active: z.boolean().optional(),
  // If set, notification is automatically emailed to this email-address on create/update Personal Campaign Page
  notify_email: z.string().nullable().optional(),
})
export type PCPBlock = z.infer<typeof PCPBlockSchema>


// ────────────────────────────────────────────────────────
// Participant
// ────────────────────────────────────────────────────────

export const ParticipantSchema = z.object({
  // Teilnehmer-ID
  id: z.number().int().optional(),
  // FK zu Kontakt ID
  contact_id: z.number().int(),
  // FK to Event ID
  event_id: z.number().int(),
  // Teilnahmestatus ID. FK zu civicrm_participant_status_type. Default of 1 should map to status = Registered.
  status_id: z.number().int().optional(),
  // Teilnehmendenrolle ID. Impliziter FK zu civicrm_option_value where option_group = participant_role.
  role_id: z.string().nullable().optional(),
  // Wann hat sich der Kontakt angemeldet?
  register_date: z.string().nullable().optional(),
  // Quelle dieser Veranstaltungsanmeldung.
  source: z.string().nullable().optional(),
  // Bei kostenpflichtigen Veranstaltungen mit mehreren Gebührenstufen wird die Bezeichnung (Text) mit der Gebührenstufe verknüpft. Beachten Sie, dass wir den Wert der Bezeichnung und nicht den Schlüssel speichern
  fee_level: z.string().nullable().optional(),
  // Test
  is_test: z.boolean().optional().default(false),
  // Ist "Später zahlen"
  is_pay_later: z.boolean().optional(),
  // aktuelle Gebühr Zahlungsprozessor falls bekannt - kann 0 sein.
  fee_amount: z.number().nullable().optional(),
  // FK to Participant ID
  registered_by_id: z.number().int().nullable().optional(),
  // FK to Discount ID
  discount_id: z.number().int().nullable().optional(),
  // 3 Zeichen String, Wert aus der Konfig-Einstellung.
  fee_currency: z.string().nullable().optional(),
  // Rabattbetrag
  discount_amount: z.number().int().nullable().optional(),
  // Auf der Warteliste
  must_wait: z.number().int().nullable().optional(),
  // FK zu Kontakt ID
  transferred_to_contact_id: z.number().int().nullable().optional(),
  // Verantwortlicher Kontakt für die Teilnahmeregistrierung
  created_id: z.number().int().nullable().optional(),
  // When was the participant record was created.
  created_date: z.string().nullable().optional(),
  // When was the participant record created or modified or deleted.
  modified_date: z.string().nullable().optional(),
})
export type Participant = z.infer<typeof ParticipantSchema>


// ────────────────────────────────────────────────────────
// ParticipantStatusType
// ────────────────────────────────────────────────────────

export const ParticipantStatusTypeSchema = z.object({
  // eindeutige Teilnehmerstatus-ID
  id: z.number().int().optional(),
  // sprachenübergreifender Name der Statusart
  name: z.string().nullable().optional(),
  // sprachenbezogenes Label für die Anzeige der Statusart
  label: z.string().nullable().optional(),
  // die zugehörende, grundlegende Statuskategorie
  class: z.string().nullable().optional(),
  // ob es eine vom System verlangte Statuskategorie ist
  is_reserved: z.boolean().optional(),
  // ob dieser Statustyp aktiv ist
  is_active: z.boolean().optional(),
  // ob diese Statuskategorie auf die Anzahl der Teilnehmenden gerechnet wird
  is_counted: z.boolean().optional(),
  // kontrolliert die Sortierreihenfolge
  weight: z.number().int(),
  // ob diese Statuskategorie public, also öffentlich sichtbar ist, ein impliziter Fremdschlüssel auf option_value.value verknüpft zur `visibility` option_group
  visibility_id: z.number().int().nullable().optional(),
  // Is provided by an extension
  has_base: z.boolean().nullable().optional(),
  // Name of extension which provides this package
  base_module: z.string().nullable().optional(),
  // When the managed entity was changed from its original settings
  local_modified_date: z.string().nullable().optional(),
})
export type ParticipantStatusType = z.infer<typeof ParticipantStatusTypeSchema>


// ────────────────────────────────────────────────────────
// Payment
// ────────────────────────────────────────────────────────

export const PaymentSchema = z.object({
  // Financial Transaction ID
  id: z.number().int().optional(),
  // FK to financial_account table.
  from_financial_account_id: z.number().int().nullable().optional(),
  // FK to financial_financial_account table.
  to_financial_account_id: z.number().int().nullable().optional(),
  // date transaction occurred
  trxn_date: z.string().nullable().optional(),
  // amount of transaction
  total_amount: z.number(),
  // aktuelle Gebühr Zahlungsprozessor falls bekannt - kann 0 sein.
  fee_amount: z.number().nullable().optional(),
  // actual funds transfer amount. total less fees. if processor does not report actual fee during transaction, this is set to total_amount.
  net_amount: z.number().nullable().optional(),
  // 3 Zeichen String, Wert aus den Konfig-Einstellungen oder aus der Benutzereingabe.
  currency: z.string().nullable().optional(),
  // Is this entry either a payment or a reversal of a payment?
  is_payment: z.boolean().optional().default(true),
  // Transaction id supplied by external processor. This may not be unique.
  trxn_id: z.string().nullable().optional(),
  // processor result code
  trxn_result_code: z.string().nullable().optional(),
  // pseudo FK to civicrm_option_value of contribution_status_id option_group
  status_id: z.number().int().nullable().optional(),
  // Payment Processor for this financial transaction
  payment_processor_id: z.number().int().nullable().optional(),
  // FK to payment_instrument option group values
  payment_instrument_id: z.number().int().nullable().optional(),
  // FK to accept_creditcard option group values
  card_type_id: z.number().int().nullable().optional(),
  // Check number
  check_number: z.string().nullable().optional(),
  // Last 4 digits of credit card
  pan_truncation: z.string().nullable().optional(),
  // Payment Processor external order reference
  order_reference: z.string().nullable().optional(),
  // Contribution ID linked to the financial trxn
  contribution_id: z.number().int().nullable().optional(),
})
export type Payment = z.infer<typeof PaymentSchema>


// ────────────────────────────────────────────────────────
// PaymentProcessor
// ────────────────────────────────────────────────────────

export const PaymentProcessorSchema = z.object({
  // Zahlungsprozessor-ID
  id: z.number().int().optional(),
  // Which Domain is this match entry for
  domain_id: z.number().int(),
  // Payment Processor Name.
  name: z.string(),
  // Name of processor when shown to CiviCRM administrators.
  title: z.string().optional(),
  // Name of processor when shown to users making a payment.
  frontend_title: z.string().optional(),
  // Additional processor information shown to administrators.
  description: z.string().nullable().optional(),
  // Type ID
  payment_processor_type_id: z.number().int(),
  // Is this processor active?
  is_active: z.boolean().optional(),
  // Is this processor the default?
  is_default: z.boolean().optional(),
  // Is this processor for a test site?
  is_test: z.boolean().optional().default(false),
  // Benutzername
  user_name: z.string().nullable().optional(),
  // Passwort
  password: z.string().nullable().optional(),
  // Signatur
  signature: z.string().nullable().optional(),
  // Seiten-URL
  url_site: z.string().nullable().optional(),
  // API URL
  url_api: z.string().nullable().optional(),
  // URL für wiederkehrende Zahlungen
  url_recur: z.string().nullable().optional(),
  // Schaltflächen-URL
  url_button: z.string().nullable().optional(),
  // Thema, Betreff
  subject: z.string().nullable().optional(),
  // Suffix for PHP class name implementation
  class_name: z.string().nullable().optional(),
  // Billing Mode (deprecated)
  billing_mode: z.number().int(),
  // Can process recurring contributions
  is_recur: z.boolean().optional(),
  // Payment Type: Credit or Debit (deprecated)
  payment_type: z.number().int().nullable().optional(),
  // Payment Instrument ID
  payment_instrument_id: z.number().int().nullable().optional(),
  // array of accepted credit card types
  accepted_credit_cards: z.string().nullable().optional(),
})
export type PaymentProcessor = z.infer<typeof PaymentProcessorSchema>


// ────────────────────────────────────────────────────────
// PaymentProcessorType
// ────────────────────────────────────────────────────────

export const PaymentProcessorTypeSchema = z.object({
  // Payment Processor Type ID
  id: z.number().int().optional(),
  // Payment Processor Type Name.
  name: z.string(),
  // Payment Processor Type Title.
  title: z.string(),
  // Payment Processor Description.
  description: z.string().nullable().optional(),
  // Is this processor active?
  is_active: z.boolean().optional(),
  // Is this processor the default?
  is_default: z.boolean().optional(),
  // Label for User Name if used
  user_name_label: z.string().nullable().optional(),
  // Label for password
  password_label: z.string().nullable().optional(),
  // Label for Signature
  signature_label: z.string().nullable().optional(),
  // Label for Subject
  subject_label: z.string().nullable().optional(),
  // Suffix for PHP class name implementation
  class_name: z.string(),
  // Default Live Site URL
  url_site_default: z.string().nullable().optional(),
  // Default API Site URL
  url_api_default: z.string().nullable().optional(),
  // Default Live Recurring Payments URL
  url_recur_default: z.string().nullable().optional(),
  // Default Live Button URL
  url_button_default: z.string().nullable().optional(),
  // Default Test Site URL
  url_site_test_default: z.string().nullable().optional(),
  // Default Test API URL
  url_api_test_default: z.string().nullable().optional(),
  // Default Test Recurring Payment URL
  url_recur_test_default: z.string().nullable().optional(),
  // Default Test Button URL
  url_button_test_default: z.string().nullable().optional(),
  // Billing Mode (deprecated)
  billing_mode: z.number().int(),
  // Can process recurring contributions
  is_recur: z.boolean().optional(),
  // Payment Type: Credit or Debit (deprecated)
  payment_type: z.number().int().nullable().optional(),
  // Payment Instrument ID
  payment_instrument_id: z.number().int().nullable().optional(),
  // Is provided by an extension
  has_base: z.boolean().nullable().optional(),
  // Name of extension which provides this package
  base_module: z.string().nullable().optional(),
  // When the managed entity was changed from its original settings
  local_modified_date: z.string().nullable().optional(),
})
export type PaymentProcessorType = z.infer<typeof PaymentProcessorTypeSchema>


// ────────────────────────────────────────────────────────
// PaymentToken
// ────────────────────────────────────────────────────────

export const PaymentTokenSchema = z.object({
  // Payment Token ID
  id: z.number().int().optional(),
  // FK to Contact ID for the owner of the token
  contact_id: z.number().int(),
  // Zahlungsprozessor-ID
  payment_processor_id: z.number().int(),
  // Externally provided token string
  token: z.string(),
  // Date created
  created_date: z.string().optional(),
  // Contact ID of token creator
  created_id: z.number().int().nullable().optional(),
  // Date this token expires
  expiry_date: z.string().nullable().optional(),
  // Email at the time of token creation. Useful for fraud forensics
  email: z.string().nullable().optional(),
  // Billing first name at the time of token creation. Useful for fraud forensics
  billing_first_name: z.string().nullable().optional(),
  // Billing middle name at the time of token creation. Useful for fraud forensics
  billing_middle_name: z.string().nullable().optional(),
  // Billing last name at the time of token creation. Useful for fraud forensics
  billing_last_name: z.string().nullable().optional(),
  // Holds the part of the card number or account details that may be retained or displayed
  masked_account_number: z.string().nullable().optional(),
  // IP used when creating the token. Useful for fraud forensics
  ip_address: z.string().nullable().optional(),
})
export type PaymentToken = z.infer<typeof PaymentTokenSchema>


// ────────────────────────────────────────────────────────
// Permission
// ────────────────────────────────────────────────────────

export const PermissionSchema = z.object({
  // Group
  group: z.string().nullable().optional(),
  // Name
  name: z.string().nullable().optional(),
  // Title
  title: z.string().nullable().optional(),
  // Description
  description: z.string().nullable().optional(),
  // Is Synthetic
  is_synthetic: z.boolean().nullable().optional(),
  // Enabled
  is_active: z.boolean().nullable().optional().default(true),
  // List of sub-permissions automatically granted by this one
  implies: z.array(z.unknown()).nullable().optional(),
  // Higher permission that implies this one
  parent: z.string().nullable().optional(),
  // Tiefe in der geschachtelten Hierarchie
  _depth: z.number().int().nullable().optional().default(0),
})
export type Permission = z.infer<typeof PermissionSchema>


// ────────────────────────────────────────────────────────
// Phone
// ────────────────────────────────────────────────────────

export const PhoneSchema = z.object({
  // Eindeutige Telefon-ID
  id: z.number().int().optional(),
  // FK zu Kontakt ID
  contact_id: z.number().int().nullable().optional(),
  // Zu welchem Telefontyp diese Rufnummer gehört.
  location_type_id: z.number().int().nullable().optional(),
  // Ist das die Haupttelefonnummer für diesen Kontakt und Ort.
  is_primary: z.boolean().optional(),
  // Für Rechnungen benutzt?
  is_billing: z.boolean().optional(),
  // Zu welchem Mobilfunkanbieter dieses Telefon gehört.
  mobile_provider_id: z.number().int().nullable().optional(),
  // Komplette Telefonnummer
  phone: z.string().nullable().optional(),
  // Optionale Durchwahl für eine Telefonnummer
  phone_ext: z.string().nullable().optional(),
  // Telefonnummer ohne Leerzeichen, Buchstaben und Satzzeichen.
  phone_numeric: z.string().nullable().optional(),
  // Zu welchem Telefontyp diese Rufnummer gehört.
  phone_type_id: z.number().int().nullable().optional(),
})
export type Phone = z.infer<typeof PhoneSchema>


// ────────────────────────────────────────────────────────
// PreferencesDate
// ────────────────────────────────────────────────────────

export const PreferencesDateSchema = z.object({
  // Date Preference ID
  id: z.number().int().optional(),
  // The meta name for this date (fixed in code)
  name: z.string(),
  // Beschreibung des Datentyps
  description: z.string().nullable().optional(),
  // The start offset relative to current year
  start: z.number().int(),
  // The end offset relative to current year, can be negative
  end: z.number().int(),
  // Der Datentyp
  date_format: z.string().nullable().optional(),
  // Zeitformat
  time_format: z.string().nullable().optional(),
})
export type PreferencesDate = z.infer<typeof PreferencesDateSchema>


// ────────────────────────────────────────────────────────
// Premium
// ────────────────────────────────────────────────────────

export const PremiumSchema = z.object({
  // Premium ID
  id: z.number().int().optional(),
  // Joins these premium settings to another object. Always civicrm_contribution_page for now.
  entity_table: z.string(),
  // Premium entity ID
  entity_id: z.number().int(),
  // Is the Premiums feature enabled for this page?
  premiums_active: z.boolean().optional(),
  // Title for Premiums section.
  premiums_intro_title: z.string().nullable().optional(),
  // Displayed in <div> at top of Premiums section of page. Text and HTML allowed.
  premiums_intro_text: z.string().nullable().optional(),
  // This email address is included in receipts if it is populated and a premium has been selected.
  premiums_contact_email: z.string().nullable().optional(),
  // This phone number is included in receipts if it is populated and a premium has been selected.
  premiums_contact_phone: z.string().nullable().optional(),
  // Boolean. Should we automatically display minimum contribution amount text after the premium descriptions.
  premiums_display_min_contribution: z.boolean().optional(),
  // Label displayed for No Thank-you option in premiums block (e.g. No thank you)
  premiums_nothankyou_label: z.string().nullable().optional(),
  // No Thank-you Position
  premiums_nothankyou_position: z.number().int().nullable().optional(),
})
export type Premium = z.infer<typeof PremiumSchema>


// ────────────────────────────────────────────────────────
// PremiumsProduct
// ────────────────────────────────────────────────────────

export const PremiumsProductSchema = z.object({
  // Zuwendungs-ID
  id: z.number().int().optional(),
  // Foreign key to premiums settings record.
  premiums_id: z.number().int(),
  // Foreign key to each product object.
  product_id: z.number().int(),
  // Reihenfolge
  weight: z.number().int(),
  // FK to Financial Type.
  financial_type_id: z.number().int().nullable().optional(),
})
export type PremiumsProduct = z.infer<typeof PremiumsProductSchema>


// ────────────────────────────────────────────────────────
// PriceField
// ────────────────────────────────────────────────────────

export const PriceFieldSchema = z.object({
  // Preisfeld
  id: z.number().int().optional(),
  // FK zu civicrm_price_set
  price_set_id: z.number().int(),
  // Variablenname/programmatic handle für dieses Feld
  name: z.string(),
  // Text for form field label (also friendly name for administering this field).
  label: z.string(),
  // Html Type
  html_type: z.string(),
  // Enter a quantity for this field?
  is_enter_qty: z.boolean().optional(),
  // Description and/or help text to display before this field.
  help_pre: z.string().nullable().optional(),
  // Description and/or help text to display after this field.
  help_post: z.string().nullable().optional(),
  // Order in which the fields should appear
  weight: z.number().int().nullable().optional(),
  // Should the price be displayed next to the label for each option?
  is_display_amounts: z.boolean().optional(),
  // number of options per line for checkbox and radio
  options_per_line: z.number().int().nullable().optional(),
  // Is this price field active
  is_active: z.boolean().optional(),
  // Is this price field required (value must be > 1)
  is_required: z.boolean().optional(),
  // If non-zero, do not show this field before the date specified
  active_on: z.string().nullable().optional(),
  // If non-zero, do not show this field after the date specified
  expire_on: z.string().nullable().optional(),
  // Optional scripting attributes for field
  javascript: z.string().nullable().optional(),
  // Implicit FK to civicrm_option_group with name = 'visibility'
  visibility_id: z.number().int().nullable().optional(),
})
export type PriceField = z.infer<typeof PriceFieldSchema>


// ────────────────────────────────────────────────────────
// PriceFieldValue
// ────────────────────────────────────────────────────────

export const PriceFieldValueSchema = z.object({
  // Price Field Value
  id: z.number().int().optional(),
  // FK to civicrm_price_field
  price_field_id: z.number().int(),
  // Price field option name
  name: z.string().nullable().optional(),
  // Price field option label
  label: z.string().nullable().optional(),
  // Price field option description.
  description: z.string().nullable().optional(),
  // Price field option pre help text.
  help_pre: z.string().nullable().optional(),
  // Price field option post field help.
  help_post: z.string().nullable().optional(),
  // Price field option amount
  amount: z.number(),
  // Number of participants per field option
  count: z.number().int().nullable().optional(),
  // Max number of participants per field options
  max_value: z.number().int().nullable().optional(),
  // Order in which the field options should appear
  weight: z.number().int().nullable().optional(),
  // FK zum Mitgliedstyp
  membership_type_id: z.number().int().nullable().optional(),
  // Number of terms for this membership
  membership_num_terms: z.number().int().nullable().optional(),
  // Is this default price field option
  is_default: z.boolean().optional(),
  // Is this price field value active
  is_active: z.boolean().optional(),
  // FK to Financial Type.
  financial_type_id: z.number().int().nullable().optional(),
  // Portion of total amount which is NOT tax deductible.
  non_deductible_amount: z.number().optional(),
  // Implicit FK to civicrm_option_group with name = 'visibility'
  visibility_id: z.number().int().nullable().optional(),
})
export type PriceFieldValue = z.infer<typeof PriceFieldValueSchema>


// ────────────────────────────────────────────────────────
// PriceSet
// ────────────────────────────────────────────────────────

export const PriceSetSchema = z.object({
  // Preisschema
  id: z.number().int().optional(),
  // Which Domain is this price-set for
  domain_id: z.number().int().nullable().optional(),
  // Variable name/programmatic handle for this set of price fields.
  name: z.string(),
  // Displayed title for the Price Set.
  title: z.string(),
  // Is this price set active
  is_active: z.boolean().optional(),
  // Beschreibung und/oder Hilfetext, der vor diesem Feld auf dem Formular angezeigt wird.
  help_pre: z.string().nullable().optional(),
  // Description and/or help text to display after fields in form.
  help_post: z.string().nullable().optional(),
  // Optional Javascript script function(s) included on the form with this price_set. Can be used for conditional
  javascript: z.string().nullable().optional(),
  // What components are using this price set?
  extends: z.string(),
  // FK to Financial Type(for membership price sets only).
  financial_type_id: z.number().int().nullable().optional(),
  // Is set if edited on Contribution or Event Page rather than through Manage Price Sets
  is_quick_config: z.boolean().optional(),
  // Is this a predefined system price set (i.e. it can not be deleted, edited)?
  is_reserved: z.boolean().optional(),
  // Minimum Amount required for this set.
  min_amount: z.number().nullable().optional(),
})
export type PriceSet = z.infer<typeof PriceSetSchema>


// ────────────────────────────────────────────────────────
// PriceSetEntity
// ────────────────────────────────────────────────────────

export const PriceSetEntitySchema = z.object({
  // Price Set Entity
  id: z.number().int().optional(),
  // Table which uses this price set
  entity_table: z.string(),
  // Item in table
  entity_id: z.number().int(),
  // price set being used
  price_set_id: z.number().int(),
})
export type PriceSetEntity = z.infer<typeof PriceSetEntitySchema>


// ────────────────────────────────────────────────────────
// PrintLabel
// ────────────────────────────────────────────────────────

export const PrintLabelSchema = z.object({
  // Drucke Etikett ID
  id: z.number().int().optional(),
  // User title for this label layout
  title: z.string().nullable().optional(),
  // Variablenname/programmatic handle für dieses Feld
  name: z.string().nullable().optional(),
  // Beschreibung des Etikettenlayouts
  description: z.string().nullable().optional(),
  // This refers to name column of civicrm_option_value row in name_badge option group
  label_format_name: z.string().nullable().optional(),
  // Implicit FK to civicrm_option_value row in NEW label_type option group
  label_type_id: z.number().int().nullable().optional(),
  // enthält JSON-codierte Konfigurationsoptionen
  data: z.string().nullable().optional(),
  // Ist das der Standard?
  is_default: z.boolean().optional(),
  // Ist diese Option aktiv?
  is_active: z.boolean().optional(),
  // Is this reserved label?
  is_reserved: z.boolean().optional(),
  // FK to civicrm_contact, who created this label layout
  created_id: z.number().int().nullable().optional(),
})
export type PrintLabel = z.infer<typeof PrintLabelSchema>


// ────────────────────────────────────────────────────────
// Product
// ────────────────────────────────────────────────────────

export const ProductSchema = z.object({
  // Product ID
  id: z.number().int().optional(),
  // Required product/premium name
  name: z.string(),
  // Optional description of the product/premium.
  description: z.string().nullable().optional(),
  // Optional product sku or code.
  sku: z.string().nullable().optional(),
  // Store comma-delimited list of color, size, etc. options for the product.
  options: z.string().nullable().optional(),
  // Full or relative URL to uploaded image - fullsize.
  image: z.string().nullable().optional(),
  // Full or relative URL to image thumbnail.
  thumbnail: z.string().nullable().optional(),
  // Sell price or market value for premiums. For tax-deductible contributions, this will be stored as non_deductible_amount in the contribution record.
  price: z.number().nullable().optional(),
  // 3 Zeichen String, Wert aus den Konfig-Einstellungen oder aus der Benutzereingabe.
  currency: z.string().nullable().optional(),
  // FK to Financial Type.
  financial_type_id: z.number().int().nullable().optional(),
  // Minimum contribution required to be eligible to select this premium.
  min_contribution: z.number().nullable().optional(),
  // Actual cost of this product. Useful to determine net return from sale or using this as an incentive.
  cost: z.number().nullable().optional(),
  // Disabling premium removes it from the premiums_premium join table below.
  is_active: z.boolean().optional(),
  // Rolling means we set start/end based on current day, fixed means we set start/end for current year or month (e.g. 1 year + fixed -> we would set start/end for 1/1/06 thru 12/31/06 for any premium chosen in 2006)
  period_type: z.string().nullable().optional(),
  // Month and day (MMDD) that fixed period type subscription or membership starts.
  fixed_period_start_day: z.number().int().nullable().optional(),
  // Duration Unit
  duration_unit: z.string().nullable().optional(),
  // Number of units for total duration of subscription, service, membership (e.g. 12 Months).
  duration_interval: z.number().int().nullable().optional(),
  // Frequency unit and interval allow option to store actual delivery frequency for a subscription or service.
  frequency_unit: z.string().nullable().optional(),
  // Number of units for delivery frequency of subscription, service, membership (e.g. every 3 Months).
  frequency_interval: z.number().int().nullable().optional(),
})
export type Product = z.infer<typeof ProductSchema>


// ────────────────────────────────────────────────────────
// Queue
// ────────────────────────────────────────────────────────

export const QueueSchema = z.object({
  // System-Queue-ID
  id: z.number().int().optional(),
  // Name der Queue
  name: z.string(),
  // Art der Queue
  type: z.string(),
  // Name vom Taskrunner
  runner: z.string().nullable().optional(),
  // Maximale Anzahl der Elemente im Batch
  batch_limit: z.number().int().optional(),
  // Wenn ein Eintrag (oder Batch von Einträgen) zur Bearbeitung benannt ist, wie lange sollen die Einträg(e) reserviert sein. (Sekunden)
  lease_time: z.number().int().optional(),
  // Anzahl der erlaubten Wiederholungen. Setze es auf null (0) um es zu deaktivieren.
  retry_limit: z.number().int().optional(),
  // Anzahl von Sekunden, die gewartet wird, bevor eine fehlgeschlagene Ausführung wiederholt wird.
  retry_interval: z.number().int().nullable().optional(),
  // Ausführungsstatus
  status: z.string().nullable().optional(),
  // Fallback-Verhalten für nicht abgefangene Fehler
  error: z.string().nullable().optional(),
  // Ist das eine Vorlagenkonfiguration (zur Benutzung von anderen/zukünftigen Queues)?
  is_template: z.boolean().optional().default(false),
  // Is provided by an extension
  has_base: z.boolean().nullable().optional(),
  // Name of extension which provides this package
  base_module: z.string().nullable().optional(),
  // When the managed entity was changed from its original settings
  local_modified_date: z.string().nullable().optional(),
})
export type Queue = z.infer<typeof QueueSchema>


// ────────────────────────────────────────────────────────
// QueueItem
// ────────────────────────────────────────────────────────

export const QueueItemSchema = z.object({
  // Queue Item ID
  id: z.number().int().optional(),
  // Name der Warteschlange diese Elements
  queue_name: z.string(),
  // Reihenfolge
  weight: z.number().int(),
  // Datum an dem das Element der Warteschlange hinzugefügt wurde
  submit_time: z.string(),
  // Datum, wann diese Aufgabe verfügbar sein soll; null wenn sofort
  release_time: z.string().nullable().optional(),
  // Zahl, wie oft die Ausführung unternommen wurde.
  run_count: z.number().int().optional(),
  // Is provided by an extension
  has_base: z.boolean().nullable().optional(),
  // Name of extension which provides this package
  base_module: z.string().nullable().optional(),
  // When the managed entity was changed from its original settings
  local_modified_date: z.string().nullable().optional(),
})
export type QueueItem = z.infer<typeof QueueItemSchema>


// ────────────────────────────────────────────────────────
// RecentItem
// ────────────────────────────────────────────────────────

export const RecentItemSchema = z.object({
  // Entity Id
  entity_id: z.number().int(),
  // Entity Type
  entity_type: z.string(),
  // Title
  title: z.string().nullable().optional(),
  // Is Deleted
  is_deleted: z.boolean().nullable().optional(),
  // Icon
  icon: z.string().nullable().optional(),
  // View URL
  view_url: z.string().nullable().optional(),
  // Edit URL
  edit_url: z.string().nullable().optional(),
  // Delete URL
  delete_url: z.string().nullable().optional(),
})
export type RecentItem = z.infer<typeof RecentItemSchema>


// ────────────────────────────────────────────────────────
// Relationship
// ────────────────────────────────────────────────────────

export const RelationshipSchema = z.object({
  // Beziehungs ID
  id: z.number().int().optional(),
  // ID des ersten Kontakts
  contact_id_a: z.number().int(),
  // ID des zweiten Kontakts
  contact_id_b: z.number().int(),
  // Beziehungsart
  relationship_type_id: z.number().int(),
  // Startdatum der Beziehung
  start_date: z.string().nullable().optional(),
  // Enddatum der Beziehung
  end_date: z.string().nullable().optional(),
  // Ist diese Beziehung aktiv?
  is_active: z.boolean().optional(),
  // Optional verbose description for the relationship.
  description: z.string().nullable().optional(),
  // Berechtigungen von Kontakt A, um Kontakt B aufzurufen oder zu bearbeiten
  is_permission_a_b: z.number().int().optional(),
  // Berechtigungen von Kontakt B, um Kontakt A aufzurufen oder zu bearbeiten
  is_permission_b_a: z.number().int().optional(),
  // FK zu civicrm_case
  case_id: z.number().int().nullable().optional(),
  // Beziehung Erstelldatum.
  created_date: z.string().optional(),
  // Beziehung zu letzt bearbeitet.
  modified_date: z.string().optional(),
  // Is active with a non-past end-date
  is_current: z.boolean().nullable().optional(),
})
export type Relationship = z.infer<typeof RelationshipSchema>


// ────────────────────────────────────────────────────────
// RelationshipCache
// ────────────────────────────────────────────────────────

export const RelationshipCacheSchema = z.object({
  // Beziehungscache-ID
  id: z.number().int().optional(),
  // id of the relationship (FK to civicrm_relationship.id)
  relationship_id: z.number().int(),
  // ID der Beziehungsart
  relationship_type_id: z.number().int(),
  // The cache record is a permutation of the original relationship record. The orientation indicates whether it is forward (a_b) or reverse (b_a) relationship.
  orientation: z.string(),
  // ID des ersten Kontakts
  near_contact_id: z.number().int(),
  // Name für die Beziehung von naher_Kontakt zu entfernter_Kontakt.
  near_relation: z.string().nullable().optional(),
  // ID des zweiten Kontakts
  far_contact_id: z.number().int(),
  // Name für die Beziehung von entfernter_Kontakt zu nahen_Kontakt.
  far_relation: z.string().nullable().optional(),
  // Ist diese Beziehung aktiv?
  is_active: z.boolean().optional(),
  // Startdatum der Beziehung
  start_date: z.string().nullable().optional(),
  // Enddatum der Beziehung
  end_date: z.string().nullable().optional(),
  // FK zu civicrm_case
  case_id: z.number().int().nullable().optional(),
  // Is active with a non-past end-date
  is_current: z.boolean().nullable().optional(),
  // Optional verbose description for the relationship.
  description: z.string().nullable().optional(),
  // Beziehung Erstelldatum.
  relationship_created_date: z.string().nullable().optional(),
  // Beziehung zu letzt bearbeitet.
  relationship_modified_date: z.string().nullable().optional(),
  // Whether contact has permission to view or update update the related contact
  permission_near_to_far: z.number().int().nullable().optional(),
  // Whether related contact has permission to view or update this contact
  permission_far_to_near: z.number().int().nullable().optional(),
})
export type RelationshipCache = z.infer<typeof RelationshipCacheSchema>


// ────────────────────────────────────────────────────────
// RelationshipType
// ────────────────────────────────────────────────────────

export const RelationshipTypeSchema = z.object({
  // Primärschlüssel
  id: z.number().int().optional(),
  // Name für die Beziehung von Kontakt_A zu Kontakt_B.
  name_a_b: z.string().nullable().optional(),
  // Bezeichnung für die Beziehung von Kontakt_A zu Kontakt_B.
  label_a_b: z.string().nullable().optional(),
  // Optionaler Name für die Beziehung von Kontakt_B zu Kontakt_A.
  name_b_a: z.string().nullable().optional(),
  // Optionale Bezeichnung für die Beziehung von Kontakt_B zu Kontakt_A.
  label_b_a: z.string().nullable().optional(),
  // Optional verbose description of the relationship type.
  description: z.string().nullable().optional(),
  // If defined, contact_a in a relationship of this type must be a specific contact_type.
  contact_type_a: z.string().nullable().optional(),
  // If defined, contact_b in a relationship of this type must be a specific contact_type.
  contact_type_b: z.string().nullable().optional(),
  // If defined, contact_sub_type_a in a relationship of this type must be a specific contact_sub_type.
  contact_sub_type_a: z.string().nullable().optional(),
  // If defined, contact_sub_type_b in a relationship of this type must be a specific contact_sub_type.
  contact_sub_type_b: z.string().nullable().optional(),
  // Is this relationship type a predefined system type (can not be changed or de-activated)?
  is_reserved: z.boolean().optional(),
  // Ist diese Beziehungsart gegenwärtig aktiv (z.B. kann es benutzt werden um Beziehungen zu erstellen oder zu bearbeiten)?
  is_active: z.boolean().optional(),
  // Is provided by an extension
  has_base: z.boolean().nullable().optional(),
  // Name of extension which provides this package
  base_module: z.string().nullable().optional(),
  // When the managed entity was changed from its original settings
  local_modified_date: z.string().nullable().optional(),
})
export type RelationshipType = z.infer<typeof RelationshipTypeSchema>


// ────────────────────────────────────────────────────────
// ReportInstance
// ────────────────────────────────────────────────────────

export const ReportInstanceSchema = z.object({
  // ID der Reportinstanz
  id: z.number().int().optional(),
  // Für welche Domain ist diese Instanz
  domain_id: z.number().int(),
  // Titel der Berichtinstanz
  title: z.string().nullable().optional(),
  // FK zu civicrm_option_value für die Berichtsvorlage
  report_id: z.string(),
  // when combined with report_id/template uniquely identifies the instance
  name: z.string().nullable().optional(),
  // arguments that are passed in the url when invoking the instance
  args: z.string().nullable().optional(),
  // Beschreibung der Berichtinstanz.
  description: z.string().nullable().optional(),
  // permission required to be able to run this instance
  permission: z.string().nullable().optional(),
  // role required to be able to run this instance
  grouprole: z.string().nullable().optional(),
  // Submitted form values for this report
  form_values: z.string().nullable().optional(),
  // Ist dieser Eintrag aktiv?
  is_active: z.boolean().optional(),
  // FK zur contact Tabelle.
  created_id: z.number().int().nullable().optional(),
  // FK zur contact Tabelle.
  owner_id: z.number().int().nullable().optional(),
  // Betreff der E-Mail
  email_subject: z.string().nullable().optional(),
  // komma-getrennte Liste von E-Mail-Adressen, an die dieser Bericht gesendet wird
  email_to: z.string().nullable().optional(),
  // komma-getrennte Liste von E-Mail-Adressen, an die dieser Bericht gesendet wird
  email_cc: z.string().nullable().optional(),
  // komma-getrennte Liste von E-Mail-Adressen, an die dieser Bericht gesendet wird
  header: z.string().nullable().optional(),
  // komma-getrennte Liste von E-Mail-Adressen, an die dieser Bericht gesendet wird
  footer: z.string().nullable().optional(),
  // FK zur Navigation ID
  navigation_id: z.number().int().nullable().optional(),
  // FK to instance ID drilldown to
  drilldown_id: z.number().int().nullable().optional(),
  // Instanz ist reserviert
  is_reserved: z.boolean().optional(),
  // Is provided by an extension
  has_base: z.boolean().nullable().optional(),
  // Name of extension which provides this package
  base_module: z.string().nullable().optional(),
  // When the managed entity was changed from its original settings
  local_modified_date: z.string().nullable().optional(),
})
export type ReportInstance = z.infer<typeof ReportInstanceSchema>


// ────────────────────────────────────────────────────────
// RiverleaStream
// ────────────────────────────────────────────────────────

export const RiverleaStreamSchema = z.object({
  // Unique RiverleaStream ID
  id: z.number().int().optional(),
  // Machine-name for this stream.
  name: z.string().nullable().optional(),
  // User-facing name for this stream
  label: z.string(),
  // Description of this stream
  description: z.string().nullable().optional(),
  // Reserved streams are not editable through the UI
  is_reserved: z.boolean().optional(),
  // Extension that provides this stream.
  extension: z.string().nullable().optional(),
  // File prefix to stream files within extension
  file_prefix: z.string().nullable().optional(),
  // A file containing stream css - path should be relative to the extension and file_prefix
  css_file: z.string().nullable().optional(),
  // A file containing stream css for darkmode - path should be relative to the extension and file_prefix
  css_file_dark: z.string().nullable().optional(),
  // Variable declarations for this stream
  vars: z.string().nullable().optional(),
  // Variable declaration overrides for the dark mode of this stream
  vars_dark: z.string().nullable().optional(),
  // Custom css for this stream
  custom_css: z.string().nullable().optional(),
  // Custom css for the darkmode of this stream
  custom_css_dark: z.string().nullable().optional(),
  // When the stream was last modified - helps with cache busting.
  modified_date: z.string().nullable().optional(),
  // Is provided by an extension
  has_base: z.boolean().nullable().optional(),
  // Name of extension which provides this package
  base_module: z.string().nullable().optional(),
  // When the managed entity was changed from its original settings
  local_modified_date: z.string().nullable().optional(),
})
export type RiverleaStream = z.infer<typeof RiverleaStreamSchema>


// ────────────────────────────────────────────────────────
// Role
// ────────────────────────────────────────────────────────

export const RoleSchema = z.object({
  // Eindeutige Rollen-ID
  id: z.number().int().optional(),
  // Maschinenname für diese Rolle
  name: z.string(),
  // Menschenlesbarer Name für diese Rolle
  label: z.string(),
  // Liste der für diese Rolle gewährten Berechtigungen
  permissions: z.string(),
  // Only active roles grant permissions
  is_active: z.boolean().optional(),
  // Is provided by an extension
  has_base: z.boolean().nullable().optional(),
  // Name of extension which provides this package
  base_module: z.string().nullable().optional(),
  // When the managed entity was changed from its original settings
  local_modified_date: z.string().nullable().optional(),
})
export type Role = z.infer<typeof RoleSchema>


// ────────────────────────────────────────────────────────
// RolePermission
// ────────────────────────────────────────────────────────

export const RolePermissionSchema = z.object({
  // Group
  group: z.string().nullable().optional(),
  // Name
  name: z.string().nullable().optional(),
  // Permission Title
  title: z.string().nullable().optional(),
  // Description
  description: z.string().nullable().optional(),
  // Permission that implies this one
  parent: z.string().nullable().optional(),
  // Tiefe in der geschachtelten Hierarchie
  _depth: z.number().int().nullable().optional().default(0),
  // Berechtigung explizit für Rolle "Everyone, including anonymous users" gewährt.
  granted_everyone: z.boolean().nullable().optional(),
  // Berechtigung für Rolle "Everyone, including anonymous users" gewährt.
  implied_everyone: z.boolean().nullable().optional(),
  // Berechtigung explizit für Rolle "Staff" gewährt.
  granted_staff: z.boolean().nullable().optional(),
  // Berechtigung für Rolle "Staff" gewährt.
  implied_staff: z.boolean().nullable().optional(),
})
export type RolePermission = z.infer<typeof RolePermissionSchema>


// ────────────────────────────────────────────────────────
// Route
// ────────────────────────────────────────────────────────

export const RouteSchema = z.object({
  // Relative Path
  path: z.string().nullable().optional(),
  // Page Title
  title: z.string().nullable().optional(),
  // Page Callback
  page_callback: z.string().nullable().optional(),
  // Page Arguments
  page_arguments: z.string().nullable().optional(),
  // Path Arguments
  path_arguments: z.string().nullable().optional(),
  // Access Arguments
  access_arguments: z.array(z.unknown()).nullable().optional(),
})
export type Route = z.infer<typeof RouteSchema>


// ────────────────────────────────────────────────────────
// SavedSearch
// ────────────────────────────────────────────────────────

export const SavedSearchSchema = z.object({
  // ID der gespeicherten Suche
  id: z.number().int().optional(),
  // Eindeutiger Name der gespeicherten Suche
  name: z.string().nullable().optional(),
  // Administrative label for search
  label: z.string().nullable().optional(),
  // Übermittelte Formular-Werte für diese Suche
  form_values: z.string().nullable().optional(),
  // Foreign key to civicrm_mapping used for saved search-builder searches.
  mapping_id: z.number().int().nullable().optional(),
  // Foreign key to civicrm_option value table used for saved custom searches.
  search_custom_id: z.number().int().nullable().optional(),
  // Entity name for API based search
  api_entity: z.string().nullable().optional(),
  // Parameters for API based search
  api_params: z.string().nullable().optional(),
  // FK zur contact Tabelle.
  created_id: z.number().int().nullable().optional(),
  // FK zur contact Tabelle.
  modified_id: z.number().int().nullable().optional(),
  // Optional date after which the search is not needed
  expires_date: z.string().nullable().optional(),
  // Wann die Suche erstellt wurde.
  created_date: z.string().optional(),
  // Wann die Suche zu letzt verändert wurde.
  modified_date: z.string().optional(),
  // Beschreibung der gespeicherten Suche
  description: z.string().nullable().optional(),
  // Search templates are used as a starting point for building new searches
  is_template: z.boolean().optional().default(false),
  // Filter by tags (including child tags)
  tags: z.array(z.unknown()).nullable().optional(),
  // Is active with a non-past end-date
  is_current: z.boolean().nullable().optional(),
  // Is provided by an extension
  has_base: z.boolean().nullable().optional(),
  // Name of extension which provides this package
  base_module: z.string().nullable().optional(),
  // When the managed entity was changed from its original settings
  local_modified_date: z.string().nullable().optional(),
})
export type SavedSearch = z.infer<typeof SavedSearchSchema>


// ────────────────────────────────────────────────────────
// SearchDisplay
// ────────────────────────────────────────────────────────

export const SearchDisplaySchema = z.object({
  // Unique SearchDisplay ID
  id: z.number().int().optional(),
  // Unique name for identifying search display
  name: z.string().optional(),
  // Label for identifying search display to administrators
  label: z.string(),
  // FK to saved search table.
  saved_search_id: z.number().int(),
  // Type of display
  type: z.string(),
  // Configuration data for the search display
  settings: z.string().nullable().optional(),
  // Skip permission checks and ACLs when running this display.
  acl_bypass: z.boolean().nullable().optional(),
  // Is provided by an extension
  has_base: z.boolean().nullable().optional(),
  // Name of extension which provides this package
  base_module: z.string().nullable().optional(),
  // When the managed entity was changed from its original settings
  local_modified_date: z.string().nullable().optional(),
  // Is this the default autocomplete display for this entity
  is_autocomplete_default: z.boolean().nullable().optional(),
})
export type SearchDisplay = z.infer<typeof SearchDisplaySchema>


// ────────────────────────────────────────────────────────
// SearchParamSet
// ────────────────────────────────────────────────────────

export const SearchParamSetSchema = z.object({
  // Unique Search Param Set ID
  id: z.number().int().optional(),
  // Name of the form
  afform_name: z.string().nullable().optional(),
  // Editable label for this set of filters
  label: z.string(),
  // JSON filter configuration
  filters: z.string().nullable().optional(),
  // JSON of picked search display columns, indexed by search display
  columns: z.string().nullable().optional(),
  // Icon for this search param set
  icon: z.string().nullable().optional(),
  // Erstellt von
  created_by: z.number().int().nullable().optional(),
  // When created.
  created_date: z.string().nullable().optional(),
  // When this search param set was last modified.
  modified_date: z.string().nullable().optional(),
})
export type SearchParamSet = z.infer<typeof SearchParamSetSchema>


// ────────────────────────────────────────────────────────
// SearchSegment
// ────────────────────────────────────────────────────────

export const SearchSegmentSchema = z.object({
  // Unique SearchSegment ID
  id: z.number().int().optional(),
  // Unique name
  name: z.string().optional(),
  // Label for identifying search segment (will appear as name of calculated field)
  label: z.string(),
  // Description will appear when selecting SearchSegment in the fields dropdown.
  description: z.string().nullable().optional(),
  // Entity for which this set is used.
  entity_name: z.string(),
  // All items in set
  items: z.string().nullable().optional(),
  // Is provided by an extension
  has_base: z.boolean().nullable().optional(),
  // Name of extension which provides this package
  base_module: z.string().nullable().optional(),
  // When the managed entity was changed from its original settings
  local_modified_date: z.string().nullable().optional(),
})
export type SearchSegment = z.infer<typeof SearchSegmentSchema>


// ────────────────────────────────────────────────────────
// Session
// ────────────────────────────────────────────────────────

export const SessionSchema = z.object({
  // Eindeutige Session-ID
  id: z.number().int().optional(),
  // Hexadecimal Session Identifier
  session_id: z.string(),
  // Daten der Session
  data: z.string().nullable().optional(),
  // Timestamp of the last session access
  last_accessed: z.string().nullable().optional(),
})
export type Session = z.infer<typeof SessionSchema>


// ────────────────────────────────────────────────────────
// Setting
// ────────────────────────────────────────────────────────

export const SettingSchema = z.object({
  // Address Standardization Provider.
  address_standardization_provider: z.unknown().nullable().optional(),
  // Provider Consumer Key
  address_standardization_key: z.unknown().nullable().optional(),
  // Provider Consumer Secret
  address_standardization_secret: z.unknown().nullable().optional(),
  // Hide Country in Mailing Labels when same as domain country
  hideCountryMailingLabels: z.unknown().nullable().optional(),
  // A flag indicating whether this system has run a post-installation routine
  installed: z.unknown().nullable().optional(),
  // Enable Components
  enable_components: z.unknown().nullable().optional(),
  // The current domain if CiviCRM is running multi-site.
  domain: z.unknown().nullable().optional(),
  // Enabling Maintenance Mode will restrict certain functionality such as scheduled job runs and REST api calls. If not set, CiviCRM will attempt to check whether the CMS is in maintenance mode.
  core_maintenance_mode: z.unknown().nullable().optional(),
  // Viewing Contacts
  contact_view_options: z.unknown().nullable().optional(),
  // Editing Contacts
  contact_edit_options: z.unknown().nullable().optional(),
  // Advanced Search
  advanced_search_options: z.unknown().nullable().optional(),
  // Contact Dashboard
  user_dashboard_options: z.unknown().nullable().optional(),
  // Address Fields
  address_options: z.unknown().nullable().optional(),
  // Address Display Format
  address_format: z.unknown().nullable().optional(),
  // Mailing Label Format
  mailing_format: z.unknown().nullable().optional(),
  // Individual Display Name Format
  display_name_format: z.unknown().nullable().optional(),
  // Individual Sort Name Format
  sort_name_format: z.unknown().nullable().optional(),
  // Accept profile submissions from external sites
  remote_profile_submissions: z.unknown().nullable().optional(),
  // Allow alerts to auto-dismiss?
  allow_alert_autodismissal: z.unknown().nullable().optional(),
  // Wysiwyg Editor
  editor_id: z.unknown().nullable().optional(),
  // Check for Similar Contacts
  contact_ajax_check_similar: z.unknown().nullable().optional(),
  // Enable Popup Forms
  ajaxPopupsEnabled: z.unknown().nullable().optional(),
  // Background Queues
  enableBackgroundQueue: z.unknown().nullable().optional(),
  // Extern URL Style
  defaultExternUrl: z.unknown().nullable().optional(),
  // Notify Activity Assignees
  activity_assignee_notification: z.unknown().nullable().optional(),
  // Include ICal Invite to Activity Assignees
  activity_assignee_notification_ics: z.unknown().nullable().optional(),
  // Quicksearch results
  contact_autocomplete_options: z.unknown().nullable().optional(),
  // Contact Reference Options
  contact_reference_options: z.unknown().nullable().optional(),
  // Viewing Smart Groups
  contact_smart_group_display: z.unknown().nullable().optional(),
  // Should the smart groups be flushed by cron jobs or user actions
  smart_group_cache_refresh_mode: z.unknown().nullable().optional(),
  // Should the acl cache be flushed by cron jobs or user actions
  acl_cache_refresh_mode: z.unknown().nullable().optional(),
  // Maximum Attachments
  max_attachments: z.unknown().nullable().optional(),
  // Maximum Attachments For Backend Processes
  max_attachments_backend: z.unknown().nullable().optional(),
  // Maximum File Size (in MB)
  maxFileSize: z.unknown().nullable().optional(),
  // Contact Trash and Undelete
  contact_undelete: z.unknown().nullable().optional(),
  // Allow Permanent Delete for contacts who are linked to live financial transactions
  allowPermDeleteFinancial: z.unknown().nullable().optional(),
  // If enabled, CiviCRM will display pop-up notifications (no more than once per day) for security and misconfiguration issues identified in the system check.
  securityAlert: z.unknown().nullable().optional(),
  // Attach PDF copy to receipts
  doNotAttachPDFReceipt: z.unknown().nullable().optional(),
  // Record generated letters
  recordGeneratedLetters: z.unknown().nullable().optional(),
  // DOMPDF Font Folder
  dompdf_font_dir: z.unknown().nullable().optional(),
  // DOMPDF Local Images Folder
  dompdf_chroot: z.unknown().nullable().optional(),
  // DOMPDF Enable Remote Images
  dompdf_enable_remote: z.unknown().nullable().optional(),
  // DOMPDF Log File
  dompdf_log_output_file: z.unknown().nullable().optional(),
  // Path to weasyprint executable
  weasyprint_path: z.unknown().nullable().optional(),
  // Path to wkhtmltopdf executable
  wkhtmltopdfPath: z.unknown().nullable().optional(),
  // Checksum Lifespan
  checksum_timeout: z.unknown().nullable().optional(),
  // Blog feed URL used by the blog dashlet
  blogUrl: z.unknown().nullable().optional(),
  // Service providing CiviCRM community messages
  communityMessagesUrl: z.unknown().nullable().optional(),
  // Service providing the Getting Started data
  gettingStartedUrl: z.unknown().nullable().optional(),
  // Code appended to resource URLs (JS/CSS) to coerce HTTP caching
  resCacheCode: z.unknown().nullable().optional(),
  // Verify SSL?
  verifySSL: z.unknown().nullable().optional(),
  // Force SSL?
  enableSSL: z.unknown().nullable().optional(),
  // WordPress Base Page
  wpBasePage: z.unknown().nullable().optional(),
  // Allow second-degree relationship permissions
  secondDegRelPermissions: z.unknown().nullable().optional(),
  // Disable CiviCRM css
  disable_core_css: z.unknown().nullable().optional(),
  // Display "empowered by CiviCRM"
  empoweredBy: z.unknown().nullable().optional(),
  // Set this when you intend to manage trigger creation outside of CiviCRM
  logging_no_trigger_permission: z.unknown().nullable().optional(),
  // Logging
  logging: z.unknown().nullable().optional(),
  // This is the date when CRM-18193 was implemented
  logging_uniqueid_date: z.unknown().nullable().optional(),
  // Do some tables pre-date CRM-18193?
  logging_all_tables_uniquid: z.unknown().nullable().optional(),
  // CMS Users Table Name
  userFrameworkUsersTableName: z.unknown().nullable().optional(),
  // CiviCRM will use this setting as path to bootstrap WP.
  wpLoadPhp: z.unknown().nullable().optional(),
  // Maximum number of minutes that secure form data should linger
  secure_cache_timeout_minutes: z.unknown().nullable().optional(),
  // Unique Site ID
  site_id: z.unknown().nullable().optional(),
  // Recent Items
  recentItemsMaxCount: z.unknown().nullable().optional(),
  // Recent Items Providers
  recentItemsProviders: z.unknown().nullable().optional(),
  // Import Batch Size
  import_batch_size: z.unknown().nullable().optional(),
  // Disable SQL MEMORY Engine
  disable_sql_memory_engine: z.unknown().nullable().optional(),
  // Default limit for dedupe screen
  dedupe_default_limit: z.unknown().nullable().optional(),
  // Sync CMS Email
  syncCMSEmail: z.unknown().nullable().optional(),
  // Preserve activity filters as a user preference
  preserve_activity_tab_filter: z.unknown().nullable().optional(),
  // Do not notify assignees for
  do_not_notify_assignees_for: z.unknown().nullable().optional(),
  // Acceptable Mime Types that can be used as part of file urls
  requestableMimeTypes: z.unknown().nullable().optional(),
  // Frontend Theme
  theme_frontend: z.unknown().nullable().optional(),
  // Backend Theme
  theme_backend: z.unknown().nullable().optional(),
  // How long should HTTP requests through Guzzle application run for in seconds
  http_timeout: z.unknown().nullable().optional(),
  // If enabled, CiviCRM will not process background queues.
  queue_paused: z.unknown().nullable().optional(),
  // The DSN for the CiviCRM Database.
  civicrm_db_dsn: z.unknown().nullable().optional(),
  // The database name component of the DSN for the CiviCRM Database.
  civicrm_db_name: z.unknown().nullable().optional(),
  // The database user component of the DSN for the CiviCRM Database.
  civicrm_db_user: z.unknown().nullable().optional(),
  // The database password component of the DSN for the CiviCRM Database.
  civicrm_db_password: z.unknown().nullable().optional(),
  // The database host component of the DSN for the CiviCRM Database.
  civicrm_db_host: z.unknown().nullable().optional(),
  // The database port component of the DSN for the CiviCRM Database.
  civicrm_db_port: z.unknown().nullable().optional(),
  // Asset Caching
  assetCache: z.unknown().nullable().optional(),
  // Send CiviCRM errors to CMS logs
  userFrameworkLogging: z.unknown().nullable().optional(),
  // Enable Debugging
  debug_enabled: z.unknown().nullable().optional(),
  // Display Backtrace
  backtrace: z.unknown().nullable().optional(),
  // Environment
  environment: z.unknown().nullable().optional(),
  // ECMAScript Module Loader
  esm_loader: z.unknown().nullable().optional(),
  // Fatal Error Handler
  fatalErrorHandler: z.unknown().nullable().optional(),
  // Resource Base
  resourceBase: z.unknown().nullable().optional(),
  // Temporary Files Directory
  uploadDir: z.unknown().nullable().optional(),
  // Image Directory
  imageUploadDir: z.unknown().nullable().optional(),
  // Custom Files Directory
  customFileUploadDir: z.unknown().nullable().optional(),
  // Custom Template Directory
  customTemplateDir: z.unknown().nullable().optional(),
  // Custom PHP Directory
  customPHPPathDir: z.unknown().nullable().optional(),
  // Extensions Directory
  extensionsDir: z.unknown().nullable().optional(),
  // Extension Repo URL
  ext_repo_url: z.unknown().nullable().optional(),
  // Extension Depth
  ext_max_depth: z.unknown().nullable().optional(),
  // Custom Translate Function
  customTranslateFunction: z.unknown().nullable().optional(),
  // Thousands Separator
  monetaryThousandSeparator: z.unknown().nullable().optional(),
  // Decimal Delimiter
  monetaryDecimalPoint: z.unknown().nullable().optional(),
  // Monetary Amount Display
  moneyformat: z.unknown().nullable().optional(),
  // Monetary Value Display
  moneyvalueformat: z.unknown().nullable().optional(),
  // Default Currency
  defaultCurrency: z.unknown().nullable().optional(),
  // Default Country
  defaultContactCountry: z.unknown().nullable().optional(),
  // Default State/Province
  defaultContactStateProvince: z.unknown().nullable().optional(),
  // Available Countries
  countryLimit: z.unknown().nullable().optional(),
  // Available States and Provinces (by Country)
  provinceLimit: z.unknown().nullable().optional(),
  // If Yes, the system will use the Language set on the logged-in user's record. This can be changed later if using the CiviCRM language switcher.
  inheritLocale: z.unknown().nullable().optional(),
  // Date Format: Complete Date and Time
  dateformatDatetime: z.unknown().nullable().optional(),
  // Date Format: Complete Date
  dateformatFull: z.unknown().nullable().optional(),
  // Date Format: Month and Year
  dateformatPartial: z.unknown().nullable().optional(),
  // Date Format: Time Only
  dateformatTime: z.unknown().nullable().optional(),
  // Date Format: Year Only
  dateformatYear: z.unknown().nullable().optional(),
  // Date Format: Financial Batch
  dateformatFinancialBatch: z.unknown().nullable().optional(),
  // Date Format: Short date Month Day Year
  dateformatshortdate: z.unknown().nullable().optional(),
  // Date Input Format
  dateInputFormat: z.unknown().nullable().optional(),
  // Import / Export Field Separator
  fieldSeparator: z.unknown().nullable().optional(),
  // Fiscal Year Start
  fiscalYearStart: z.unknown().nullable().optional(),
  // Available Languages (Multi-lingual)
  languageLimit: z.unknown().nullable().optional(),
  // Partial Locales
  partial_locales: z.unknown().nullable().optional(),
  // Formatting locale
  format_locale: z.unknown().nullable().optional(),
  // Available Languages
  uiLanguages: z.unknown().nullable().optional(),
  // Default Language
  lcMessages: z.unknown().nullable().optional(),
  // Legacy Encoding
  legacyEncoding: z.unknown().nullable().optional(),
  // Time Input Format
  timeInputFormat: z.unknown().nullable().optional(),
  // Week begins on
  weekBegins: z.unknown().nullable().optional(),
  // Default Language for contacts
  contact_default_language: z.unknown().nullable().optional(),
  // Pinned Countries
  pinnedContactCountries: z.unknown().nullable().optional(),
  // Force source translations to the default language
  force_translation_source_locale: z.unknown().nullable().optional(),
  // Mailing Backend
  mailing_backend: z.unknown().nullable().optional(),
  // VERP Separator
  verpSeparator: z.unknown().nullable().optional(),
  // Simple mail limit
  simple_mail_limit: z.unknown().nullable().optional(),
  // Allow mail from logged in contact
  allow_mail_from_logged_in_contact: z.unknown().nullable().optional(),
  // Use Smarty in scheduled reminders
  scheduled_reminder_smarty: z.unknown().nullable().optional(),
  // Treat SMTP Error 450 4.1.2 as permanent
  smtp_450_is_permanent: z.unknown().nullable().optional(),
  // Geo Provider Key
  geoAPIKey: z.unknown().nullable().optional(),
  // Geocoding Provider
  geoProvider: z.unknown().nullable().optional(),
  // Map Provider Key
  mapAPIKey: z.unknown().nullable().optional(),
  // Mapping Provider
  mapProvider: z.unknown().nullable().optional(),
  // Make CiviCRM aware of multiple domains. You should configure a domain group if enabled
  multisite_is_enabled: z.unknown().nullable().optional(),
  // Contacts created on this site are added to this group
  domain_group_id: z.unknown().nullable().optional(),
  // Domain Event Price Set
  event_price_set_domain_id: z.unknown().nullable().optional(),
  // Unique Email per Domain?
  uniq_email_per_site: z.unknown().nullable().optional(),
  // Menubar position
  menubar_position: z.unknown().nullable().optional(),
  // Menubar color
  menubar_color: z.unknown().nullable().optional(),
  // Autocomplete Results
  search_autocomplete_count: z.unknown().nullable().optional(),
  // Include Order By Clause
  includeOrderByClause: z.unknown().nullable().optional(),
  // Automatic Wildcard
  includeWildCardInName: z.unknown().nullable().optional(),
  // Include Email
  includeEmailInName: z.unknown().nullable().optional(),
  // Include Nickname
  includeNickNameInName: z.unknown().nullable().optional(),
  // Include Alphabetical Pager
  includeAlphabeticalPager: z.unknown().nullable().optional(),
  // Smart group cache timeout
  smartGroupCacheTimeout: z.unknown().nullable().optional(),
  // Advanced Search Profile
  defaultSearchProfileID: z.unknown().nullable().optional(),
  // PrevNext Cache
  prevNextBackend: z.unknown().nullable().optional(),
  // Search Primary Details Only
  searchPrimaryDetailsOnly: z.unknown().nullable().optional(),
  // Quicksearch options
  quicksearch_options: z.unknown().nullable().optional(),
  // Autocomplete Search Displays
  autocomplete_displays: z.unknown().nullable().optional(),
  // Default Search Pager size
  default_pager_size: z.unknown().nullable().optional(),
  // The base URL of the user framework or CMS that CiviCRM is running in.
  userFrameworkBaseURL: z.unknown().nullable().optional(),
  // CiviCRM Resource URL
  userFrameworkResourceURL: z.unknown().nullable().optional(),
  // Image Upload URL
  imageUploadURL: z.unknown().nullable().optional(),
  // Custom CSS URL
  customCSSURL: z.unknown().nullable().optional(),
  // Extension Resource URL
  extensionsURL: z.unknown().nullable().optional(),
  // Authentication guard
  authx_guards: z.unknown().nullable().optional(),
  // Acceptable credentials (Auto Login)
  authx_auto_cred: z.unknown().nullable().optional(),
  // User account requirements (Auto Login)
  authx_auto_user: z.unknown().nullable().optional(),
  // Acceptable credentials (HTTP Header)
  authx_header_cred: z.unknown().nullable().optional(),
  // User account requirements (HTTP Header)
  authx_header_user: z.unknown().nullable().optional(),
  // Acceptable credentials (HTTP Session Login)
  authx_login_cred: z.unknown().nullable().optional(),
  // User account requirements (HTTP Session Login)
  authx_login_user: z.unknown().nullable().optional(),
  // Acceptable credentials (HTTP Parameter)
  authx_param_cred: z.unknown().nullable().optional(),
  // User account requirements (HTTP Parameter)
  authx_param_user: z.unknown().nullable().optional(),
  // Acceptable credentials (HTTP X-Header)
  authx_xheader_cred: z.unknown().nullable().optional(),
  // User account requirements (HTTP X-Header)
  authx_xheader_user: z.unknown().nullable().optional(),
  // Acceptable credentials (Legacy REST)
  authx_legacyrest_cred: z.unknown().nullable().optional(),
  // User account requirements (Legacy REST)
  authx_legacyrest_user: z.unknown().nullable().optional(),
  // Acceptable credentials (Pipe)
  authx_pipe_cred: z.unknown().nullable().optional(),
  // User account requirements (Pipe)
  authx_pipe_user: z.unknown().nullable().optional(),
  // Acceptable credentials (Script)
  authx_script_cred: z.unknown().nullable().optional(),
  // User account requirements (Script)
  authx_script_user: z.unknown().nullable().optional(),
  // Redact Activity Email
  civicaseRedactActivityEmail: z.unknown().nullable().optional(),
  // Allow Multiple Case Clients
  civicaseAllowMultipleClients: z.unknown().nullable().optional(),
  // Activity Type Sorting
  civicaseNaturalActivityTypeSort: z.unknown().nullable().optional(),
  // Include case activities in general activity views.
  civicaseShowCaseActivities: z.unknown().nullable().optional(),
  // CVV required for backoffice?
  cvv_backoffice_required: z.unknown().nullable().optional(),
  // Deprecated, virtualized setting
  contribution_invoice_settings: z.unknown().nullable().optional(),
  // Enable Tax and Invoicing
  invoicing: z.unknown().nullable().optional(),
  // Invoice Prefix
  invoice_prefix: z.unknown().nullable().optional(),
  // Due Date
  invoice_due_date: z.unknown().nullable().optional(),
  // For transmission
  invoice_due_date_period: z.unknown().nullable().optional(),
  // Notes or Standard Terms
  invoice_notes: z.unknown().nullable().optional(),
  // Automatically email invoice when user purchases online
  invoice_is_email_pdf: z.unknown().nullable().optional(),
  // Tax Term
  tax_term: z.unknown().nullable().optional(),
  // Tax Display Settings
  tax_display_settings: z.unknown().nullable().optional(),
  // Enable Deferred Revenue
  deferred_revenue_enabled: z.unknown().nullable().optional(),
  // Default invoice payment page
  default_invoice_page: z.unknown().nullable().optional(),
  // Always post to Accounts Receivable?
  always_post_to_accounts_receivable: z.unknown().nullable().optional(),
  // Automatically update related contributions when Membership Type is changed
  update_contribution_on_membership_type_change: z.unknown().nullable().optional(),
  // Configure how many events should be shown on the dashboard. This overrides the default value of 10 entries.
  show_events: z.unknown().nullable().optional(),
  // Should payment element be shown on the confirmation page instead of the first page?
  event_show_payment_on_confirm: z.unknown().nullable().optional(),
  // Enable Double Opt-in for Profile Group(s) field
  profile_double_optin: z.unknown().nullable().optional(),
  // No-Reply Address
  no_reply_email_address: z.unknown().nullable().optional(),
  // Track replies using VERP in Reply-To header
  track_civimail_replies: z.unknown().nullable().optional(),
  // Enable workflow support for CiviMail
  civimail_workflow: z.unknown().nullable().optional(),
  // Enable global server wide lock for CiviMail
  civimail_server_wide_lock: z.unknown().nullable().optional(),
  // Unsubscribe Methods
  civimail_unsubscribe_methods: z.unknown().nullable().optional(),
  // Enable Custom Reply-To
  replyTo: z.unknown().nullable().optional(),
  // Enable Double Opt-in for Profiles which use the "Add to Group" setting
  profile_add_to_group_double_optin: z.unknown().nullable().optional(),
  // Disable check for mandatory tokens
  disable_mandatory_tokens_check: z.unknown().nullable().optional(),
  // CiviMail dedupes e-mail addresses by default
  dedupe_email_default: z.unknown().nullable().optional(),
  // Hashed Mailing URL's
  hash_mailing_url: z.unknown().nullable().optional(),
  // Enable multiple bulk email address for a contact.
  civimail_multiple_bulk_emails: z.unknown().nullable().optional(),
  // Enable CiviMail to generate Message-ID header
  include_message_id: z.unknown().nullable().optional(),
  // Mailer Batch Limit
  mailerBatchLimit: z.unknown().nullable().optional(),
  // Mailer Job Size
  mailerJobSize: z.unknown().nullable().optional(),
  // Mailer Cron Job Limit
  mailerJobsMax: z.unknown().nullable().optional(),
  // Mailer Throttle Time
  mailThrottleTime: z.unknown().nullable().optional(),
  // Enable CiviMail to create activities on delivery
  write_activity_record: z.unknown().nullable().optional(),
  // Enable automatic CiviMail recipient count display
  auto_recipient_rebuild: z.unknown().nullable().optional(),
  // Enable click-through tracking by default
  url_tracking_default: z.unknown().nullable().optional(),
  // Enable open tracking by default
  open_tracking_default: z.unknown().nullable().optional(),
  // Database Update Frequency
  civimail_sync_interval: z.unknown().nullable().optional(),
  // Default One Click Unsubscribe Mode
  default_oneclick_unsubscribe_mode: z.unknown().nullable().optional(),
  // Default online membership renewal page
  default_renewal_contribution_page: z.unknown().nullable().optional(),
  // reCAPTCHA Site Key
  recaptchaPublicKey: z.unknown().nullable().optional(),
  // reCAPTCHA Secret Key
  recaptchaPrivateKey: z.unknown().nullable().optional(),
  // If enabled, reCAPTCHA will show on all contribution pages.
  forceRecaptcha: z.unknown().nullable().optional(),
  // Backend Dark Mode Control
  riverlea_dark_mode_backend: z.unknown().nullable().optional(),
  // Frontend Dark Mode Control
  riverlea_dark_mode_frontend: z.unknown().nullable().optional(),
  // Credit Notes Prefix
  credit_notes_prefix: z.unknown().nullable().optional(),
  // Duration (in minutes) until a user session expires
  standaloneusers_session_max_lifetime: z.unknown().nullable().optional(),
  // Choose which multi-factor options are required/accepted. Leave blank to disable MFA. TOTP is Time-based One-Time Password which requires an authenticator app to provide a code.
  standalone_mfa_enabled: z.unknown().nullable().optional(),
  // How many days should ‘remember this device’ allow a user to bypass MFA? Use zero to disable this feature.
  standalone_mfa_remember: z.unknown().nullable().optional(),
  // How should the CiviMail composition screen look?
  mosaico_layout: z.unknown().nullable().optional(),
  // Which backend should process images?
  mosaico_graphics: z.unknown().nullable().optional(),
  // for resize of images with width
  mosaico_scale_factor1: z.unknown().nullable().optional(),
  // for resize of images with width
  mosaico_scale_factor2: z.unknown().nullable().optional(),
  // When uploading images, the mosaico editor trims it down to very required size (in pixels). Use scale factor setting to keep some buffer (2x or 3x) so upscale doesn't look distorted or low resolution. Example<br/>3x => Upto 285 pixels (covers both 2 and 3 column block images)<br/>2x => All other sizes (single column block images)
  mosaico_scale_width_limit1: z.unknown().nullable().optional(),
  // When uploading images, the mosaico editor trims it down to very required size (in pixels). Use scale factor setting to keep some buffer (2x or 3x) so upscale doesn't look distorted or low resolution. Example<br/>3x => Upto 285 pixels (covers both 2 and 3 column block images)<br/>2x => All other sizes (single column block images)
  mosaico_scale_width_limit2: z.unknown().nullable().optional(),
  // Mosaico Custom Templates Directory
  mosaico_custom_templates_dir: z.unknown().nullable().optional(),
  // Mosaico Custom Templates URL
  mosaico_custom_templates_url: z.unknown().nullable().optional(),
  // Hide these base templates
  mosaico_hide_base_templates: z.unknown().nullable().optional(),
  // Mosaico hotlist tokens
  mosaico_hotlist_tokens: z.unknown().nullable().optional(),
  // Add a comma-separated list of css color values to override the default palette in the color picker
  mosaico_custom_theme_colors: z.unknown().nullable().optional(),
})
export type Setting = z.infer<typeof SettingSchema>


// ────────────────────────────────────────────────────────
// SiteEmailAddress
// ────────────────────────────────────────────────────────

export const SiteEmailAddressSchema = z.object({
  //  E-Mail-Adresse ID
  id: z.number().int().optional(),
  // Vollständiger Name des Absenders
  display_name: z.string(),
  //  E-Mail-Adresse des Absenders
  email: z.string(),
  // Verwendung dieser E-Mail-Adresse
  description: z.string().nullable().optional(),
  // Ist diese E-Mail-Adresse aktiviert?
  is_active: z.boolean().optional(),
  // Ist das die Standard-E-Mail für die Organisation?
  is_default: z.boolean().optional(),
  // Which Domain is this option value for
  domain_id: z.number().int(),
})
export type SiteEmailAddress = z.infer<typeof SiteEmailAddressSchema>


// ────────────────────────────────────────────────────────
// SiteToken
// ────────────────────────────────────────────────────────

export const SiteTokenSchema = z.object({
  // Seiten-Token ID
  id: z.number().int().optional(),
  // Organisation
  domain_id: z.number().int(),
  // Token-String, z.B. {site.[name]}
  name: z.string(),
  // Benutzersichtbare Bezeichnung in der Token UI
  label: z.string(),
  // Wert des Tokens im HTML-Format.
  body_html: z.string().nullable().optional(),
  // Wert des Tokens im Text-Format.
  body_text: z.string().nullable().optional(),
  // Ist dieser Token aktiv?
  is_active: z.boolean().optional(),
  // Ist dieser Token reserviert?
  is_reserved: z.boolean().optional(),
  // Verantwortlicher Kontakt für die Token-Erstellung
  created_id: z.number().int().nullable().optional(),
  // FK zur contact Tabelle.
  modified_id: z.number().int().nullable().optional(),
  // Wann der Token erstellt oder bearbeitet oder gelöscht wurde.
  modified_date: z.string().nullable().optional(),
  // Is provided by an extension
  has_base: z.boolean().nullable().optional(),
  // Name of extension which provides this package
  base_module: z.string().nullable().optional(),
  // When the managed entity was changed from its original settings
  local_modified_date: z.string().nullable().optional(),
})
export type SiteToken = z.infer<typeof SiteTokenSchema>


// ────────────────────────────────────────────────────────
// SmsProvider
// ────────────────────────────────────────────────────────

export const SmsProviderSchema = z.object({
  // ID des SMS-Anbieters
  id: z.number().int().optional(),
  // Provider internal name points to option_value of option_group sms_provider_name
  name: z.string().nullable().optional(),
  // Provider name visible to user
  title: z.string().nullable().optional(),
  // SMS Provider Username
  username: z.string().nullable().optional(),
  // SMS Provider Password
  password: z.string().nullable().optional(),
  // points to value in civicrm_option_value for group sms_api_type
  api_type: z.number().int(),
  // SMS Provider API URL
  api_url: z.string().nullable().optional(),
  // the api params in xml, http or smtp format
  api_params: z.string().nullable().optional(),
  // SMS Provider is Default?
  is_default: z.boolean().optional(),
  // SMS Provider is Active?
  is_active: z.boolean().optional(),
  // Which Domain is this sms provider for
  domain_id: z.number().int().nullable().optional(),
})
export type SmsProvider = z.infer<typeof SmsProviderSchema>


// ────────────────────────────────────────────────────────
// StateProvince
// ────────────────────────────────────────────────────────

export const StateProvinceSchema = z.object({
  // Bundesland/Provinz-ID
  id: z.number().int().optional(),
  // Name Bundesland/Provinz
  name: z.string().nullable().optional(),
  // 2-4 Zeichen Abkürzung von Bundesland/Provinz
  abbreviation: z.string().nullable().optional(),
  // ID vom Land, zu dem Bundesland/Provinz gehört
  country_id: z.number().int(),
  // Ist diese/r Bundesland/Provinz aktiv?
  is_active: z.boolean().optional(),
})
export type StateProvince = z.infer<typeof StateProvinceSchema>


// ────────────────────────────────────────────────────────
// StatusPreference
// ────────────────────────────────────────────────────────

export const StatusPreferenceSchema = z.object({
  // Unique Status Preference ID
  id: z.number().int().optional(),
  // Which Domain is this Status Preference for
  domain_id: z.number().int(),
  // Name der Statusprüfung, auf welche diese Präferenz referenziert.
  name: z.string(),
  // expires ignore_severity. NULL never hushes.
  hush_until: z.string().nullable().optional(),
  // Hush messages up to and including this severity.
  ignore_severity: z.number().int().nullable().optional(),
  // These settings are per-check, and can't be compared across checks.
  prefs: z.string().nullable().optional(),
  // These values are per-check, and can't be compared across checks.
  check_info: z.string().nullable().optional(),
  // Ist diese Statusprüfung aktiv?
  is_active: z.boolean().optional(),
})
export type StatusPreference = z.infer<typeof StatusPreferenceSchema>


// ────────────────────────────────────────────────────────
// SubscriptionHistory
// ────────────────────────────────────────────────────────

export const SubscriptionHistorySchema = z.object({
  // Interne ID
  id: z.number().int().optional(),
  // CiviCRM-ID
  contact_id: z.number().int(),
  // Gruppen-ID
  group_id: z.number().int().nullable().optional(),
  // Datum der (Ab-)Bestellung
  date: z.string().optional(),
  // How the (un)subscription was triggered
  method: z.string().nullable().optional(),
  // The state of the contact within the group
  status: z.string().nullable().optional(),
  // IP-Adresse oder andere Trackinginfo
  tracking: z.string().nullable().optional(),
})
export type SubscriptionHistory = z.infer<typeof SubscriptionHistorySchema>


// ────────────────────────────────────────────────────────
// Tag
// ────────────────────────────────────────────────────────

export const TagSchema = z.object({
  // Tag-ID
  id: z.number().int().optional(),
  // Eindeutiger Maschinenname
  name: z.string(),
  // User-facing tag name
  label: z.string(),
  // Optionale, ausführliche Beschreibung des Tags
  description: z.string().nullable().optional(),
  // Optional parent id for this tag.
  parent_id: z.number().int().nullable().optional(),
  // Ist dieser Tag suchbar / angezeigt
  is_selectable: z.boolean().optional(),
  // Reserviert
  is_reserved: z.boolean().optional(),
  // Tagset
  is_tagset: z.boolean().optional(),
  // Benutzt für
  used_for: z.string().nullable().optional(),
  // FK to civicrm_contact, wer diesen Tag angelegt hat
  created_id: z.number().int().nullable().optional(),
  // Hex Farbwerte, z.B. #ffffff
  color: z.string().nullable().optional(),
  // Datum und Uhrzeit, an dem der Tag erstellt wurde.
  created_date: z.string().nullable().optional(),
  // Is provided by an extension
  has_base: z.boolean().nullable().optional(),
  // Name of extension which provides this package
  base_module: z.string().nullable().optional(),
  // When the managed entity was changed from its original settings
  local_modified_date: z.string().nullable().optional(),
})
export type Tag = z.infer<typeof TagSchema>


// ────────────────────────────────────────────────────────
// Totp
// ────────────────────────────────────────────────────────

export const TotpSchema = z.object({
  // Eindeutige TOTP-ID
  id: z.number().int().optional(),
  // Reference to User (UFMatch) ID
  user_id: z.number().int(),
  // Encrypted Base64 encoded TOTP Seed
  seed: z.string(),
  // benutzter Hash-Algorithmus
  hash: z.string().optional(),
  // Seconds each code lasts
  period: z.number().int().optional(),
  // Length of codes
  length: z.number().int().optional(),
})
export type Totp = z.infer<typeof TotpSchema>


// ────────────────────────────────────────────────────────
// Translation
// ────────────────────────────────────────────────────────

export const TranslationSchema = z.object({
  // Eindeutige String-ID
  id: z.number().int().optional(),
  // Table where referenced item is stored
  entity_table: z.string().nullable().optional(),
  // Field where referenced item is stored
  entity_field: z.string().nullable().optional(),
  // ID der betroffenen Entität.
  entity_id: z.number().int().nullable().optional(),
  // Relevante Sprache
  language: z.string(),
  // Specify whether the string is active, draft, etc
  status_id: z.number().int().optional(),
  // Übersetzter String
  string: z.string().nullable().optional(),
  // Alternate FK when using translation_source instead of entity_table / entity_id
  source_key: z.string().nullable().optional(),
})
export type Translation = z.infer<typeof TranslationSchema>


// ────────────────────────────────────────────────────────
// TranslationSource
// ────────────────────────────────────────────────────────

export const TranslationSourceSchema = z.object({
  // Unique Source ID
  id: z.number().int().optional(),
  // Table where referenced item is stored
  entity: z.string(),
  // Field where referenced item is stored
  entity_field: z.string().nullable().optional(),
  // ID der betroffenen Entität.
  entity_id: z.number().int().nullable().optional(),
  // hash(entity_name,entity_id,entity_field,entity)
  context_key: z.string(),
  // Source text for referencing translations
  source: z.string(),
  // hash(source)
  source_key: z.string(),
})
export type TranslationSource = z.infer<typeof TranslationSourceSchema>


// ────────────────────────────────────────────────────────
// UFField
// ────────────────────────────────────────────────────────

export const UFFieldSchema = z.object({
  // Eindeutige Tabellen-ID
  id: z.number().int().optional(),
  // Zu welchem Formular dieses Feld gehört.
  uf_group_id: z.number().int(),
  // Name for CiviCRM field which is being exposed for sharing.
  field_name: z.string(),
  // Kann dieses Feld gegenwärtig geteilt werden? Versteckt das Feld im Sharing-Kontext, falls FALSCH.
  is_active: z.boolean().optional(),
  // das Feld kann nur angesehen werden und ist nicht bearbeitbar in Benutzerformularen.
  is_view: z.boolean().optional(),
  // Ist es ein Pflichtfeld, wenn es für ein Nutzer- oder Registrierungsformular benutzt wird?
  is_required: z.boolean().optional(),
  // Steuert die Anzeigereihenfolge der Felder bei der Anzeige von Benutzerrahmenfeldern in Registrierungs- und Kontobearbeitungsformularen.
  weight: z.number().int().optional(),
  // Description and/or help text to display after this field.
  help_post: z.string().nullable().optional(),
  // Description and/or help text to display before this field.
  help_pre: z.string().nullable().optional(),
  // In welchen Kontext(en) ist dieses Feld sichtbar?
  visibility: z.string().nullable().optional(),
  // Is this field included as a column in the selector table?
  in_selector: z.boolean().optional(),
  // Ist dieses Feld in Suchformularen oder Profilen enthalten?
  is_searchable: z.boolean().optional(),
  // Adresskategorie für diese Zuordnung / Mapping, falls benötigt
  location_type_id: z.number().int().nullable().optional(),
  // Phone Type ID, if required
  phone_type_id: z.number().int().nullable().optional(),
  // Webseitenart, falls benötigt
  website_type_id: z.number().int().nullable().optional(),
  // Um Etiketten für die Felder zu speichern.
  label: z.string(),
  // Das Feld speichert den Feldtyp (z.B. Person, Haushalt... Feld etc.).
  field_type: z.string().nullable().optional(),
  // Ist dieses Feld durch andere CiviCRM-Funktionen reserviert?
  is_reserved: z.boolean().optional(),
  // Mehrfacheintragung auflisten?
  is_multi_summary: z.boolean().optional(),
  // Is provided by an extension
  has_base: z.boolean().nullable().optional(),
  // Name of extension which provides this package
  base_module: z.string().nullable().optional(),
  // When the managed entity was changed from its original settings
  local_modified_date: z.string().nullable().optional(),
})
export type UFField = z.infer<typeof UFFieldSchema>


// ────────────────────────────────────────────────────────
// UFGroup
// ────────────────────────────────────────────────────────

export const UFGroupSchema = z.object({
  // Eindeutige Tabellen-ID
  id: z.number().int().optional(),
  // Name of the UF group for directly addressing it in the codebase
  name: z.string(),
  // Ist dieses Profil aktuell aktiv? Versteckt alle relevanten Felder in Sharing-Kontext, falls FALSCH.
  is_active: z.boolean().optional(),
  // Komma-getrennte Liste mit Art(en) der Profilfelder.
  group_type: z.string().nullable().optional(),
  // Formulartitel.
  title: z.string().optional(),
  // Öffentlicher Titel des Profilformulars
  frontend_title: z.string().optional(),
  // Optional verbose description of the profile.
  description: z.string().nullable().optional(),
  // Beschreibung und/oder Hilfetext, der vor diesem Feld auf dem Formular angezeigt wird.
  help_pre: z.string().nullable().optional(),
  // Description and/or help text to display after fields in form.
  help_post: z.string().nullable().optional(),
  // Group id, foreign key from civicrm_group
  limit_listings_group_id: z.number().int().nullable().optional(),
  // Weiterleiten zur URL nach Einreichung.
  post_url: z.string().nullable().optional(),
  // foreign key to civicrm_group_id
  add_to_group_id: z.number().int().nullable().optional(),
  // Soll ein CAPTCHA Widget auf diesem Profilformular verwendet werden.
  add_captcha: z.boolean().optional(),
  // Wollen wir Ergebnisse von diesem Profil mappen.
  is_map: z.boolean().optional(),
  // Should edit link display in profile selector
  is_edit_link: z.boolean().optional(),
  // Sollten wir einen Link zum Profil der Website in der Profilauswahl anzeigen?
  is_uf_link: z.boolean().optional(),
  // Sollen wir den Kontaktdatensatz aktualisieren wenn ein Duplikat gefunden wird
  is_update_dupe: z.boolean().optional(),
  // Weiterleiten zur URL, wenn der Abbrechen-Button geklickt wird.
  cancel_url: z.string().nullable().optional(),
  // Sollen wir einen CMS-Benutzer für dieses Profil erstellen
  is_cms_user: z.boolean().optional(),
  // Benachrichtigung nach Profileingabe
  notify: z.string().nullable().optional(),
  // Ist dieses Feld durch andere CiviCRM-Funktionen reserviert?
  is_reserved: z.boolean().optional(),
  // FK to civicrm_contact, who created this UF group
  created_id: z.number().int().nullable().optional(),
  // Date and time this UF group was created.
  created_date: z.string().nullable().optional(),
  // Sollen wir die Umkreissuche in dieses Profil-Suchfeld einschließen?
  is_proximity_search: z.boolean().optional(),
  // Text, der auf dem Abbrechen-Button angezeigt wird im Erstell- oder Editiermodus
  cancel_button_text: z.string().nullable().optional(),
  // Benutzerdefinierter Text, der auf dem Submit-Button angezeigt wird im Erstell- oder Editiermodus von Profilen
  submit_button_text: z.string().nullable().optional(),
  // Soll ein Abbrechen-Knopf in dieses Profil aufgenommen werden.
  add_cancel_button: z.boolean().optional(),
  // Is provided by an extension
  has_base: z.boolean().nullable().optional(),
  // Name of extension which provides this package
  base_module: z.string().nullable().optional(),
  // When the managed entity was changed from its original settings
  local_modified_date: z.string().nullable().optional(),
})
export type UFGroup = z.infer<typeof UFGroupSchema>


// ────────────────────────────────────────────────────────
// UFJoin
// ────────────────────────────────────────────────────────

export const UFJoinSchema = z.object({
  // Eindeutige Tabellen-ID
  id: z.number().int().optional(),
  // Is this join currently active?
  is_active: z.boolean().optional(),
  // Module which owns this uf_join instance, e.g. User Registration, CiviDonate, etc.
  module: z.string(),
  // Name of table where item being referenced is stored. Modules which only need a single collection of uf_join instances may choose not to populate entity_table and entity_id.
  entity_table: z.string().nullable().optional(),
  // Foreign key to the referenced item.
  entity_id: z.number().int().nullable().optional(),
  // Controls display order when multiple user framework groups are setup for concurrent display.
  weight: z.number().int().optional(),
  // Zu welchem Formular dieses Feld gehört.
  uf_group_id: z.number().int(),
  // Json serialized array of data used by the ufjoin.module
  module_data: z.string().nullable().optional(),
})
export type UFJoin = z.infer<typeof UFJoinSchema>


// ────────────────────────────────────────────────────────
// UFMatch
// ────────────────────────────────────────────────────────

export const UFMatchSchema = z.object({
  // Systemgenerierte ID.
  id: z.number().int().optional(),
  // Which Domain is this match entry for
  domain_id: z.number().int(),
  // UF-ID
  uf_id: z.number().int(),
  // UF-Name
  uf_name: z.string().nullable().optional(),
  // FK zu Kontakt ID
  contact_id: z.number().int().nullable().optional(),
  // UI language preferred by the given user/contact
  language: z.string().nullable().optional(),
})
export type UFMatch = z.infer<typeof UFMatchSchema>


// ────────────────────────────────────────────────────────
// User
// ────────────────────────────────────────────────────────

export const UserSchema = z.object({
  // Unique User ID
  id: z.number().int().optional(),
  // Which Domain is this match entry for
  domain_id: z.number().int(),
  // UF ID. Redundant in Standalone. Needs to be identical to id.
  uf_id: z.number().int().optional(),
  // E-Mail (z.B. zum Passwort zurücksetzen)
  uf_name: z.string().nullable().optional(),
  // FK zu Kontakt ID
  contact_id: z.number().int().nullable().optional(),
  // Benutzername
  username: z.string(),
  // Hashed, not plaintext password
  hashed_password: z.string(),
  // Wann erstellt
  when_created: z.string().nullable().optional(),
  // Wann zuletzt zugegriffen
  when_last_accessed: z.string().nullable().optional(),
  // Wann aktualisiert
  when_updated: z.string().nullable().optional(),
  // Aktiviert
  is_active: z.boolean().optional(),
  // Zeitzone des Benutzers
  timezone: z.string().nullable().optional(),
  // UI language preferred by the given user/contact
  language: z.string().nullable().optional(),
  // The unspent token
  password_reset_token: z.string().nullable().optional(),
  // Role ids belonging to this user.
  roles: z.array(z.unknown()).nullable().optional(),
})
export type User = z.infer<typeof UserSchema>


// ────────────────────────────────────────────────────────
// UserJob
// ────────────────────────────────────────────────────────

export const UserJobSchema = z.object({
  // Job ID
  id: z.number().int().optional(),
  // Eindeutiger Name für die Aufgabe.
  name: z.string().nullable().optional(),
  // Label for job.
  label: z.string().nullable().optional(),
  // FK zur contact Tabelle.
  created_id: z.number().int().nullable().optional(),
  // Datum und Uhrzeit, an dem die Aufgabe erstellt wurde.
  created_date: z.string().optional(),
  // Datum und Uhrzeit, an dem die Import-Aufgabe gestartet wurde.
  start_date: z.string().nullable().optional(),
  // Datum und Uhrzeit, an dem die Import-Aufgabe beendet wurde.
  end_date: z.string().nullable().optional(),
  // Date and time to clean up after this import job (temp table deletion date).
  expires_date: z.string().nullable().optional(),
  // User Job Status ID
  status_id: z.number().int(),
  // Name of the job type, which will allow finding the correct class
  job_type: z.string(),
  // FK to Queue
  queue_id: z.number().int().nullable().optional(),
  // Batch import search display
  search_display_id: z.number().int().nullable().optional(),
  // Data pertaining to job configuration
  metadata: z.string().nullable().optional(),
  // Ist das eine Vorlagenkonfiguration (zur Benutzung von anderen/zukünftigen Aufagben)?
  is_template: z.boolean().optional().default(false),
  // Is active with a non-past end-date
  is_current: z.boolean().nullable().optional(),
  // Is provided by an extension
  has_base: z.boolean().nullable().optional(),
  // Name of extension which provides this package
  base_module: z.string().nullable().optional(),
  // When the managed entity was changed from its original settings
  local_modified_date: z.string().nullable().optional(),
})
export type UserJob = z.infer<typeof UserJobSchema>


// ────────────────────────────────────────────────────────
// UserRole
// ────────────────────────────────────────────────────────

export const UserRoleSchema = z.object({
  // Unique UserRole ID
  id: z.number().int().optional(),
  // FK zu User
  user_id: z.number().int().nullable().optional(),
  // FK zu Rolle
  role_id: z.number().int().nullable().optional(),
})
export type UserRole = z.infer<typeof UserRoleSchema>


// ────────────────────────────────────────────────────────
// Website
// ────────────────────────────────────────────────────────

export const WebsiteSchema = z.object({
  // Eindeutige Webseiten-ID
  id: z.number().int().optional(),
  // FK zu Kontakt ID
  contact_id: z.number().int().nullable().optional(),
  // Webseite
  url: z.string().nullable().optional(),
  // Zu welchem Webseitentyp diese Seite gehört
  website_type_id: z.number().int().nullable().optional(),
})
export type Website = z.infer<typeof WebsiteSchema>


// ────────────────────────────────────────────────────────
// WordReplacement
// ────────────────────────────────────────────────────────

export const WordReplacementSchema = z.object({
  // Wortersetzungs-ID
  id: z.number().int().optional(),
  // Wort, welches ersetzt werden muss
  find_word: z.string().nullable().optional(),
  // Wort, welches das zu ersetzende Wort ersetzt
  replace_word: z.string().nullable().optional(),
  // Ist dieser Eintrag aktiv?
  is_active: z.boolean().optional(),
  // Word Replacement Match Type
  match_type: z.string().nullable().optional(),
  // FK to Domain ID. This is for Domain specific word replacement
  domain_id: z.number().int().nullable().optional(),
})
export type WordReplacement = z.infer<typeof WordReplacementSchema>


// ────────────────────────────────────────────────────────
// WorkflowMessage
// ────────────────────────────────────────────────────────

export const WorkflowMessageSchema = z.object({
  // Name
  name: z.string().nullable().optional(),
  // Group
  group: z.string().nullable().optional(),
  // Class
  class: z.string().nullable().optional(),
  // Description
  description: z.string().nullable().optional(),
  // Support Level
  support: z.string().nullable().optional(),
})
export type WorkflowMessage = z.infer<typeof WorkflowMessageSchema>


// ────────────────────────────────────────────────────────
// WorldRegion
// ────────────────────────────────────────────────────────

export const WorldRegionSchema = z.object({
  // Länder-ID
  id: z.number().int().optional(),
  // Name der Region ist mit dem Land verbunden
  name: z.string().nullable().optional(),
})
export type WorldRegion = z.infer<typeof WorldRegionSchema>
