// Auto-generated — do not edit manually
// Run generate_civicrm_types.py to regenerate

export type WhereClause = [string, string, unknown?][]
export type OrderByClause = Record<string, "ASC" | "DESC">

export interface CiviCRMClientOptions {
  baseUrl: string
  apiKey: string
  /** Skip TLS verification (Node.js only, dev use) */
  checkPermissions?: boolean
}

interface ApiResponse<T> {
  values: T[]
  count: number
  error_code?: number
  error_message?: string
}

async function apiCall<T>(
  opts: CiviCRMClientOptions,
  entity: string,
  action: string,
  params: Record<string, unknown>,
): Promise<T[]> {
  const url = `${opts.baseUrl.replace(/\/$/, "")}/civicrm/ajax/api4/${entity}/${action}`
  const body = new URLSearchParams({
    params: JSON.stringify({
      checkPermissions: opts.checkPermissions ?? false,
      ...params,
    }),
  })

  const res = await fetch(url, {
    method: "POST",
    headers: {
      Authorization: `Bearer ${opts.apiKey}`,
      "X-Requested-With": "XMLHttpRequest",
      "Content-Type": "application/x-www-form-urlencoded",
      Accept: "application/json",
    },
    body: body.toString(),
  })

  if (!res.ok) {
    throw new Error(`CiviCRM API HTTP ${res.status}: ${res.statusText}`)
  }

  const json: ApiResponse<T> = await res.json()
  if (json.error_message) {
    throw new Error(`CiviCRM API error ${json.error_code}: ${json.error_message}`)
  }

  return json.values
}

import { ACLSchema, type ACL, ACLEntityRoleSchema, type ACLEntityRole, ActionScheduleSchema, type ActionSchedule, ActivitySchema, type Activity, ActivityContactSchema, type ActivityContact, AddressSchema, type Address, AfformSchema, type Afform, AfformBehaviorSchema, type AfformBehavior, AfformSubmissionSchema, type AfformSubmission, BatchSchema, type Batch, BouncePatternSchema, type BouncePattern, BounceTypeSchema, type BounceType, CaseSchema, type Case, CaseActivitySchema, type CaseActivity, CaseContactSchema, type CaseContact, CaseTypeSchema, type CaseType, ContactSchema, type Contact, ContactTypeSchema, type ContactType, ContributionSchema, type Contribution, ContributionPageSchema, type ContributionPage, ContributionProductSchema, type ContributionProduct, ContributionRecurSchema, type ContributionRecur, ContributionSoftSchema, type ContributionSoft, CountrySchema, type Country, CountySchema, type County, CustomFieldSchema, type CustomField, CustomGroupSchema, type CustomGroup, DashboardSchema, type Dashboard, DashboardContactSchema, type DashboardContact, DedupeExceptionSchema, type DedupeException, DedupeRuleSchema, type DedupeRule, DedupeRuleGroupSchema, type DedupeRuleGroup, DiscountSchema, type Discount, DomainSchema, type Domain, EmailSchema, type Email, EntitySchema, type Entity, EntityBatchSchema, type EntityBatch, EntityFileSchema, type EntityFile, EntityFinancialAccountSchema, type EntityFinancialAccount, EntityFinancialTrxnSchema, type EntityFinancialTrxn, EntityTagSchema, type EntityTag, EventSchema, type Event, ExampleDataSchema, type ExampleData, ExtensionSchema, type Extension, FileSchema, type File, FinancialAccountSchema, type FinancialAccount, FinancialItemSchema, type FinancialItem, FinancialTrxnSchema, type FinancialTrxn, FinancialTypeSchema, type FinancialType, GroupSchema, type Group, GroupContactSchema, type GroupContact, GroupNestingSchema, type GroupNesting, GroupOrganizationSchema, type GroupOrganization, GroupSubscriptionSchema, type GroupSubscription, HouseholdSchema, type Household, IMSchema, type IM, IndividualSchema, type Individual, JobSchema, type Job, JobLogSchema, type JobLog, LineItemSchema, type LineItem, LocBlockSchema, type LocBlock, LocationTypeSchema, type LocationType, LogSchema, type Log, MailSettingsSchema, type MailSettings, MailingSchema, type Mailing, MailingComponentSchema, type MailingComponent, MailingEventBounceSchema, type MailingEventBounce, MailingEventConfirmSchema, type MailingEventConfirm, MailingEventDeliveredSchema, type MailingEventDelivered, MailingEventOpenedSchema, type MailingEventOpened, MailingEventQueueSchema, type MailingEventQueue, MailingEventReplySchema, type MailingEventReply, MailingEventSubscribeSchema, type MailingEventSubscribe, MailingEventTrackableURLOpenSchema, type MailingEventTrackableURLOpen, MailingEventUnsubscribeSchema, type MailingEventUnsubscribe, MailingGroupSchema, type MailingGroup, MailingJobSchema, type MailingJob, MailingTrackableURLSchema, type MailingTrackableURL, ManagedSchema, type Managed, MappingSchema, type Mapping, MappingFieldSchema, type MappingField, MembershipSchema, type Membership, MembershipBlockSchema, type MembershipBlock, MembershipLogSchema, type MembershipLog, MembershipStatusSchema, type MembershipStatus, MembershipTypeSchema, type MembershipType, MessageTemplateSchema, type MessageTemplate, MosaicoTemplateSchema, type MosaicoTemplate, NavigationSchema, type Navigation, NoteSchema, type Note, OpenIDSchema, type OpenID, OptionGroupSchema, type OptionGroup, OptionValueSchema, type OptionValue, OrganizationSchema, type Organization, PCPSchema, type PCP, PCPBlockSchema, type PCPBlock, ParticipantSchema, type Participant, ParticipantStatusTypeSchema, type ParticipantStatusType, PaymentSchema, type Payment, PaymentProcessorSchema, type PaymentProcessor, PaymentProcessorTypeSchema, type PaymentProcessorType, PaymentTokenSchema, type PaymentToken, PermissionSchema, type Permission, PhoneSchema, type Phone, PreferencesDateSchema, type PreferencesDate, PremiumSchema, type Premium, PremiumsProductSchema, type PremiumsProduct, PriceFieldSchema, type PriceField, PriceFieldValueSchema, type PriceFieldValue, PriceSetSchema, type PriceSet, PriceSetEntitySchema, type PriceSetEntity, PrintLabelSchema, type PrintLabel, ProductSchema, type Product, QueueSchema, type Queue, QueueItemSchema, type QueueItem, RecentItemSchema, type RecentItem, RelationshipSchema, type Relationship, RelationshipCacheSchema, type RelationshipCache, RelationshipTypeSchema, type RelationshipType, ReportInstanceSchema, type ReportInstance, RiverleaStreamSchema, type RiverleaStream, RoleSchema, type Role, RolePermissionSchema, type RolePermission, RouteSchema, type Route, SavedSearchSchema, type SavedSearch, SearchDisplaySchema, type SearchDisplay, SearchParamSetSchema, type SearchParamSet, SearchSegmentSchema, type SearchSegment, SessionSchema, type Session, SettingSchema, type Setting, SiteEmailAddressSchema, type SiteEmailAddress, SiteTokenSchema, type SiteToken, SmsProviderSchema, type SmsProvider, StateProvinceSchema, type StateProvince, StatusPreferenceSchema, type StatusPreference, SubscriptionHistorySchema, type SubscriptionHistory, TagSchema, type Tag, TotpSchema, type Totp, TranslationSchema, type Translation, TranslationSourceSchema, type TranslationSource, UFFieldSchema, type UFField, UFGroupSchema, type UFGroup, UFJoinSchema, type UFJoin, UFMatchSchema, type UFMatch, UserSchema, type User, UserJobSchema, type UserJob, UserRoleSchema, type UserRole, WebsiteSchema, type Website, WordReplacementSchema, type WordReplacement, WorkflowMessageSchema, type WorkflowMessage, WorldRegionSchema, type WorldRegion } from "./civicrm.schemas"

export function createACLClient(opts: CiviCRMClientOptions) {
  return {
    async get(params: {
      where?: WhereClause
      select?: (keyof ACL)[]
      orderBy?: OrderByClause
      limit?: number
      offset?: number
    } = {}): Promise<ACL[]> {
      const raw = await apiCall<unknown>(opts, "ACL", "get", params)
      return raw.map((v) => ACLSchema.parse(v))
    },

    async getById(id: number, select?: (keyof ACL)[]): Promise<ACL> {
      const results = await apiCall<unknown>(opts, "ACL", "get", {
        where: [["id", "=", id]],
        ...(select ? { select } : {}),
      })
      if (!results.length) throw new Error(`ACL ${id} not found`)
      return ACLSchema.parse(results[0])
    },

    async create(values: Partial<Omit<ACL, "id">>): Promise<ACL> {
      const results = await apiCall<unknown>(opts, "ACL", "create", { values })
      if (!results.length) throw new Error("No value returned from ACL.create")
      return ACLSchema.parse(results[0])
    },

    async update(id: number, values: Partial<Omit<ACL, "id">>): Promise<ACL> {
      const results = await apiCall<unknown>(opts, "ACL", "update", {
        where: [["id", "=", id]],
        values,
      })
      if (!results.length) throw new Error("No value returned from ACL.update")
      return ACLSchema.parse(results[0])
    },

    async delete(id: number): Promise<void> {
      await apiCall<unknown>(opts, "ACL", "delete", {
        where: [["id", "=", id]],
      })
    },
  }
}

export function createACLEntityRoleClient(opts: CiviCRMClientOptions) {
  return {
    async get(params: {
      where?: WhereClause
      select?: (keyof ACLEntityRole)[]
      orderBy?: OrderByClause
      limit?: number
      offset?: number
    } = {}): Promise<ACLEntityRole[]> {
      const raw = await apiCall<unknown>(opts, "ACLEntityRole", "get", params)
      return raw.map((v) => ACLEntityRoleSchema.parse(v))
    },

    async getById(id: number, select?: (keyof ACLEntityRole)[]): Promise<ACLEntityRole> {
      const results = await apiCall<unknown>(opts, "ACLEntityRole", "get", {
        where: [["id", "=", id]],
        ...(select ? { select } : {}),
      })
      if (!results.length) throw new Error(`ACLEntityRole ${id} not found`)
      return ACLEntityRoleSchema.parse(results[0])
    },

    async create(values: Partial<Omit<ACLEntityRole, "id">>): Promise<ACLEntityRole> {
      const results = await apiCall<unknown>(opts, "ACLEntityRole", "create", { values })
      if (!results.length) throw new Error("No value returned from ACLEntityRole.create")
      return ACLEntityRoleSchema.parse(results[0])
    },

    async update(id: number, values: Partial<Omit<ACLEntityRole, "id">>): Promise<ACLEntityRole> {
      const results = await apiCall<unknown>(opts, "ACLEntityRole", "update", {
        where: [["id", "=", id]],
        values,
      })
      if (!results.length) throw new Error("No value returned from ACLEntityRole.update")
      return ACLEntityRoleSchema.parse(results[0])
    },

    async delete(id: number): Promise<void> {
      await apiCall<unknown>(opts, "ACLEntityRole", "delete", {
        where: [["id", "=", id]],
      })
    },
  }
}

export function createActionScheduleClient(opts: CiviCRMClientOptions) {
  return {
    async get(params: {
      where?: WhereClause
      select?: (keyof ActionSchedule)[]
      orderBy?: OrderByClause
      limit?: number
      offset?: number
    } = {}): Promise<ActionSchedule[]> {
      const raw = await apiCall<unknown>(opts, "ActionSchedule", "get", params)
      return raw.map((v) => ActionScheduleSchema.parse(v))
    },

    async getById(id: number, select?: (keyof ActionSchedule)[]): Promise<ActionSchedule> {
      const results = await apiCall<unknown>(opts, "ActionSchedule", "get", {
        where: [["id", "=", id]],
        ...(select ? { select } : {}),
      })
      if (!results.length) throw new Error(`ActionSchedule ${id} not found`)
      return ActionScheduleSchema.parse(results[0])
    },

    async create(values: Partial<Omit<ActionSchedule, "id">>): Promise<ActionSchedule> {
      const results = await apiCall<unknown>(opts, "ActionSchedule", "create", { values })
      if (!results.length) throw new Error("No value returned from ActionSchedule.create")
      return ActionScheduleSchema.parse(results[0])
    },

    async update(id: number, values: Partial<Omit<ActionSchedule, "id">>): Promise<ActionSchedule> {
      const results = await apiCall<unknown>(opts, "ActionSchedule", "update", {
        where: [["id", "=", id]],
        values,
      })
      if (!results.length) throw new Error("No value returned from ActionSchedule.update")
      return ActionScheduleSchema.parse(results[0])
    },

    async delete(id: number): Promise<void> {
      await apiCall<unknown>(opts, "ActionSchedule", "delete", {
        where: [["id", "=", id]],
      })
    },
  }
}

export function createActivityClient(opts: CiviCRMClientOptions) {
  return {
    async get(params: {
      where?: WhereClause
      select?: (keyof Activity)[]
      orderBy?: OrderByClause
      limit?: number
      offset?: number
    } = {}): Promise<Activity[]> {
      const raw = await apiCall<unknown>(opts, "Activity", "get", params)
      return raw.map((v) => ActivitySchema.parse(v))
    },

    async getById(id: number, select?: (keyof Activity)[]): Promise<Activity> {
      const results = await apiCall<unknown>(opts, "Activity", "get", {
        where: [["id", "=", id]],
        ...(select ? { select } : {}),
      })
      if (!results.length) throw new Error(`Activity ${id} not found`)
      return ActivitySchema.parse(results[0])
    },

    async create(values: Partial<Omit<Activity, "id">>): Promise<Activity> {
      const results = await apiCall<unknown>(opts, "Activity", "create", { values })
      if (!results.length) throw new Error("No value returned from Activity.create")
      return ActivitySchema.parse(results[0])
    },

    async update(id: number, values: Partial<Omit<Activity, "id">>): Promise<Activity> {
      const results = await apiCall<unknown>(opts, "Activity", "update", {
        where: [["id", "=", id]],
        values,
      })
      if (!results.length) throw new Error("No value returned from Activity.update")
      return ActivitySchema.parse(results[0])
    },

    async delete(id: number): Promise<void> {
      await apiCall<unknown>(opts, "Activity", "delete", {
        where: [["id", "=", id]],
      })
    },
  }
}

export function createActivityContactClient(opts: CiviCRMClientOptions) {
  return {
    async get(params: {
      where?: WhereClause
      select?: (keyof ActivityContact)[]
      orderBy?: OrderByClause
      limit?: number
      offset?: number
    } = {}): Promise<ActivityContact[]> {
      const raw = await apiCall<unknown>(opts, "ActivityContact", "get", params)
      return raw.map((v) => ActivityContactSchema.parse(v))
    },

    async getById(id: number, select?: (keyof ActivityContact)[]): Promise<ActivityContact> {
      const results = await apiCall<unknown>(opts, "ActivityContact", "get", {
        where: [["id", "=", id]],
        ...(select ? { select } : {}),
      })
      if (!results.length) throw new Error(`ActivityContact ${id} not found`)
      return ActivityContactSchema.parse(results[0])
    },

    async create(values: Partial<Omit<ActivityContact, "id">>): Promise<ActivityContact> {
      const results = await apiCall<unknown>(opts, "ActivityContact", "create", { values })
      if (!results.length) throw new Error("No value returned from ActivityContact.create")
      return ActivityContactSchema.parse(results[0])
    },

    async update(id: number, values: Partial<Omit<ActivityContact, "id">>): Promise<ActivityContact> {
      const results = await apiCall<unknown>(opts, "ActivityContact", "update", {
        where: [["id", "=", id]],
        values,
      })
      if (!results.length) throw new Error("No value returned from ActivityContact.update")
      return ActivityContactSchema.parse(results[0])
    },

    async delete(id: number): Promise<void> {
      await apiCall<unknown>(opts, "ActivityContact", "delete", {
        where: [["id", "=", id]],
      })
    },
  }
}

export function createAddressClient(opts: CiviCRMClientOptions) {
  return {
    async get(params: {
      where?: WhereClause
      select?: (keyof Address)[]
      orderBy?: OrderByClause
      limit?: number
      offset?: number
    } = {}): Promise<Address[]> {
      const raw = await apiCall<unknown>(opts, "Address", "get", params)
      return raw.map((v) => AddressSchema.parse(v))
    },

    async getById(id: number, select?: (keyof Address)[]): Promise<Address> {
      const results = await apiCall<unknown>(opts, "Address", "get", {
        where: [["id", "=", id]],
        ...(select ? { select } : {}),
      })
      if (!results.length) throw new Error(`Address ${id} not found`)
      return AddressSchema.parse(results[0])
    },

    async create(values: Partial<Omit<Address, "id">>): Promise<Address> {
      const results = await apiCall<unknown>(opts, "Address", "create", { values })
      if (!results.length) throw new Error("No value returned from Address.create")
      return AddressSchema.parse(results[0])
    },

    async update(id: number, values: Partial<Omit<Address, "id">>): Promise<Address> {
      const results = await apiCall<unknown>(opts, "Address", "update", {
        where: [["id", "=", id]],
        values,
      })
      if (!results.length) throw new Error("No value returned from Address.update")
      return AddressSchema.parse(results[0])
    },

    async delete(id: number): Promise<void> {
      await apiCall<unknown>(opts, "Address", "delete", {
        where: [["id", "=", id]],
      })
    },
  }
}

export function createAfformClient(opts: CiviCRMClientOptions) {
  return {
    async get(params: {
      where?: WhereClause
      select?: (keyof Afform)[]
      orderBy?: OrderByClause
      limit?: number
      offset?: number
    } = {}): Promise<Afform[]> {
      const raw = await apiCall<unknown>(opts, "Afform", "get", params)
      return raw.map((v) => AfformSchema.parse(v))
    },

    async getById(id: number, select?: (keyof Afform)[]): Promise<Afform> {
      const results = await apiCall<unknown>(opts, "Afform", "get", {
        where: [["id", "=", id]],
        ...(select ? { select } : {}),
      })
      if (!results.length) throw new Error(`Afform ${id} not found`)
      return AfformSchema.parse(results[0])
    },

    async create(values: Partial<Omit<Afform, "id">>): Promise<Afform> {
      const results = await apiCall<unknown>(opts, "Afform", "create", { values })
      if (!results.length) throw new Error("No value returned from Afform.create")
      return AfformSchema.parse(results[0])
    },

    async update(id: number, values: Partial<Omit<Afform, "id">>): Promise<Afform> {
      const results = await apiCall<unknown>(opts, "Afform", "update", {
        where: [["id", "=", id]],
        values,
      })
      if (!results.length) throw new Error("No value returned from Afform.update")
      return AfformSchema.parse(results[0])
    },

    async delete(id: number): Promise<void> {
      await apiCall<unknown>(opts, "Afform", "delete", {
        where: [["id", "=", id]],
      })
    },
  }
}

export function createAfformBehaviorClient(opts: CiviCRMClientOptions) {
  return {
    async get(params: {
      where?: WhereClause
      select?: (keyof AfformBehavior)[]
      orderBy?: OrderByClause
      limit?: number
      offset?: number
    } = {}): Promise<AfformBehavior[]> {
      const raw = await apiCall<unknown>(opts, "AfformBehavior", "get", params)
      return raw.map((v) => AfformBehaviorSchema.parse(v))
    },

    async getById(id: number, select?: (keyof AfformBehavior)[]): Promise<AfformBehavior> {
      const results = await apiCall<unknown>(opts, "AfformBehavior", "get", {
        where: [["id", "=", id]],
        ...(select ? { select } : {}),
      })
      if (!results.length) throw new Error(`AfformBehavior ${id} not found`)
      return AfformBehaviorSchema.parse(results[0])
    },

    async create(values: Partial<Omit<AfformBehavior, "id">>): Promise<AfformBehavior> {
      const results = await apiCall<unknown>(opts, "AfformBehavior", "create", { values })
      if (!results.length) throw new Error("No value returned from AfformBehavior.create")
      return AfformBehaviorSchema.parse(results[0])
    },

    async update(id: number, values: Partial<Omit<AfformBehavior, "id">>): Promise<AfformBehavior> {
      const results = await apiCall<unknown>(opts, "AfformBehavior", "update", {
        where: [["id", "=", id]],
        values,
      })
      if (!results.length) throw new Error("No value returned from AfformBehavior.update")
      return AfformBehaviorSchema.parse(results[0])
    },

    async delete(id: number): Promise<void> {
      await apiCall<unknown>(opts, "AfformBehavior", "delete", {
        where: [["id", "=", id]],
      })
    },
  }
}

export function createAfformSubmissionClient(opts: CiviCRMClientOptions) {
  return {
    async get(params: {
      where?: WhereClause
      select?: (keyof AfformSubmission)[]
      orderBy?: OrderByClause
      limit?: number
      offset?: number
    } = {}): Promise<AfformSubmission[]> {
      const raw = await apiCall<unknown>(opts, "AfformSubmission", "get", params)
      return raw.map((v) => AfformSubmissionSchema.parse(v))
    },

    async getById(id: number, select?: (keyof AfformSubmission)[]): Promise<AfformSubmission> {
      const results = await apiCall<unknown>(opts, "AfformSubmission", "get", {
        where: [["id", "=", id]],
        ...(select ? { select } : {}),
      })
      if (!results.length) throw new Error(`AfformSubmission ${id} not found`)
      return AfformSubmissionSchema.parse(results[0])
    },

    async create(values: Partial<Omit<AfformSubmission, "id">>): Promise<AfformSubmission> {
      const results = await apiCall<unknown>(opts, "AfformSubmission", "create", { values })
      if (!results.length) throw new Error("No value returned from AfformSubmission.create")
      return AfformSubmissionSchema.parse(results[0])
    },

    async update(id: number, values: Partial<Omit<AfformSubmission, "id">>): Promise<AfformSubmission> {
      const results = await apiCall<unknown>(opts, "AfformSubmission", "update", {
        where: [["id", "=", id]],
        values,
      })
      if (!results.length) throw new Error("No value returned from AfformSubmission.update")
      return AfformSubmissionSchema.parse(results[0])
    },

    async delete(id: number): Promise<void> {
      await apiCall<unknown>(opts, "AfformSubmission", "delete", {
        where: [["id", "=", id]],
      })
    },
  }
}

export function createBatchClient(opts: CiviCRMClientOptions) {
  return {
    async get(params: {
      where?: WhereClause
      select?: (keyof Batch)[]
      orderBy?: OrderByClause
      limit?: number
      offset?: number
    } = {}): Promise<Batch[]> {
      const raw = await apiCall<unknown>(opts, "Batch", "get", params)
      return raw.map((v) => BatchSchema.parse(v))
    },

    async getById(id: number, select?: (keyof Batch)[]): Promise<Batch> {
      const results = await apiCall<unknown>(opts, "Batch", "get", {
        where: [["id", "=", id]],
        ...(select ? { select } : {}),
      })
      if (!results.length) throw new Error(`Batch ${id} not found`)
      return BatchSchema.parse(results[0])
    },

    async create(values: Partial<Omit<Batch, "id">>): Promise<Batch> {
      const results = await apiCall<unknown>(opts, "Batch", "create", { values })
      if (!results.length) throw new Error("No value returned from Batch.create")
      return BatchSchema.parse(results[0])
    },

    async update(id: number, values: Partial<Omit<Batch, "id">>): Promise<Batch> {
      const results = await apiCall<unknown>(opts, "Batch", "update", {
        where: [["id", "=", id]],
        values,
      })
      if (!results.length) throw new Error("No value returned from Batch.update")
      return BatchSchema.parse(results[0])
    },

    async delete(id: number): Promise<void> {
      await apiCall<unknown>(opts, "Batch", "delete", {
        where: [["id", "=", id]],
      })
    },
  }
}

export function createBouncePatternClient(opts: CiviCRMClientOptions) {
  return {
    async get(params: {
      where?: WhereClause
      select?: (keyof BouncePattern)[]
      orderBy?: OrderByClause
      limit?: number
      offset?: number
    } = {}): Promise<BouncePattern[]> {
      const raw = await apiCall<unknown>(opts, "BouncePattern", "get", params)
      return raw.map((v) => BouncePatternSchema.parse(v))
    },

    async getById(id: number, select?: (keyof BouncePattern)[]): Promise<BouncePattern> {
      const results = await apiCall<unknown>(opts, "BouncePattern", "get", {
        where: [["id", "=", id]],
        ...(select ? { select } : {}),
      })
      if (!results.length) throw new Error(`BouncePattern ${id} not found`)
      return BouncePatternSchema.parse(results[0])
    },

    async create(values: Partial<Omit<BouncePattern, "id">>): Promise<BouncePattern> {
      const results = await apiCall<unknown>(opts, "BouncePattern", "create", { values })
      if (!results.length) throw new Error("No value returned from BouncePattern.create")
      return BouncePatternSchema.parse(results[0])
    },

    async update(id: number, values: Partial<Omit<BouncePattern, "id">>): Promise<BouncePattern> {
      const results = await apiCall<unknown>(opts, "BouncePattern", "update", {
        where: [["id", "=", id]],
        values,
      })
      if (!results.length) throw new Error("No value returned from BouncePattern.update")
      return BouncePatternSchema.parse(results[0])
    },

    async delete(id: number): Promise<void> {
      await apiCall<unknown>(opts, "BouncePattern", "delete", {
        where: [["id", "=", id]],
      })
    },
  }
}

export function createBounceTypeClient(opts: CiviCRMClientOptions) {
  return {
    async get(params: {
      where?: WhereClause
      select?: (keyof BounceType)[]
      orderBy?: OrderByClause
      limit?: number
      offset?: number
    } = {}): Promise<BounceType[]> {
      const raw = await apiCall<unknown>(opts, "BounceType", "get", params)
      return raw.map((v) => BounceTypeSchema.parse(v))
    },

    async getById(id: number, select?: (keyof BounceType)[]): Promise<BounceType> {
      const results = await apiCall<unknown>(opts, "BounceType", "get", {
        where: [["id", "=", id]],
        ...(select ? { select } : {}),
      })
      if (!results.length) throw new Error(`BounceType ${id} not found`)
      return BounceTypeSchema.parse(results[0])
    },

    async create(values: Partial<Omit<BounceType, "id">>): Promise<BounceType> {
      const results = await apiCall<unknown>(opts, "BounceType", "create", { values })
      if (!results.length) throw new Error("No value returned from BounceType.create")
      return BounceTypeSchema.parse(results[0])
    },

    async update(id: number, values: Partial<Omit<BounceType, "id">>): Promise<BounceType> {
      const results = await apiCall<unknown>(opts, "BounceType", "update", {
        where: [["id", "=", id]],
        values,
      })
      if (!results.length) throw new Error("No value returned from BounceType.update")
      return BounceTypeSchema.parse(results[0])
    },

    async delete(id: number): Promise<void> {
      await apiCall<unknown>(opts, "BounceType", "delete", {
        where: [["id", "=", id]],
      })
    },
  }
}

export function createCaseClient(opts: CiviCRMClientOptions) {
  return {
    async get(params: {
      where?: WhereClause
      select?: (keyof Case)[]
      orderBy?: OrderByClause
      limit?: number
      offset?: number
    } = {}): Promise<Case[]> {
      const raw = await apiCall<unknown>(opts, "Case", "get", params)
      return raw.map((v) => CaseSchema.parse(v))
    },

    async getById(id: number, select?: (keyof Case)[]): Promise<Case> {
      const results = await apiCall<unknown>(opts, "Case", "get", {
        where: [["id", "=", id]],
        ...(select ? { select } : {}),
      })
      if (!results.length) throw new Error(`Case ${id} not found`)
      return CaseSchema.parse(results[0])
    },

    async create(values: Partial<Omit<Case, "id">>): Promise<Case> {
      const results = await apiCall<unknown>(opts, "Case", "create", { values })
      if (!results.length) throw new Error("No value returned from Case.create")
      return CaseSchema.parse(results[0])
    },

    async update(id: number, values: Partial<Omit<Case, "id">>): Promise<Case> {
      const results = await apiCall<unknown>(opts, "Case", "update", {
        where: [["id", "=", id]],
        values,
      })
      if (!results.length) throw new Error("No value returned from Case.update")
      return CaseSchema.parse(results[0])
    },

    async delete(id: number): Promise<void> {
      await apiCall<unknown>(opts, "Case", "delete", {
        where: [["id", "=", id]],
      })
    },
  }
}

export function createCaseActivityClient(opts: CiviCRMClientOptions) {
  return {
    async get(params: {
      where?: WhereClause
      select?: (keyof CaseActivity)[]
      orderBy?: OrderByClause
      limit?: number
      offset?: number
    } = {}): Promise<CaseActivity[]> {
      const raw = await apiCall<unknown>(opts, "CaseActivity", "get", params)
      return raw.map((v) => CaseActivitySchema.parse(v))
    },

    async getById(id: number, select?: (keyof CaseActivity)[]): Promise<CaseActivity> {
      const results = await apiCall<unknown>(opts, "CaseActivity", "get", {
        where: [["id", "=", id]],
        ...(select ? { select } : {}),
      })
      if (!results.length) throw new Error(`CaseActivity ${id} not found`)
      return CaseActivitySchema.parse(results[0])
    },

    async create(values: Partial<Omit<CaseActivity, "id">>): Promise<CaseActivity> {
      const results = await apiCall<unknown>(opts, "CaseActivity", "create", { values })
      if (!results.length) throw new Error("No value returned from CaseActivity.create")
      return CaseActivitySchema.parse(results[0])
    },

    async update(id: number, values: Partial<Omit<CaseActivity, "id">>): Promise<CaseActivity> {
      const results = await apiCall<unknown>(opts, "CaseActivity", "update", {
        where: [["id", "=", id]],
        values,
      })
      if (!results.length) throw new Error("No value returned from CaseActivity.update")
      return CaseActivitySchema.parse(results[0])
    },

    async delete(id: number): Promise<void> {
      await apiCall<unknown>(opts, "CaseActivity", "delete", {
        where: [["id", "=", id]],
      })
    },
  }
}

export function createCaseContactClient(opts: CiviCRMClientOptions) {
  return {
    async get(params: {
      where?: WhereClause
      select?: (keyof CaseContact)[]
      orderBy?: OrderByClause
      limit?: number
      offset?: number
    } = {}): Promise<CaseContact[]> {
      const raw = await apiCall<unknown>(opts, "CaseContact", "get", params)
      return raw.map((v) => CaseContactSchema.parse(v))
    },

    async getById(id: number, select?: (keyof CaseContact)[]): Promise<CaseContact> {
      const results = await apiCall<unknown>(opts, "CaseContact", "get", {
        where: [["id", "=", id]],
        ...(select ? { select } : {}),
      })
      if (!results.length) throw new Error(`CaseContact ${id} not found`)
      return CaseContactSchema.parse(results[0])
    },

    async create(values: Partial<Omit<CaseContact, "id">>): Promise<CaseContact> {
      const results = await apiCall<unknown>(opts, "CaseContact", "create", { values })
      if (!results.length) throw new Error("No value returned from CaseContact.create")
      return CaseContactSchema.parse(results[0])
    },

    async update(id: number, values: Partial<Omit<CaseContact, "id">>): Promise<CaseContact> {
      const results = await apiCall<unknown>(opts, "CaseContact", "update", {
        where: [["id", "=", id]],
        values,
      })
      if (!results.length) throw new Error("No value returned from CaseContact.update")
      return CaseContactSchema.parse(results[0])
    },

    async delete(id: number): Promise<void> {
      await apiCall<unknown>(opts, "CaseContact", "delete", {
        where: [["id", "=", id]],
      })
    },
  }
}

export function createCaseTypeClient(opts: CiviCRMClientOptions) {
  return {
    async get(params: {
      where?: WhereClause
      select?: (keyof CaseType)[]
      orderBy?: OrderByClause
      limit?: number
      offset?: number
    } = {}): Promise<CaseType[]> {
      const raw = await apiCall<unknown>(opts, "CaseType", "get", params)
      return raw.map((v) => CaseTypeSchema.parse(v))
    },

    async getById(id: number, select?: (keyof CaseType)[]): Promise<CaseType> {
      const results = await apiCall<unknown>(opts, "CaseType", "get", {
        where: [["id", "=", id]],
        ...(select ? { select } : {}),
      })
      if (!results.length) throw new Error(`CaseType ${id} not found`)
      return CaseTypeSchema.parse(results[0])
    },

    async create(values: Partial<Omit<CaseType, "id">>): Promise<CaseType> {
      const results = await apiCall<unknown>(opts, "CaseType", "create", { values })
      if (!results.length) throw new Error("No value returned from CaseType.create")
      return CaseTypeSchema.parse(results[0])
    },

    async update(id: number, values: Partial<Omit<CaseType, "id">>): Promise<CaseType> {
      const results = await apiCall<unknown>(opts, "CaseType", "update", {
        where: [["id", "=", id]],
        values,
      })
      if (!results.length) throw new Error("No value returned from CaseType.update")
      return CaseTypeSchema.parse(results[0])
    },

    async delete(id: number): Promise<void> {
      await apiCall<unknown>(opts, "CaseType", "delete", {
        where: [["id", "=", id]],
      })
    },
  }
}

export function createContactClient(opts: CiviCRMClientOptions) {
  return {
    async get(params: {
      where?: WhereClause
      select?: (keyof Contact)[]
      orderBy?: OrderByClause
      limit?: number
      offset?: number
    } = {}): Promise<Contact[]> {
      const raw = await apiCall<unknown>(opts, "Contact", "get", params)
      return raw.map((v) => ContactSchema.parse(v))
    },

    async getById(id: number, select?: (keyof Contact)[]): Promise<Contact> {
      const results = await apiCall<unknown>(opts, "Contact", "get", {
        where: [["id", "=", id]],
        ...(select ? { select } : {}),
      })
      if (!results.length) throw new Error(`Contact ${id} not found`)
      return ContactSchema.parse(results[0])
    },

    async create(values: Partial<Omit<Contact, "id">>): Promise<Contact> {
      const results = await apiCall<unknown>(opts, "Contact", "create", { values })
      if (!results.length) throw new Error("No value returned from Contact.create")
      return ContactSchema.parse(results[0])
    },

    async update(id: number, values: Partial<Omit<Contact, "id">>): Promise<Contact> {
      const results = await apiCall<unknown>(opts, "Contact", "update", {
        where: [["id", "=", id]],
        values,
      })
      if (!results.length) throw new Error("No value returned from Contact.update")
      return ContactSchema.parse(results[0])
    },

    async delete(id: number): Promise<void> {
      await apiCall<unknown>(opts, "Contact", "delete", {
        where: [["id", "=", id]],
      })
    },
  }
}

export function createContactTypeClient(opts: CiviCRMClientOptions) {
  return {
    async get(params: {
      where?: WhereClause
      select?: (keyof ContactType)[]
      orderBy?: OrderByClause
      limit?: number
      offset?: number
    } = {}): Promise<ContactType[]> {
      const raw = await apiCall<unknown>(opts, "ContactType", "get", params)
      return raw.map((v) => ContactTypeSchema.parse(v))
    },

    async getById(id: number, select?: (keyof ContactType)[]): Promise<ContactType> {
      const results = await apiCall<unknown>(opts, "ContactType", "get", {
        where: [["id", "=", id]],
        ...(select ? { select } : {}),
      })
      if (!results.length) throw new Error(`ContactType ${id} not found`)
      return ContactTypeSchema.parse(results[0])
    },

    async create(values: Partial<Omit<ContactType, "id">>): Promise<ContactType> {
      const results = await apiCall<unknown>(opts, "ContactType", "create", { values })
      if (!results.length) throw new Error("No value returned from ContactType.create")
      return ContactTypeSchema.parse(results[0])
    },

    async update(id: number, values: Partial<Omit<ContactType, "id">>): Promise<ContactType> {
      const results = await apiCall<unknown>(opts, "ContactType", "update", {
        where: [["id", "=", id]],
        values,
      })
      if (!results.length) throw new Error("No value returned from ContactType.update")
      return ContactTypeSchema.parse(results[0])
    },

    async delete(id: number): Promise<void> {
      await apiCall<unknown>(opts, "ContactType", "delete", {
        where: [["id", "=", id]],
      })
    },
  }
}

export function createContributionClient(opts: CiviCRMClientOptions) {
  return {
    async get(params: {
      where?: WhereClause
      select?: (keyof Contribution)[]
      orderBy?: OrderByClause
      limit?: number
      offset?: number
    } = {}): Promise<Contribution[]> {
      const raw = await apiCall<unknown>(opts, "Contribution", "get", params)
      return raw.map((v) => ContributionSchema.parse(v))
    },

    async getById(id: number, select?: (keyof Contribution)[]): Promise<Contribution> {
      const results = await apiCall<unknown>(opts, "Contribution", "get", {
        where: [["id", "=", id]],
        ...(select ? { select } : {}),
      })
      if (!results.length) throw new Error(`Contribution ${id} not found`)
      return ContributionSchema.parse(results[0])
    },

    async create(values: Partial<Omit<Contribution, "id">>): Promise<Contribution> {
      const results = await apiCall<unknown>(opts, "Contribution", "create", { values })
      if (!results.length) throw new Error("No value returned from Contribution.create")
      return ContributionSchema.parse(results[0])
    },

    async update(id: number, values: Partial<Omit<Contribution, "id">>): Promise<Contribution> {
      const results = await apiCall<unknown>(opts, "Contribution", "update", {
        where: [["id", "=", id]],
        values,
      })
      if (!results.length) throw new Error("No value returned from Contribution.update")
      return ContributionSchema.parse(results[0])
    },

    async delete(id: number): Promise<void> {
      await apiCall<unknown>(opts, "Contribution", "delete", {
        where: [["id", "=", id]],
      })
    },
  }
}

export function createContributionPageClient(opts: CiviCRMClientOptions) {
  return {
    async get(params: {
      where?: WhereClause
      select?: (keyof ContributionPage)[]
      orderBy?: OrderByClause
      limit?: number
      offset?: number
    } = {}): Promise<ContributionPage[]> {
      const raw = await apiCall<unknown>(opts, "ContributionPage", "get", params)
      return raw.map((v) => ContributionPageSchema.parse(v))
    },

    async getById(id: number, select?: (keyof ContributionPage)[]): Promise<ContributionPage> {
      const results = await apiCall<unknown>(opts, "ContributionPage", "get", {
        where: [["id", "=", id]],
        ...(select ? { select } : {}),
      })
      if (!results.length) throw new Error(`ContributionPage ${id} not found`)
      return ContributionPageSchema.parse(results[0])
    },

    async create(values: Partial<Omit<ContributionPage, "id">>): Promise<ContributionPage> {
      const results = await apiCall<unknown>(opts, "ContributionPage", "create", { values })
      if (!results.length) throw new Error("No value returned from ContributionPage.create")
      return ContributionPageSchema.parse(results[0])
    },

    async update(id: number, values: Partial<Omit<ContributionPage, "id">>): Promise<ContributionPage> {
      const results = await apiCall<unknown>(opts, "ContributionPage", "update", {
        where: [["id", "=", id]],
        values,
      })
      if (!results.length) throw new Error("No value returned from ContributionPage.update")
      return ContributionPageSchema.parse(results[0])
    },

    async delete(id: number): Promise<void> {
      await apiCall<unknown>(opts, "ContributionPage", "delete", {
        where: [["id", "=", id]],
      })
    },
  }
}

export function createContributionProductClient(opts: CiviCRMClientOptions) {
  return {
    async get(params: {
      where?: WhereClause
      select?: (keyof ContributionProduct)[]
      orderBy?: OrderByClause
      limit?: number
      offset?: number
    } = {}): Promise<ContributionProduct[]> {
      const raw = await apiCall<unknown>(opts, "ContributionProduct", "get", params)
      return raw.map((v) => ContributionProductSchema.parse(v))
    },

    async getById(id: number, select?: (keyof ContributionProduct)[]): Promise<ContributionProduct> {
      const results = await apiCall<unknown>(opts, "ContributionProduct", "get", {
        where: [["id", "=", id]],
        ...(select ? { select } : {}),
      })
      if (!results.length) throw new Error(`ContributionProduct ${id} not found`)
      return ContributionProductSchema.parse(results[0])
    },

    async create(values: Partial<Omit<ContributionProduct, "id">>): Promise<ContributionProduct> {
      const results = await apiCall<unknown>(opts, "ContributionProduct", "create", { values })
      if (!results.length) throw new Error("No value returned from ContributionProduct.create")
      return ContributionProductSchema.parse(results[0])
    },

    async update(id: number, values: Partial<Omit<ContributionProduct, "id">>): Promise<ContributionProduct> {
      const results = await apiCall<unknown>(opts, "ContributionProduct", "update", {
        where: [["id", "=", id]],
        values,
      })
      if (!results.length) throw new Error("No value returned from ContributionProduct.update")
      return ContributionProductSchema.parse(results[0])
    },

    async delete(id: number): Promise<void> {
      await apiCall<unknown>(opts, "ContributionProduct", "delete", {
        where: [["id", "=", id]],
      })
    },
  }
}

export function createContributionRecurClient(opts: CiviCRMClientOptions) {
  return {
    async get(params: {
      where?: WhereClause
      select?: (keyof ContributionRecur)[]
      orderBy?: OrderByClause
      limit?: number
      offset?: number
    } = {}): Promise<ContributionRecur[]> {
      const raw = await apiCall<unknown>(opts, "ContributionRecur", "get", params)
      return raw.map((v) => ContributionRecurSchema.parse(v))
    },

    async getById(id: number, select?: (keyof ContributionRecur)[]): Promise<ContributionRecur> {
      const results = await apiCall<unknown>(opts, "ContributionRecur", "get", {
        where: [["id", "=", id]],
        ...(select ? { select } : {}),
      })
      if (!results.length) throw new Error(`ContributionRecur ${id} not found`)
      return ContributionRecurSchema.parse(results[0])
    },

    async create(values: Partial<Omit<ContributionRecur, "id">>): Promise<ContributionRecur> {
      const results = await apiCall<unknown>(opts, "ContributionRecur", "create", { values })
      if (!results.length) throw new Error("No value returned from ContributionRecur.create")
      return ContributionRecurSchema.parse(results[0])
    },

    async update(id: number, values: Partial<Omit<ContributionRecur, "id">>): Promise<ContributionRecur> {
      const results = await apiCall<unknown>(opts, "ContributionRecur", "update", {
        where: [["id", "=", id]],
        values,
      })
      if (!results.length) throw new Error("No value returned from ContributionRecur.update")
      return ContributionRecurSchema.parse(results[0])
    },

    async delete(id: number): Promise<void> {
      await apiCall<unknown>(opts, "ContributionRecur", "delete", {
        where: [["id", "=", id]],
      })
    },
  }
}

export function createContributionSoftClient(opts: CiviCRMClientOptions) {
  return {
    async get(params: {
      where?: WhereClause
      select?: (keyof ContributionSoft)[]
      orderBy?: OrderByClause
      limit?: number
      offset?: number
    } = {}): Promise<ContributionSoft[]> {
      const raw = await apiCall<unknown>(opts, "ContributionSoft", "get", params)
      return raw.map((v) => ContributionSoftSchema.parse(v))
    },

    async getById(id: number, select?: (keyof ContributionSoft)[]): Promise<ContributionSoft> {
      const results = await apiCall<unknown>(opts, "ContributionSoft", "get", {
        where: [["id", "=", id]],
        ...(select ? { select } : {}),
      })
      if (!results.length) throw new Error(`ContributionSoft ${id} not found`)
      return ContributionSoftSchema.parse(results[0])
    },

    async create(values: Partial<Omit<ContributionSoft, "id">>): Promise<ContributionSoft> {
      const results = await apiCall<unknown>(opts, "ContributionSoft", "create", { values })
      if (!results.length) throw new Error("No value returned from ContributionSoft.create")
      return ContributionSoftSchema.parse(results[0])
    },

    async update(id: number, values: Partial<Omit<ContributionSoft, "id">>): Promise<ContributionSoft> {
      const results = await apiCall<unknown>(opts, "ContributionSoft", "update", {
        where: [["id", "=", id]],
        values,
      })
      if (!results.length) throw new Error("No value returned from ContributionSoft.update")
      return ContributionSoftSchema.parse(results[0])
    },

    async delete(id: number): Promise<void> {
      await apiCall<unknown>(opts, "ContributionSoft", "delete", {
        where: [["id", "=", id]],
      })
    },
  }
}

export function createCountryClient(opts: CiviCRMClientOptions) {
  return {
    async get(params: {
      where?: WhereClause
      select?: (keyof Country)[]
      orderBy?: OrderByClause
      limit?: number
      offset?: number
    } = {}): Promise<Country[]> {
      const raw = await apiCall<unknown>(opts, "Country", "get", params)
      return raw.map((v) => CountrySchema.parse(v))
    },

    async getById(id: number, select?: (keyof Country)[]): Promise<Country> {
      const results = await apiCall<unknown>(opts, "Country", "get", {
        where: [["id", "=", id]],
        ...(select ? { select } : {}),
      })
      if (!results.length) throw new Error(`Country ${id} not found`)
      return CountrySchema.parse(results[0])
    },

    async create(values: Partial<Omit<Country, "id">>): Promise<Country> {
      const results = await apiCall<unknown>(opts, "Country", "create", { values })
      if (!results.length) throw new Error("No value returned from Country.create")
      return CountrySchema.parse(results[0])
    },

    async update(id: number, values: Partial<Omit<Country, "id">>): Promise<Country> {
      const results = await apiCall<unknown>(opts, "Country", "update", {
        where: [["id", "=", id]],
        values,
      })
      if (!results.length) throw new Error("No value returned from Country.update")
      return CountrySchema.parse(results[0])
    },

    async delete(id: number): Promise<void> {
      await apiCall<unknown>(opts, "Country", "delete", {
        where: [["id", "=", id]],
      })
    },
  }
}

export function createCountyClient(opts: CiviCRMClientOptions) {
  return {
    async get(params: {
      where?: WhereClause
      select?: (keyof County)[]
      orderBy?: OrderByClause
      limit?: number
      offset?: number
    } = {}): Promise<County[]> {
      const raw = await apiCall<unknown>(opts, "County", "get", params)
      return raw.map((v) => CountySchema.parse(v))
    },

    async getById(id: number, select?: (keyof County)[]): Promise<County> {
      const results = await apiCall<unknown>(opts, "County", "get", {
        where: [["id", "=", id]],
        ...(select ? { select } : {}),
      })
      if (!results.length) throw new Error(`County ${id} not found`)
      return CountySchema.parse(results[0])
    },

    async create(values: Partial<Omit<County, "id">>): Promise<County> {
      const results = await apiCall<unknown>(opts, "County", "create", { values })
      if (!results.length) throw new Error("No value returned from County.create")
      return CountySchema.parse(results[0])
    },

    async update(id: number, values: Partial<Omit<County, "id">>): Promise<County> {
      const results = await apiCall<unknown>(opts, "County", "update", {
        where: [["id", "=", id]],
        values,
      })
      if (!results.length) throw new Error("No value returned from County.update")
      return CountySchema.parse(results[0])
    },

    async delete(id: number): Promise<void> {
      await apiCall<unknown>(opts, "County", "delete", {
        where: [["id", "=", id]],
      })
    },
  }
}

export function createCustomFieldClient(opts: CiviCRMClientOptions) {
  return {
    async get(params: {
      where?: WhereClause
      select?: (keyof CustomField)[]
      orderBy?: OrderByClause
      limit?: number
      offset?: number
    } = {}): Promise<CustomField[]> {
      const raw = await apiCall<unknown>(opts, "CustomField", "get", params)
      return raw.map((v) => CustomFieldSchema.parse(v))
    },

    async getById(id: number, select?: (keyof CustomField)[]): Promise<CustomField> {
      const results = await apiCall<unknown>(opts, "CustomField", "get", {
        where: [["id", "=", id]],
        ...(select ? { select } : {}),
      })
      if (!results.length) throw new Error(`CustomField ${id} not found`)
      return CustomFieldSchema.parse(results[0])
    },

    async create(values: Partial<Omit<CustomField, "id">>): Promise<CustomField> {
      const results = await apiCall<unknown>(opts, "CustomField", "create", { values })
      if (!results.length) throw new Error("No value returned from CustomField.create")
      return CustomFieldSchema.parse(results[0])
    },

    async update(id: number, values: Partial<Omit<CustomField, "id">>): Promise<CustomField> {
      const results = await apiCall<unknown>(opts, "CustomField", "update", {
        where: [["id", "=", id]],
        values,
      })
      if (!results.length) throw new Error("No value returned from CustomField.update")
      return CustomFieldSchema.parse(results[0])
    },

    async delete(id: number): Promise<void> {
      await apiCall<unknown>(opts, "CustomField", "delete", {
        where: [["id", "=", id]],
      })
    },
  }
}

export function createCustomGroupClient(opts: CiviCRMClientOptions) {
  return {
    async get(params: {
      where?: WhereClause
      select?: (keyof CustomGroup)[]
      orderBy?: OrderByClause
      limit?: number
      offset?: number
    } = {}): Promise<CustomGroup[]> {
      const raw = await apiCall<unknown>(opts, "CustomGroup", "get", params)
      return raw.map((v) => CustomGroupSchema.parse(v))
    },

    async getById(id: number, select?: (keyof CustomGroup)[]): Promise<CustomGroup> {
      const results = await apiCall<unknown>(opts, "CustomGroup", "get", {
        where: [["id", "=", id]],
        ...(select ? { select } : {}),
      })
      if (!results.length) throw new Error(`CustomGroup ${id} not found`)
      return CustomGroupSchema.parse(results[0])
    },

    async create(values: Partial<Omit<CustomGroup, "id">>): Promise<CustomGroup> {
      const results = await apiCall<unknown>(opts, "CustomGroup", "create", { values })
      if (!results.length) throw new Error("No value returned from CustomGroup.create")
      return CustomGroupSchema.parse(results[0])
    },

    async update(id: number, values: Partial<Omit<CustomGroup, "id">>): Promise<CustomGroup> {
      const results = await apiCall<unknown>(opts, "CustomGroup", "update", {
        where: [["id", "=", id]],
        values,
      })
      if (!results.length) throw new Error("No value returned from CustomGroup.update")
      return CustomGroupSchema.parse(results[0])
    },

    async delete(id: number): Promise<void> {
      await apiCall<unknown>(opts, "CustomGroup", "delete", {
        where: [["id", "=", id]],
      })
    },
  }
}

export function createDashboardClient(opts: CiviCRMClientOptions) {
  return {
    async get(params: {
      where?: WhereClause
      select?: (keyof Dashboard)[]
      orderBy?: OrderByClause
      limit?: number
      offset?: number
    } = {}): Promise<Dashboard[]> {
      const raw = await apiCall<unknown>(opts, "Dashboard", "get", params)
      return raw.map((v) => DashboardSchema.parse(v))
    },

    async getById(id: number, select?: (keyof Dashboard)[]): Promise<Dashboard> {
      const results = await apiCall<unknown>(opts, "Dashboard", "get", {
        where: [["id", "=", id]],
        ...(select ? { select } : {}),
      })
      if (!results.length) throw new Error(`Dashboard ${id} not found`)
      return DashboardSchema.parse(results[0])
    },

    async create(values: Partial<Omit<Dashboard, "id">>): Promise<Dashboard> {
      const results = await apiCall<unknown>(opts, "Dashboard", "create", { values })
      if (!results.length) throw new Error("No value returned from Dashboard.create")
      return DashboardSchema.parse(results[0])
    },

    async update(id: number, values: Partial<Omit<Dashboard, "id">>): Promise<Dashboard> {
      const results = await apiCall<unknown>(opts, "Dashboard", "update", {
        where: [["id", "=", id]],
        values,
      })
      if (!results.length) throw new Error("No value returned from Dashboard.update")
      return DashboardSchema.parse(results[0])
    },

    async delete(id: number): Promise<void> {
      await apiCall<unknown>(opts, "Dashboard", "delete", {
        where: [["id", "=", id]],
      })
    },
  }
}

export function createDashboardContactClient(opts: CiviCRMClientOptions) {
  return {
    async get(params: {
      where?: WhereClause
      select?: (keyof DashboardContact)[]
      orderBy?: OrderByClause
      limit?: number
      offset?: number
    } = {}): Promise<DashboardContact[]> {
      const raw = await apiCall<unknown>(opts, "DashboardContact", "get", params)
      return raw.map((v) => DashboardContactSchema.parse(v))
    },

    async getById(id: number, select?: (keyof DashboardContact)[]): Promise<DashboardContact> {
      const results = await apiCall<unknown>(opts, "DashboardContact", "get", {
        where: [["id", "=", id]],
        ...(select ? { select } : {}),
      })
      if (!results.length) throw new Error(`DashboardContact ${id} not found`)
      return DashboardContactSchema.parse(results[0])
    },

    async create(values: Partial<Omit<DashboardContact, "id">>): Promise<DashboardContact> {
      const results = await apiCall<unknown>(opts, "DashboardContact", "create", { values })
      if (!results.length) throw new Error("No value returned from DashboardContact.create")
      return DashboardContactSchema.parse(results[0])
    },

    async update(id: number, values: Partial<Omit<DashboardContact, "id">>): Promise<DashboardContact> {
      const results = await apiCall<unknown>(opts, "DashboardContact", "update", {
        where: [["id", "=", id]],
        values,
      })
      if (!results.length) throw new Error("No value returned from DashboardContact.update")
      return DashboardContactSchema.parse(results[0])
    },

    async delete(id: number): Promise<void> {
      await apiCall<unknown>(opts, "DashboardContact", "delete", {
        where: [["id", "=", id]],
      })
    },
  }
}

export function createDedupeExceptionClient(opts: CiviCRMClientOptions) {
  return {
    async get(params: {
      where?: WhereClause
      select?: (keyof DedupeException)[]
      orderBy?: OrderByClause
      limit?: number
      offset?: number
    } = {}): Promise<DedupeException[]> {
      const raw = await apiCall<unknown>(opts, "DedupeException", "get", params)
      return raw.map((v) => DedupeExceptionSchema.parse(v))
    },

    async getById(id: number, select?: (keyof DedupeException)[]): Promise<DedupeException> {
      const results = await apiCall<unknown>(opts, "DedupeException", "get", {
        where: [["id", "=", id]],
        ...(select ? { select } : {}),
      })
      if (!results.length) throw new Error(`DedupeException ${id} not found`)
      return DedupeExceptionSchema.parse(results[0])
    },

    async create(values: Partial<Omit<DedupeException, "id">>): Promise<DedupeException> {
      const results = await apiCall<unknown>(opts, "DedupeException", "create", { values })
      if (!results.length) throw new Error("No value returned from DedupeException.create")
      return DedupeExceptionSchema.parse(results[0])
    },

    async update(id: number, values: Partial<Omit<DedupeException, "id">>): Promise<DedupeException> {
      const results = await apiCall<unknown>(opts, "DedupeException", "update", {
        where: [["id", "=", id]],
        values,
      })
      if (!results.length) throw new Error("No value returned from DedupeException.update")
      return DedupeExceptionSchema.parse(results[0])
    },

    async delete(id: number): Promise<void> {
      await apiCall<unknown>(opts, "DedupeException", "delete", {
        where: [["id", "=", id]],
      })
    },
  }
}

export function createDedupeRuleClient(opts: CiviCRMClientOptions) {
  return {
    async get(params: {
      where?: WhereClause
      select?: (keyof DedupeRule)[]
      orderBy?: OrderByClause
      limit?: number
      offset?: number
    } = {}): Promise<DedupeRule[]> {
      const raw = await apiCall<unknown>(opts, "DedupeRule", "get", params)
      return raw.map((v) => DedupeRuleSchema.parse(v))
    },

    async getById(id: number, select?: (keyof DedupeRule)[]): Promise<DedupeRule> {
      const results = await apiCall<unknown>(opts, "DedupeRule", "get", {
        where: [["id", "=", id]],
        ...(select ? { select } : {}),
      })
      if (!results.length) throw new Error(`DedupeRule ${id} not found`)
      return DedupeRuleSchema.parse(results[0])
    },

    async create(values: Partial<Omit<DedupeRule, "id">>): Promise<DedupeRule> {
      const results = await apiCall<unknown>(opts, "DedupeRule", "create", { values })
      if (!results.length) throw new Error("No value returned from DedupeRule.create")
      return DedupeRuleSchema.parse(results[0])
    },

    async update(id: number, values: Partial<Omit<DedupeRule, "id">>): Promise<DedupeRule> {
      const results = await apiCall<unknown>(opts, "DedupeRule", "update", {
        where: [["id", "=", id]],
        values,
      })
      if (!results.length) throw new Error("No value returned from DedupeRule.update")
      return DedupeRuleSchema.parse(results[0])
    },

    async delete(id: number): Promise<void> {
      await apiCall<unknown>(opts, "DedupeRule", "delete", {
        where: [["id", "=", id]],
      })
    },
  }
}

export function createDedupeRuleGroupClient(opts: CiviCRMClientOptions) {
  return {
    async get(params: {
      where?: WhereClause
      select?: (keyof DedupeRuleGroup)[]
      orderBy?: OrderByClause
      limit?: number
      offset?: number
    } = {}): Promise<DedupeRuleGroup[]> {
      const raw = await apiCall<unknown>(opts, "DedupeRuleGroup", "get", params)
      return raw.map((v) => DedupeRuleGroupSchema.parse(v))
    },

    async getById(id: number, select?: (keyof DedupeRuleGroup)[]): Promise<DedupeRuleGroup> {
      const results = await apiCall<unknown>(opts, "DedupeRuleGroup", "get", {
        where: [["id", "=", id]],
        ...(select ? { select } : {}),
      })
      if (!results.length) throw new Error(`DedupeRuleGroup ${id} not found`)
      return DedupeRuleGroupSchema.parse(results[0])
    },

    async create(values: Partial<Omit<DedupeRuleGroup, "id">>): Promise<DedupeRuleGroup> {
      const results = await apiCall<unknown>(opts, "DedupeRuleGroup", "create", { values })
      if (!results.length) throw new Error("No value returned from DedupeRuleGroup.create")
      return DedupeRuleGroupSchema.parse(results[0])
    },

    async update(id: number, values: Partial<Omit<DedupeRuleGroup, "id">>): Promise<DedupeRuleGroup> {
      const results = await apiCall<unknown>(opts, "DedupeRuleGroup", "update", {
        where: [["id", "=", id]],
        values,
      })
      if (!results.length) throw new Error("No value returned from DedupeRuleGroup.update")
      return DedupeRuleGroupSchema.parse(results[0])
    },

    async delete(id: number): Promise<void> {
      await apiCall<unknown>(opts, "DedupeRuleGroup", "delete", {
        where: [["id", "=", id]],
      })
    },
  }
}

export function createDiscountClient(opts: CiviCRMClientOptions) {
  return {
    async get(params: {
      where?: WhereClause
      select?: (keyof Discount)[]
      orderBy?: OrderByClause
      limit?: number
      offset?: number
    } = {}): Promise<Discount[]> {
      const raw = await apiCall<unknown>(opts, "Discount", "get", params)
      return raw.map((v) => DiscountSchema.parse(v))
    },

    async getById(id: number, select?: (keyof Discount)[]): Promise<Discount> {
      const results = await apiCall<unknown>(opts, "Discount", "get", {
        where: [["id", "=", id]],
        ...(select ? { select } : {}),
      })
      if (!results.length) throw new Error(`Discount ${id} not found`)
      return DiscountSchema.parse(results[0])
    },

    async create(values: Partial<Omit<Discount, "id">>): Promise<Discount> {
      const results = await apiCall<unknown>(opts, "Discount", "create", { values })
      if (!results.length) throw new Error("No value returned from Discount.create")
      return DiscountSchema.parse(results[0])
    },

    async update(id: number, values: Partial<Omit<Discount, "id">>): Promise<Discount> {
      const results = await apiCall<unknown>(opts, "Discount", "update", {
        where: [["id", "=", id]],
        values,
      })
      if (!results.length) throw new Error("No value returned from Discount.update")
      return DiscountSchema.parse(results[0])
    },

    async delete(id: number): Promise<void> {
      await apiCall<unknown>(opts, "Discount", "delete", {
        where: [["id", "=", id]],
      })
    },
  }
}

export function createDomainClient(opts: CiviCRMClientOptions) {
  return {
    async get(params: {
      where?: WhereClause
      select?: (keyof Domain)[]
      orderBy?: OrderByClause
      limit?: number
      offset?: number
    } = {}): Promise<Domain[]> {
      const raw = await apiCall<unknown>(opts, "Domain", "get", params)
      return raw.map((v) => DomainSchema.parse(v))
    },

    async getById(id: number, select?: (keyof Domain)[]): Promise<Domain> {
      const results = await apiCall<unknown>(opts, "Domain", "get", {
        where: [["id", "=", id]],
        ...(select ? { select } : {}),
      })
      if (!results.length) throw new Error(`Domain ${id} not found`)
      return DomainSchema.parse(results[0])
    },

    async create(values: Partial<Omit<Domain, "id">>): Promise<Domain> {
      const results = await apiCall<unknown>(opts, "Domain", "create", { values })
      if (!results.length) throw new Error("No value returned from Domain.create")
      return DomainSchema.parse(results[0])
    },

    async update(id: number, values: Partial<Omit<Domain, "id">>): Promise<Domain> {
      const results = await apiCall<unknown>(opts, "Domain", "update", {
        where: [["id", "=", id]],
        values,
      })
      if (!results.length) throw new Error("No value returned from Domain.update")
      return DomainSchema.parse(results[0])
    },

    async delete(id: number): Promise<void> {
      await apiCall<unknown>(opts, "Domain", "delete", {
        where: [["id", "=", id]],
      })
    },
  }
}

export function createEmailClient(opts: CiviCRMClientOptions) {
  return {
    async get(params: {
      where?: WhereClause
      select?: (keyof Email)[]
      orderBy?: OrderByClause
      limit?: number
      offset?: number
    } = {}): Promise<Email[]> {
      const raw = await apiCall<unknown>(opts, "Email", "get", params)
      return raw.map((v) => EmailSchema.parse(v))
    },

    async getById(id: number, select?: (keyof Email)[]): Promise<Email> {
      const results = await apiCall<unknown>(opts, "Email", "get", {
        where: [["id", "=", id]],
        ...(select ? { select } : {}),
      })
      if (!results.length) throw new Error(`Email ${id} not found`)
      return EmailSchema.parse(results[0])
    },

    async create(values: Partial<Omit<Email, "id">>): Promise<Email> {
      const results = await apiCall<unknown>(opts, "Email", "create", { values })
      if (!results.length) throw new Error("No value returned from Email.create")
      return EmailSchema.parse(results[0])
    },

    async update(id: number, values: Partial<Omit<Email, "id">>): Promise<Email> {
      const results = await apiCall<unknown>(opts, "Email", "update", {
        where: [["id", "=", id]],
        values,
      })
      if (!results.length) throw new Error("No value returned from Email.update")
      return EmailSchema.parse(results[0])
    },

    async delete(id: number): Promise<void> {
      await apiCall<unknown>(opts, "Email", "delete", {
        where: [["id", "=", id]],
      })
    },
  }
}

export function createEntityClient(opts: CiviCRMClientOptions) {
  return {
    async get(params: {
      where?: WhereClause
      select?: (keyof Entity)[]
      orderBy?: OrderByClause
      limit?: number
      offset?: number
    } = {}): Promise<Entity[]> {
      const raw = await apiCall<unknown>(opts, "Entity", "get", params)
      return raw.map((v) => EntitySchema.parse(v))
    },

    async getById(id: number, select?: (keyof Entity)[]): Promise<Entity> {
      const results = await apiCall<unknown>(opts, "Entity", "get", {
        where: [["id", "=", id]],
        ...(select ? { select } : {}),
      })
      if (!results.length) throw new Error(`Entity ${id} not found`)
      return EntitySchema.parse(results[0])
    },

    async create(values: Partial<Omit<Entity, "id">>): Promise<Entity> {
      const results = await apiCall<unknown>(opts, "Entity", "create", { values })
      if (!results.length) throw new Error("No value returned from Entity.create")
      return EntitySchema.parse(results[0])
    },

    async update(id: number, values: Partial<Omit<Entity, "id">>): Promise<Entity> {
      const results = await apiCall<unknown>(opts, "Entity", "update", {
        where: [["id", "=", id]],
        values,
      })
      if (!results.length) throw new Error("No value returned from Entity.update")
      return EntitySchema.parse(results[0])
    },

    async delete(id: number): Promise<void> {
      await apiCall<unknown>(opts, "Entity", "delete", {
        where: [["id", "=", id]],
      })
    },
  }
}

export function createEntityBatchClient(opts: CiviCRMClientOptions) {
  return {
    async get(params: {
      where?: WhereClause
      select?: (keyof EntityBatch)[]
      orderBy?: OrderByClause
      limit?: number
      offset?: number
    } = {}): Promise<EntityBatch[]> {
      const raw = await apiCall<unknown>(opts, "EntityBatch", "get", params)
      return raw.map((v) => EntityBatchSchema.parse(v))
    },

    async getById(id: number, select?: (keyof EntityBatch)[]): Promise<EntityBatch> {
      const results = await apiCall<unknown>(opts, "EntityBatch", "get", {
        where: [["id", "=", id]],
        ...(select ? { select } : {}),
      })
      if (!results.length) throw new Error(`EntityBatch ${id} not found`)
      return EntityBatchSchema.parse(results[0])
    },

    async create(values: Partial<Omit<EntityBatch, "id">>): Promise<EntityBatch> {
      const results = await apiCall<unknown>(opts, "EntityBatch", "create", { values })
      if (!results.length) throw new Error("No value returned from EntityBatch.create")
      return EntityBatchSchema.parse(results[0])
    },

    async update(id: number, values: Partial<Omit<EntityBatch, "id">>): Promise<EntityBatch> {
      const results = await apiCall<unknown>(opts, "EntityBatch", "update", {
        where: [["id", "=", id]],
        values,
      })
      if (!results.length) throw new Error("No value returned from EntityBatch.update")
      return EntityBatchSchema.parse(results[0])
    },

    async delete(id: number): Promise<void> {
      await apiCall<unknown>(opts, "EntityBatch", "delete", {
        where: [["id", "=", id]],
      })
    },
  }
}

export function createEntityFileClient(opts: CiviCRMClientOptions) {
  return {
    async get(params: {
      where?: WhereClause
      select?: (keyof EntityFile)[]
      orderBy?: OrderByClause
      limit?: number
      offset?: number
    } = {}): Promise<EntityFile[]> {
      const raw = await apiCall<unknown>(opts, "EntityFile", "get", params)
      return raw.map((v) => EntityFileSchema.parse(v))
    },

    async getById(id: number, select?: (keyof EntityFile)[]): Promise<EntityFile> {
      const results = await apiCall<unknown>(opts, "EntityFile", "get", {
        where: [["id", "=", id]],
        ...(select ? { select } : {}),
      })
      if (!results.length) throw new Error(`EntityFile ${id} not found`)
      return EntityFileSchema.parse(results[0])
    },

    async create(values: Partial<Omit<EntityFile, "id">>): Promise<EntityFile> {
      const results = await apiCall<unknown>(opts, "EntityFile", "create", { values })
      if (!results.length) throw new Error("No value returned from EntityFile.create")
      return EntityFileSchema.parse(results[0])
    },

    async update(id: number, values: Partial<Omit<EntityFile, "id">>): Promise<EntityFile> {
      const results = await apiCall<unknown>(opts, "EntityFile", "update", {
        where: [["id", "=", id]],
        values,
      })
      if (!results.length) throw new Error("No value returned from EntityFile.update")
      return EntityFileSchema.parse(results[0])
    },

    async delete(id: number): Promise<void> {
      await apiCall<unknown>(opts, "EntityFile", "delete", {
        where: [["id", "=", id]],
      })
    },
  }
}

export function createEntityFinancialAccountClient(opts: CiviCRMClientOptions) {
  return {
    async get(params: {
      where?: WhereClause
      select?: (keyof EntityFinancialAccount)[]
      orderBy?: OrderByClause
      limit?: number
      offset?: number
    } = {}): Promise<EntityFinancialAccount[]> {
      const raw = await apiCall<unknown>(opts, "EntityFinancialAccount", "get", params)
      return raw.map((v) => EntityFinancialAccountSchema.parse(v))
    },

    async getById(id: number, select?: (keyof EntityFinancialAccount)[]): Promise<EntityFinancialAccount> {
      const results = await apiCall<unknown>(opts, "EntityFinancialAccount", "get", {
        where: [["id", "=", id]],
        ...(select ? { select } : {}),
      })
      if (!results.length) throw new Error(`EntityFinancialAccount ${id} not found`)
      return EntityFinancialAccountSchema.parse(results[0])
    },

    async create(values: Partial<Omit<EntityFinancialAccount, "id">>): Promise<EntityFinancialAccount> {
      const results = await apiCall<unknown>(opts, "EntityFinancialAccount", "create", { values })
      if (!results.length) throw new Error("No value returned from EntityFinancialAccount.create")
      return EntityFinancialAccountSchema.parse(results[0])
    },

    async update(id: number, values: Partial<Omit<EntityFinancialAccount, "id">>): Promise<EntityFinancialAccount> {
      const results = await apiCall<unknown>(opts, "EntityFinancialAccount", "update", {
        where: [["id", "=", id]],
        values,
      })
      if (!results.length) throw new Error("No value returned from EntityFinancialAccount.update")
      return EntityFinancialAccountSchema.parse(results[0])
    },

    async delete(id: number): Promise<void> {
      await apiCall<unknown>(opts, "EntityFinancialAccount", "delete", {
        where: [["id", "=", id]],
      })
    },
  }
}

export function createEntityFinancialTrxnClient(opts: CiviCRMClientOptions) {
  return {
    async get(params: {
      where?: WhereClause
      select?: (keyof EntityFinancialTrxn)[]
      orderBy?: OrderByClause
      limit?: number
      offset?: number
    } = {}): Promise<EntityFinancialTrxn[]> {
      const raw = await apiCall<unknown>(opts, "EntityFinancialTrxn", "get", params)
      return raw.map((v) => EntityFinancialTrxnSchema.parse(v))
    },

    async getById(id: number, select?: (keyof EntityFinancialTrxn)[]): Promise<EntityFinancialTrxn> {
      const results = await apiCall<unknown>(opts, "EntityFinancialTrxn", "get", {
        where: [["id", "=", id]],
        ...(select ? { select } : {}),
      })
      if (!results.length) throw new Error(`EntityFinancialTrxn ${id} not found`)
      return EntityFinancialTrxnSchema.parse(results[0])
    },

    async create(values: Partial<Omit<EntityFinancialTrxn, "id">>): Promise<EntityFinancialTrxn> {
      const results = await apiCall<unknown>(opts, "EntityFinancialTrxn", "create", { values })
      if (!results.length) throw new Error("No value returned from EntityFinancialTrxn.create")
      return EntityFinancialTrxnSchema.parse(results[0])
    },

    async update(id: number, values: Partial<Omit<EntityFinancialTrxn, "id">>): Promise<EntityFinancialTrxn> {
      const results = await apiCall<unknown>(opts, "EntityFinancialTrxn", "update", {
        where: [["id", "=", id]],
        values,
      })
      if (!results.length) throw new Error("No value returned from EntityFinancialTrxn.update")
      return EntityFinancialTrxnSchema.parse(results[0])
    },

    async delete(id: number): Promise<void> {
      await apiCall<unknown>(opts, "EntityFinancialTrxn", "delete", {
        where: [["id", "=", id]],
      })
    },
  }
}

export function createEntityTagClient(opts: CiviCRMClientOptions) {
  return {
    async get(params: {
      where?: WhereClause
      select?: (keyof EntityTag)[]
      orderBy?: OrderByClause
      limit?: number
      offset?: number
    } = {}): Promise<EntityTag[]> {
      const raw = await apiCall<unknown>(opts, "EntityTag", "get", params)
      return raw.map((v) => EntityTagSchema.parse(v))
    },

    async getById(id: number, select?: (keyof EntityTag)[]): Promise<EntityTag> {
      const results = await apiCall<unknown>(opts, "EntityTag", "get", {
        where: [["id", "=", id]],
        ...(select ? { select } : {}),
      })
      if (!results.length) throw new Error(`EntityTag ${id} not found`)
      return EntityTagSchema.parse(results[0])
    },

    async create(values: Partial<Omit<EntityTag, "id">>): Promise<EntityTag> {
      const results = await apiCall<unknown>(opts, "EntityTag", "create", { values })
      if (!results.length) throw new Error("No value returned from EntityTag.create")
      return EntityTagSchema.parse(results[0])
    },

    async update(id: number, values: Partial<Omit<EntityTag, "id">>): Promise<EntityTag> {
      const results = await apiCall<unknown>(opts, "EntityTag", "update", {
        where: [["id", "=", id]],
        values,
      })
      if (!results.length) throw new Error("No value returned from EntityTag.update")
      return EntityTagSchema.parse(results[0])
    },

    async delete(id: number): Promise<void> {
      await apiCall<unknown>(opts, "EntityTag", "delete", {
        where: [["id", "=", id]],
      })
    },
  }
}

export function createEventClient(opts: CiviCRMClientOptions) {
  return {
    async get(params: {
      where?: WhereClause
      select?: (keyof Event)[]
      orderBy?: OrderByClause
      limit?: number
      offset?: number
    } = {}): Promise<Event[]> {
      const raw = await apiCall<unknown>(opts, "Event", "get", params)
      return raw.map((v) => EventSchema.parse(v))
    },

    async getById(id: number, select?: (keyof Event)[]): Promise<Event> {
      const results = await apiCall<unknown>(opts, "Event", "get", {
        where: [["id", "=", id]],
        ...(select ? { select } : {}),
      })
      if (!results.length) throw new Error(`Event ${id} not found`)
      return EventSchema.parse(results[0])
    },

    async create(values: Partial<Omit<Event, "id">>): Promise<Event> {
      const results = await apiCall<unknown>(opts, "Event", "create", { values })
      if (!results.length) throw new Error("No value returned from Event.create")
      return EventSchema.parse(results[0])
    },

    async update(id: number, values: Partial<Omit<Event, "id">>): Promise<Event> {
      const results = await apiCall<unknown>(opts, "Event", "update", {
        where: [["id", "=", id]],
        values,
      })
      if (!results.length) throw new Error("No value returned from Event.update")
      return EventSchema.parse(results[0])
    },

    async delete(id: number): Promise<void> {
      await apiCall<unknown>(opts, "Event", "delete", {
        where: [["id", "=", id]],
      })
    },
  }
}

export function createExampleDataClient(opts: CiviCRMClientOptions) {
  return {
    async get(params: {
      where?: WhereClause
      select?: (keyof ExampleData)[]
      orderBy?: OrderByClause
      limit?: number
      offset?: number
    } = {}): Promise<ExampleData[]> {
      const raw = await apiCall<unknown>(opts, "ExampleData", "get", params)
      return raw.map((v) => ExampleDataSchema.parse(v))
    },

    async getById(id: number, select?: (keyof ExampleData)[]): Promise<ExampleData> {
      const results = await apiCall<unknown>(opts, "ExampleData", "get", {
        where: [["id", "=", id]],
        ...(select ? { select } : {}),
      })
      if (!results.length) throw new Error(`ExampleData ${id} not found`)
      return ExampleDataSchema.parse(results[0])
    },

    async create(values: Partial<Omit<ExampleData, "id">>): Promise<ExampleData> {
      const results = await apiCall<unknown>(opts, "ExampleData", "create", { values })
      if (!results.length) throw new Error("No value returned from ExampleData.create")
      return ExampleDataSchema.parse(results[0])
    },

    async update(id: number, values: Partial<Omit<ExampleData, "id">>): Promise<ExampleData> {
      const results = await apiCall<unknown>(opts, "ExampleData", "update", {
        where: [["id", "=", id]],
        values,
      })
      if (!results.length) throw new Error("No value returned from ExampleData.update")
      return ExampleDataSchema.parse(results[0])
    },

    async delete(id: number): Promise<void> {
      await apiCall<unknown>(opts, "ExampleData", "delete", {
        where: [["id", "=", id]],
      })
    },
  }
}

export function createExtensionClient(opts: CiviCRMClientOptions) {
  return {
    async get(params: {
      where?: WhereClause
      select?: (keyof Extension)[]
      orderBy?: OrderByClause
      limit?: number
      offset?: number
    } = {}): Promise<Extension[]> {
      const raw = await apiCall<unknown>(opts, "Extension", "get", params)
      return raw.map((v) => ExtensionSchema.parse(v))
    },

    async getById(id: number, select?: (keyof Extension)[]): Promise<Extension> {
      const results = await apiCall<unknown>(opts, "Extension", "get", {
        where: [["id", "=", id]],
        ...(select ? { select } : {}),
      })
      if (!results.length) throw new Error(`Extension ${id} not found`)
      return ExtensionSchema.parse(results[0])
    },

    async create(values: Partial<Omit<Extension, "id">>): Promise<Extension> {
      const results = await apiCall<unknown>(opts, "Extension", "create", { values })
      if (!results.length) throw new Error("No value returned from Extension.create")
      return ExtensionSchema.parse(results[0])
    },

    async update(id: number, values: Partial<Omit<Extension, "id">>): Promise<Extension> {
      const results = await apiCall<unknown>(opts, "Extension", "update", {
        where: [["id", "=", id]],
        values,
      })
      if (!results.length) throw new Error("No value returned from Extension.update")
      return ExtensionSchema.parse(results[0])
    },

    async delete(id: number): Promise<void> {
      await apiCall<unknown>(opts, "Extension", "delete", {
        where: [["id", "=", id]],
      })
    },
  }
}

export function createFileClient(opts: CiviCRMClientOptions) {
  return {
    async get(params: {
      where?: WhereClause
      select?: (keyof File)[]
      orderBy?: OrderByClause
      limit?: number
      offset?: number
    } = {}): Promise<File[]> {
      const raw = await apiCall<unknown>(opts, "File", "get", params)
      return raw.map((v) => FileSchema.parse(v))
    },

    async getById(id: number, select?: (keyof File)[]): Promise<File> {
      const results = await apiCall<unknown>(opts, "File", "get", {
        where: [["id", "=", id]],
        ...(select ? { select } : {}),
      })
      if (!results.length) throw new Error(`File ${id} not found`)
      return FileSchema.parse(results[0])
    },

    async create(values: Partial<Omit<File, "id">>): Promise<File> {
      const results = await apiCall<unknown>(opts, "File", "create", { values })
      if (!results.length) throw new Error("No value returned from File.create")
      return FileSchema.parse(results[0])
    },

    async update(id: number, values: Partial<Omit<File, "id">>): Promise<File> {
      const results = await apiCall<unknown>(opts, "File", "update", {
        where: [["id", "=", id]],
        values,
      })
      if (!results.length) throw new Error("No value returned from File.update")
      return FileSchema.parse(results[0])
    },

    async delete(id: number): Promise<void> {
      await apiCall<unknown>(opts, "File", "delete", {
        where: [["id", "=", id]],
      })
    },
  }
}

export function createFinancialAccountClient(opts: CiviCRMClientOptions) {
  return {
    async get(params: {
      where?: WhereClause
      select?: (keyof FinancialAccount)[]
      orderBy?: OrderByClause
      limit?: number
      offset?: number
    } = {}): Promise<FinancialAccount[]> {
      const raw = await apiCall<unknown>(opts, "FinancialAccount", "get", params)
      return raw.map((v) => FinancialAccountSchema.parse(v))
    },

    async getById(id: number, select?: (keyof FinancialAccount)[]): Promise<FinancialAccount> {
      const results = await apiCall<unknown>(opts, "FinancialAccount", "get", {
        where: [["id", "=", id]],
        ...(select ? { select } : {}),
      })
      if (!results.length) throw new Error(`FinancialAccount ${id} not found`)
      return FinancialAccountSchema.parse(results[0])
    },

    async create(values: Partial<Omit<FinancialAccount, "id">>): Promise<FinancialAccount> {
      const results = await apiCall<unknown>(opts, "FinancialAccount", "create", { values })
      if (!results.length) throw new Error("No value returned from FinancialAccount.create")
      return FinancialAccountSchema.parse(results[0])
    },

    async update(id: number, values: Partial<Omit<FinancialAccount, "id">>): Promise<FinancialAccount> {
      const results = await apiCall<unknown>(opts, "FinancialAccount", "update", {
        where: [["id", "=", id]],
        values,
      })
      if (!results.length) throw new Error("No value returned from FinancialAccount.update")
      return FinancialAccountSchema.parse(results[0])
    },

    async delete(id: number): Promise<void> {
      await apiCall<unknown>(opts, "FinancialAccount", "delete", {
        where: [["id", "=", id]],
      })
    },
  }
}

export function createFinancialItemClient(opts: CiviCRMClientOptions) {
  return {
    async get(params: {
      where?: WhereClause
      select?: (keyof FinancialItem)[]
      orderBy?: OrderByClause
      limit?: number
      offset?: number
    } = {}): Promise<FinancialItem[]> {
      const raw = await apiCall<unknown>(opts, "FinancialItem", "get", params)
      return raw.map((v) => FinancialItemSchema.parse(v))
    },

    async getById(id: number, select?: (keyof FinancialItem)[]): Promise<FinancialItem> {
      const results = await apiCall<unknown>(opts, "FinancialItem", "get", {
        where: [["id", "=", id]],
        ...(select ? { select } : {}),
      })
      if (!results.length) throw new Error(`FinancialItem ${id} not found`)
      return FinancialItemSchema.parse(results[0])
    },

    async create(values: Partial<Omit<FinancialItem, "id">>): Promise<FinancialItem> {
      const results = await apiCall<unknown>(opts, "FinancialItem", "create", { values })
      if (!results.length) throw new Error("No value returned from FinancialItem.create")
      return FinancialItemSchema.parse(results[0])
    },

    async update(id: number, values: Partial<Omit<FinancialItem, "id">>): Promise<FinancialItem> {
      const results = await apiCall<unknown>(opts, "FinancialItem", "update", {
        where: [["id", "=", id]],
        values,
      })
      if (!results.length) throw new Error("No value returned from FinancialItem.update")
      return FinancialItemSchema.parse(results[0])
    },

    async delete(id: number): Promise<void> {
      await apiCall<unknown>(opts, "FinancialItem", "delete", {
        where: [["id", "=", id]],
      })
    },
  }
}

export function createFinancialTrxnClient(opts: CiviCRMClientOptions) {
  return {
    async get(params: {
      where?: WhereClause
      select?: (keyof FinancialTrxn)[]
      orderBy?: OrderByClause
      limit?: number
      offset?: number
    } = {}): Promise<FinancialTrxn[]> {
      const raw = await apiCall<unknown>(opts, "FinancialTrxn", "get", params)
      return raw.map((v) => FinancialTrxnSchema.parse(v))
    },

    async getById(id: number, select?: (keyof FinancialTrxn)[]): Promise<FinancialTrxn> {
      const results = await apiCall<unknown>(opts, "FinancialTrxn", "get", {
        where: [["id", "=", id]],
        ...(select ? { select } : {}),
      })
      if (!results.length) throw new Error(`FinancialTrxn ${id} not found`)
      return FinancialTrxnSchema.parse(results[0])
    },

    async create(values: Partial<Omit<FinancialTrxn, "id">>): Promise<FinancialTrxn> {
      const results = await apiCall<unknown>(opts, "FinancialTrxn", "create", { values })
      if (!results.length) throw new Error("No value returned from FinancialTrxn.create")
      return FinancialTrxnSchema.parse(results[0])
    },

    async update(id: number, values: Partial<Omit<FinancialTrxn, "id">>): Promise<FinancialTrxn> {
      const results = await apiCall<unknown>(opts, "FinancialTrxn", "update", {
        where: [["id", "=", id]],
        values,
      })
      if (!results.length) throw new Error("No value returned from FinancialTrxn.update")
      return FinancialTrxnSchema.parse(results[0])
    },

    async delete(id: number): Promise<void> {
      await apiCall<unknown>(opts, "FinancialTrxn", "delete", {
        where: [["id", "=", id]],
      })
    },
  }
}

export function createFinancialTypeClient(opts: CiviCRMClientOptions) {
  return {
    async get(params: {
      where?: WhereClause
      select?: (keyof FinancialType)[]
      orderBy?: OrderByClause
      limit?: number
      offset?: number
    } = {}): Promise<FinancialType[]> {
      const raw = await apiCall<unknown>(opts, "FinancialType", "get", params)
      return raw.map((v) => FinancialTypeSchema.parse(v))
    },

    async getById(id: number, select?: (keyof FinancialType)[]): Promise<FinancialType> {
      const results = await apiCall<unknown>(opts, "FinancialType", "get", {
        where: [["id", "=", id]],
        ...(select ? { select } : {}),
      })
      if (!results.length) throw new Error(`FinancialType ${id} not found`)
      return FinancialTypeSchema.parse(results[0])
    },

    async create(values: Partial<Omit<FinancialType, "id">>): Promise<FinancialType> {
      const results = await apiCall<unknown>(opts, "FinancialType", "create", { values })
      if (!results.length) throw new Error("No value returned from FinancialType.create")
      return FinancialTypeSchema.parse(results[0])
    },

    async update(id: number, values: Partial<Omit<FinancialType, "id">>): Promise<FinancialType> {
      const results = await apiCall<unknown>(opts, "FinancialType", "update", {
        where: [["id", "=", id]],
        values,
      })
      if (!results.length) throw new Error("No value returned from FinancialType.update")
      return FinancialTypeSchema.parse(results[0])
    },

    async delete(id: number): Promise<void> {
      await apiCall<unknown>(opts, "FinancialType", "delete", {
        where: [["id", "=", id]],
      })
    },
  }
}

export function createGroupClient(opts: CiviCRMClientOptions) {
  return {
    async get(params: {
      where?: WhereClause
      select?: (keyof Group)[]
      orderBy?: OrderByClause
      limit?: number
      offset?: number
    } = {}): Promise<Group[]> {
      const raw = await apiCall<unknown>(opts, "Group", "get", params)
      return raw.map((v) => GroupSchema.parse(v))
    },

    async getById(id: number, select?: (keyof Group)[]): Promise<Group> {
      const results = await apiCall<unknown>(opts, "Group", "get", {
        where: [["id", "=", id]],
        ...(select ? { select } : {}),
      })
      if (!results.length) throw new Error(`Group ${id} not found`)
      return GroupSchema.parse(results[0])
    },

    async create(values: Partial<Omit<Group, "id">>): Promise<Group> {
      const results = await apiCall<unknown>(opts, "Group", "create", { values })
      if (!results.length) throw new Error("No value returned from Group.create")
      return GroupSchema.parse(results[0])
    },

    async update(id: number, values: Partial<Omit<Group, "id">>): Promise<Group> {
      const results = await apiCall<unknown>(opts, "Group", "update", {
        where: [["id", "=", id]],
        values,
      })
      if (!results.length) throw new Error("No value returned from Group.update")
      return GroupSchema.parse(results[0])
    },

    async delete(id: number): Promise<void> {
      await apiCall<unknown>(opts, "Group", "delete", {
        where: [["id", "=", id]],
      })
    },
  }
}

export function createGroupContactClient(opts: CiviCRMClientOptions) {
  return {
    async get(params: {
      where?: WhereClause
      select?: (keyof GroupContact)[]
      orderBy?: OrderByClause
      limit?: number
      offset?: number
    } = {}): Promise<GroupContact[]> {
      const raw = await apiCall<unknown>(opts, "GroupContact", "get", params)
      return raw.map((v) => GroupContactSchema.parse(v))
    },

    async getById(id: number, select?: (keyof GroupContact)[]): Promise<GroupContact> {
      const results = await apiCall<unknown>(opts, "GroupContact", "get", {
        where: [["id", "=", id]],
        ...(select ? { select } : {}),
      })
      if (!results.length) throw new Error(`GroupContact ${id} not found`)
      return GroupContactSchema.parse(results[0])
    },

    async create(values: Partial<Omit<GroupContact, "id">>): Promise<GroupContact> {
      const results = await apiCall<unknown>(opts, "GroupContact", "create", { values })
      if (!results.length) throw new Error("No value returned from GroupContact.create")
      return GroupContactSchema.parse(results[0])
    },

    async update(id: number, values: Partial<Omit<GroupContact, "id">>): Promise<GroupContact> {
      const results = await apiCall<unknown>(opts, "GroupContact", "update", {
        where: [["id", "=", id]],
        values,
      })
      if (!results.length) throw new Error("No value returned from GroupContact.update")
      return GroupContactSchema.parse(results[0])
    },

    async delete(id: number): Promise<void> {
      await apiCall<unknown>(opts, "GroupContact", "delete", {
        where: [["id", "=", id]],
      })
    },
  }
}

export function createGroupNestingClient(opts: CiviCRMClientOptions) {
  return {
    async get(params: {
      where?: WhereClause
      select?: (keyof GroupNesting)[]
      orderBy?: OrderByClause
      limit?: number
      offset?: number
    } = {}): Promise<GroupNesting[]> {
      const raw = await apiCall<unknown>(opts, "GroupNesting", "get", params)
      return raw.map((v) => GroupNestingSchema.parse(v))
    },

    async getById(id: number, select?: (keyof GroupNesting)[]): Promise<GroupNesting> {
      const results = await apiCall<unknown>(opts, "GroupNesting", "get", {
        where: [["id", "=", id]],
        ...(select ? { select } : {}),
      })
      if (!results.length) throw new Error(`GroupNesting ${id} not found`)
      return GroupNestingSchema.parse(results[0])
    },

    async create(values: Partial<Omit<GroupNesting, "id">>): Promise<GroupNesting> {
      const results = await apiCall<unknown>(opts, "GroupNesting", "create", { values })
      if (!results.length) throw new Error("No value returned from GroupNesting.create")
      return GroupNestingSchema.parse(results[0])
    },

    async update(id: number, values: Partial<Omit<GroupNesting, "id">>): Promise<GroupNesting> {
      const results = await apiCall<unknown>(opts, "GroupNesting", "update", {
        where: [["id", "=", id]],
        values,
      })
      if (!results.length) throw new Error("No value returned from GroupNesting.update")
      return GroupNestingSchema.parse(results[0])
    },

    async delete(id: number): Promise<void> {
      await apiCall<unknown>(opts, "GroupNesting", "delete", {
        where: [["id", "=", id]],
      })
    },
  }
}

export function createGroupOrganizationClient(opts: CiviCRMClientOptions) {
  return {
    async get(params: {
      where?: WhereClause
      select?: (keyof GroupOrganization)[]
      orderBy?: OrderByClause
      limit?: number
      offset?: number
    } = {}): Promise<GroupOrganization[]> {
      const raw = await apiCall<unknown>(opts, "GroupOrganization", "get", params)
      return raw.map((v) => GroupOrganizationSchema.parse(v))
    },

    async getById(id: number, select?: (keyof GroupOrganization)[]): Promise<GroupOrganization> {
      const results = await apiCall<unknown>(opts, "GroupOrganization", "get", {
        where: [["id", "=", id]],
        ...(select ? { select } : {}),
      })
      if (!results.length) throw new Error(`GroupOrganization ${id} not found`)
      return GroupOrganizationSchema.parse(results[0])
    },

    async create(values: Partial<Omit<GroupOrganization, "id">>): Promise<GroupOrganization> {
      const results = await apiCall<unknown>(opts, "GroupOrganization", "create", { values })
      if (!results.length) throw new Error("No value returned from GroupOrganization.create")
      return GroupOrganizationSchema.parse(results[0])
    },

    async update(id: number, values: Partial<Omit<GroupOrganization, "id">>): Promise<GroupOrganization> {
      const results = await apiCall<unknown>(opts, "GroupOrganization", "update", {
        where: [["id", "=", id]],
        values,
      })
      if (!results.length) throw new Error("No value returned from GroupOrganization.update")
      return GroupOrganizationSchema.parse(results[0])
    },

    async delete(id: number): Promise<void> {
      await apiCall<unknown>(opts, "GroupOrganization", "delete", {
        where: [["id", "=", id]],
      })
    },
  }
}

export function createGroupSubscriptionClient(opts: CiviCRMClientOptions) {
  return {
    async get(params: {
      where?: WhereClause
      select?: (keyof GroupSubscription)[]
      orderBy?: OrderByClause
      limit?: number
      offset?: number
    } = {}): Promise<GroupSubscription[]> {
      const raw = await apiCall<unknown>(opts, "GroupSubscription", "get", params)
      return raw.map((v) => GroupSubscriptionSchema.parse(v))
    },

    async getById(id: number, select?: (keyof GroupSubscription)[]): Promise<GroupSubscription> {
      const results = await apiCall<unknown>(opts, "GroupSubscription", "get", {
        where: [["id", "=", id]],
        ...(select ? { select } : {}),
      })
      if (!results.length) throw new Error(`GroupSubscription ${id} not found`)
      return GroupSubscriptionSchema.parse(results[0])
    },

    async create(values: Partial<Omit<GroupSubscription, "id">>): Promise<GroupSubscription> {
      const results = await apiCall<unknown>(opts, "GroupSubscription", "create", { values })
      if (!results.length) throw new Error("No value returned from GroupSubscription.create")
      return GroupSubscriptionSchema.parse(results[0])
    },

    async update(id: number, values: Partial<Omit<GroupSubscription, "id">>): Promise<GroupSubscription> {
      const results = await apiCall<unknown>(opts, "GroupSubscription", "update", {
        where: [["id", "=", id]],
        values,
      })
      if (!results.length) throw new Error("No value returned from GroupSubscription.update")
      return GroupSubscriptionSchema.parse(results[0])
    },

    async delete(id: number): Promise<void> {
      await apiCall<unknown>(opts, "GroupSubscription", "delete", {
        where: [["id", "=", id]],
      })
    },
  }
}

export function createHouseholdClient(opts: CiviCRMClientOptions) {
  return {
    async get(params: {
      where?: WhereClause
      select?: (keyof Household)[]
      orderBy?: OrderByClause
      limit?: number
      offset?: number
    } = {}): Promise<Household[]> {
      const raw = await apiCall<unknown>(opts, "Household", "get", params)
      return raw.map((v) => HouseholdSchema.parse(v))
    },

    async getById(id: number, select?: (keyof Household)[]): Promise<Household> {
      const results = await apiCall<unknown>(opts, "Household", "get", {
        where: [["id", "=", id]],
        ...(select ? { select } : {}),
      })
      if (!results.length) throw new Error(`Household ${id} not found`)
      return HouseholdSchema.parse(results[0])
    },

    async create(values: Partial<Omit<Household, "id">>): Promise<Household> {
      const results = await apiCall<unknown>(opts, "Household", "create", { values })
      if (!results.length) throw new Error("No value returned from Household.create")
      return HouseholdSchema.parse(results[0])
    },

    async update(id: number, values: Partial<Omit<Household, "id">>): Promise<Household> {
      const results = await apiCall<unknown>(opts, "Household", "update", {
        where: [["id", "=", id]],
        values,
      })
      if (!results.length) throw new Error("No value returned from Household.update")
      return HouseholdSchema.parse(results[0])
    },

    async delete(id: number): Promise<void> {
      await apiCall<unknown>(opts, "Household", "delete", {
        where: [["id", "=", id]],
      })
    },
  }
}

export function createIMClient(opts: CiviCRMClientOptions) {
  return {
    async get(params: {
      where?: WhereClause
      select?: (keyof IM)[]
      orderBy?: OrderByClause
      limit?: number
      offset?: number
    } = {}): Promise<IM[]> {
      const raw = await apiCall<unknown>(opts, "IM", "get", params)
      return raw.map((v) => IMSchema.parse(v))
    },

    async getById(id: number, select?: (keyof IM)[]): Promise<IM> {
      const results = await apiCall<unknown>(opts, "IM", "get", {
        where: [["id", "=", id]],
        ...(select ? { select } : {}),
      })
      if (!results.length) throw new Error(`IM ${id} not found`)
      return IMSchema.parse(results[0])
    },

    async create(values: Partial<Omit<IM, "id">>): Promise<IM> {
      const results = await apiCall<unknown>(opts, "IM", "create", { values })
      if (!results.length) throw new Error("No value returned from IM.create")
      return IMSchema.parse(results[0])
    },

    async update(id: number, values: Partial<Omit<IM, "id">>): Promise<IM> {
      const results = await apiCall<unknown>(opts, "IM", "update", {
        where: [["id", "=", id]],
        values,
      })
      if (!results.length) throw new Error("No value returned from IM.update")
      return IMSchema.parse(results[0])
    },

    async delete(id: number): Promise<void> {
      await apiCall<unknown>(opts, "IM", "delete", {
        where: [["id", "=", id]],
      })
    },
  }
}

export function createIndividualClient(opts: CiviCRMClientOptions) {
  return {
    async get(params: {
      where?: WhereClause
      select?: (keyof Individual)[]
      orderBy?: OrderByClause
      limit?: number
      offset?: number
    } = {}): Promise<Individual[]> {
      const raw = await apiCall<unknown>(opts, "Individual", "get", params)
      return raw.map((v) => IndividualSchema.parse(v))
    },

    async getById(id: number, select?: (keyof Individual)[]): Promise<Individual> {
      const results = await apiCall<unknown>(opts, "Individual", "get", {
        where: [["id", "=", id]],
        ...(select ? { select } : {}),
      })
      if (!results.length) throw new Error(`Individual ${id} not found`)
      return IndividualSchema.parse(results[0])
    },

    async create(values: Partial<Omit<Individual, "id">>): Promise<Individual> {
      const results = await apiCall<unknown>(opts, "Individual", "create", { values })
      if (!results.length) throw new Error("No value returned from Individual.create")
      return IndividualSchema.parse(results[0])
    },

    async update(id: number, values: Partial<Omit<Individual, "id">>): Promise<Individual> {
      const results = await apiCall<unknown>(opts, "Individual", "update", {
        where: [["id", "=", id]],
        values,
      })
      if (!results.length) throw new Error("No value returned from Individual.update")
      return IndividualSchema.parse(results[0])
    },

    async delete(id: number): Promise<void> {
      await apiCall<unknown>(opts, "Individual", "delete", {
        where: [["id", "=", id]],
      })
    },
  }
}

export function createJobClient(opts: CiviCRMClientOptions) {
  return {
    async get(params: {
      where?: WhereClause
      select?: (keyof Job)[]
      orderBy?: OrderByClause
      limit?: number
      offset?: number
    } = {}): Promise<Job[]> {
      const raw = await apiCall<unknown>(opts, "Job", "get", params)
      return raw.map((v) => JobSchema.parse(v))
    },

    async getById(id: number, select?: (keyof Job)[]): Promise<Job> {
      const results = await apiCall<unknown>(opts, "Job", "get", {
        where: [["id", "=", id]],
        ...(select ? { select } : {}),
      })
      if (!results.length) throw new Error(`Job ${id} not found`)
      return JobSchema.parse(results[0])
    },

    async create(values: Partial<Omit<Job, "id">>): Promise<Job> {
      const results = await apiCall<unknown>(opts, "Job", "create", { values })
      if (!results.length) throw new Error("No value returned from Job.create")
      return JobSchema.parse(results[0])
    },

    async update(id: number, values: Partial<Omit<Job, "id">>): Promise<Job> {
      const results = await apiCall<unknown>(opts, "Job", "update", {
        where: [["id", "=", id]],
        values,
      })
      if (!results.length) throw new Error("No value returned from Job.update")
      return JobSchema.parse(results[0])
    },

    async delete(id: number): Promise<void> {
      await apiCall<unknown>(opts, "Job", "delete", {
        where: [["id", "=", id]],
      })
    },
  }
}

export function createJobLogClient(opts: CiviCRMClientOptions) {
  return {
    async get(params: {
      where?: WhereClause
      select?: (keyof JobLog)[]
      orderBy?: OrderByClause
      limit?: number
      offset?: number
    } = {}): Promise<JobLog[]> {
      const raw = await apiCall<unknown>(opts, "JobLog", "get", params)
      return raw.map((v) => JobLogSchema.parse(v))
    },

    async getById(id: number, select?: (keyof JobLog)[]): Promise<JobLog> {
      const results = await apiCall<unknown>(opts, "JobLog", "get", {
        where: [["id", "=", id]],
        ...(select ? { select } : {}),
      })
      if (!results.length) throw new Error(`JobLog ${id} not found`)
      return JobLogSchema.parse(results[0])
    },

    async create(values: Partial<Omit<JobLog, "id">>): Promise<JobLog> {
      const results = await apiCall<unknown>(opts, "JobLog", "create", { values })
      if (!results.length) throw new Error("No value returned from JobLog.create")
      return JobLogSchema.parse(results[0])
    },

    async update(id: number, values: Partial<Omit<JobLog, "id">>): Promise<JobLog> {
      const results = await apiCall<unknown>(opts, "JobLog", "update", {
        where: [["id", "=", id]],
        values,
      })
      if (!results.length) throw new Error("No value returned from JobLog.update")
      return JobLogSchema.parse(results[0])
    },

    async delete(id: number): Promise<void> {
      await apiCall<unknown>(opts, "JobLog", "delete", {
        where: [["id", "=", id]],
      })
    },
  }
}

export function createLineItemClient(opts: CiviCRMClientOptions) {
  return {
    async get(params: {
      where?: WhereClause
      select?: (keyof LineItem)[]
      orderBy?: OrderByClause
      limit?: number
      offset?: number
    } = {}): Promise<LineItem[]> {
      const raw = await apiCall<unknown>(opts, "LineItem", "get", params)
      return raw.map((v) => LineItemSchema.parse(v))
    },

    async getById(id: number, select?: (keyof LineItem)[]): Promise<LineItem> {
      const results = await apiCall<unknown>(opts, "LineItem", "get", {
        where: [["id", "=", id]],
        ...(select ? { select } : {}),
      })
      if (!results.length) throw new Error(`LineItem ${id} not found`)
      return LineItemSchema.parse(results[0])
    },

    async create(values: Partial<Omit<LineItem, "id">>): Promise<LineItem> {
      const results = await apiCall<unknown>(opts, "LineItem", "create", { values })
      if (!results.length) throw new Error("No value returned from LineItem.create")
      return LineItemSchema.parse(results[0])
    },

    async update(id: number, values: Partial<Omit<LineItem, "id">>): Promise<LineItem> {
      const results = await apiCall<unknown>(opts, "LineItem", "update", {
        where: [["id", "=", id]],
        values,
      })
      if (!results.length) throw new Error("No value returned from LineItem.update")
      return LineItemSchema.parse(results[0])
    },

    async delete(id: number): Promise<void> {
      await apiCall<unknown>(opts, "LineItem", "delete", {
        where: [["id", "=", id]],
      })
    },
  }
}

export function createLocBlockClient(opts: CiviCRMClientOptions) {
  return {
    async get(params: {
      where?: WhereClause
      select?: (keyof LocBlock)[]
      orderBy?: OrderByClause
      limit?: number
      offset?: number
    } = {}): Promise<LocBlock[]> {
      const raw = await apiCall<unknown>(opts, "LocBlock", "get", params)
      return raw.map((v) => LocBlockSchema.parse(v))
    },

    async getById(id: number, select?: (keyof LocBlock)[]): Promise<LocBlock> {
      const results = await apiCall<unknown>(opts, "LocBlock", "get", {
        where: [["id", "=", id]],
        ...(select ? { select } : {}),
      })
      if (!results.length) throw new Error(`LocBlock ${id} not found`)
      return LocBlockSchema.parse(results[0])
    },

    async create(values: Partial<Omit<LocBlock, "id">>): Promise<LocBlock> {
      const results = await apiCall<unknown>(opts, "LocBlock", "create", { values })
      if (!results.length) throw new Error("No value returned from LocBlock.create")
      return LocBlockSchema.parse(results[0])
    },

    async update(id: number, values: Partial<Omit<LocBlock, "id">>): Promise<LocBlock> {
      const results = await apiCall<unknown>(opts, "LocBlock", "update", {
        where: [["id", "=", id]],
        values,
      })
      if (!results.length) throw new Error("No value returned from LocBlock.update")
      return LocBlockSchema.parse(results[0])
    },

    async delete(id: number): Promise<void> {
      await apiCall<unknown>(opts, "LocBlock", "delete", {
        where: [["id", "=", id]],
      })
    },
  }
}

export function createLocationTypeClient(opts: CiviCRMClientOptions) {
  return {
    async get(params: {
      where?: WhereClause
      select?: (keyof LocationType)[]
      orderBy?: OrderByClause
      limit?: number
      offset?: number
    } = {}): Promise<LocationType[]> {
      const raw = await apiCall<unknown>(opts, "LocationType", "get", params)
      return raw.map((v) => LocationTypeSchema.parse(v))
    },

    async getById(id: number, select?: (keyof LocationType)[]): Promise<LocationType> {
      const results = await apiCall<unknown>(opts, "LocationType", "get", {
        where: [["id", "=", id]],
        ...(select ? { select } : {}),
      })
      if (!results.length) throw new Error(`LocationType ${id} not found`)
      return LocationTypeSchema.parse(results[0])
    },

    async create(values: Partial<Omit<LocationType, "id">>): Promise<LocationType> {
      const results = await apiCall<unknown>(opts, "LocationType", "create", { values })
      if (!results.length) throw new Error("No value returned from LocationType.create")
      return LocationTypeSchema.parse(results[0])
    },

    async update(id: number, values: Partial<Omit<LocationType, "id">>): Promise<LocationType> {
      const results = await apiCall<unknown>(opts, "LocationType", "update", {
        where: [["id", "=", id]],
        values,
      })
      if (!results.length) throw new Error("No value returned from LocationType.update")
      return LocationTypeSchema.parse(results[0])
    },

    async delete(id: number): Promise<void> {
      await apiCall<unknown>(opts, "LocationType", "delete", {
        where: [["id", "=", id]],
      })
    },
  }
}

export function createLogClient(opts: CiviCRMClientOptions) {
  return {
    async get(params: {
      where?: WhereClause
      select?: (keyof Log)[]
      orderBy?: OrderByClause
      limit?: number
      offset?: number
    } = {}): Promise<Log[]> {
      const raw = await apiCall<unknown>(opts, "Log", "get", params)
      return raw.map((v) => LogSchema.parse(v))
    },

    async getById(id: number, select?: (keyof Log)[]): Promise<Log> {
      const results = await apiCall<unknown>(opts, "Log", "get", {
        where: [["id", "=", id]],
        ...(select ? { select } : {}),
      })
      if (!results.length) throw new Error(`Log ${id} not found`)
      return LogSchema.parse(results[0])
    },

    async create(values: Partial<Omit<Log, "id">>): Promise<Log> {
      const results = await apiCall<unknown>(opts, "Log", "create", { values })
      if (!results.length) throw new Error("No value returned from Log.create")
      return LogSchema.parse(results[0])
    },

    async update(id: number, values: Partial<Omit<Log, "id">>): Promise<Log> {
      const results = await apiCall<unknown>(opts, "Log", "update", {
        where: [["id", "=", id]],
        values,
      })
      if (!results.length) throw new Error("No value returned from Log.update")
      return LogSchema.parse(results[0])
    },

    async delete(id: number): Promise<void> {
      await apiCall<unknown>(opts, "Log", "delete", {
        where: [["id", "=", id]],
      })
    },
  }
}

export function createMailSettingsClient(opts: CiviCRMClientOptions) {
  return {
    async get(params: {
      where?: WhereClause
      select?: (keyof MailSettings)[]
      orderBy?: OrderByClause
      limit?: number
      offset?: number
    } = {}): Promise<MailSettings[]> {
      const raw = await apiCall<unknown>(opts, "MailSettings", "get", params)
      return raw.map((v) => MailSettingsSchema.parse(v))
    },

    async getById(id: number, select?: (keyof MailSettings)[]): Promise<MailSettings> {
      const results = await apiCall<unknown>(opts, "MailSettings", "get", {
        where: [["id", "=", id]],
        ...(select ? { select } : {}),
      })
      if (!results.length) throw new Error(`MailSettings ${id} not found`)
      return MailSettingsSchema.parse(results[0])
    },

    async create(values: Partial<Omit<MailSettings, "id">>): Promise<MailSettings> {
      const results = await apiCall<unknown>(opts, "MailSettings", "create", { values })
      if (!results.length) throw new Error("No value returned from MailSettings.create")
      return MailSettingsSchema.parse(results[0])
    },

    async update(id: number, values: Partial<Omit<MailSettings, "id">>): Promise<MailSettings> {
      const results = await apiCall<unknown>(opts, "MailSettings", "update", {
        where: [["id", "=", id]],
        values,
      })
      if (!results.length) throw new Error("No value returned from MailSettings.update")
      return MailSettingsSchema.parse(results[0])
    },

    async delete(id: number): Promise<void> {
      await apiCall<unknown>(opts, "MailSettings", "delete", {
        where: [["id", "=", id]],
      })
    },
  }
}

export function createMailingClient(opts: CiviCRMClientOptions) {
  return {
    async get(params: {
      where?: WhereClause
      select?: (keyof Mailing)[]
      orderBy?: OrderByClause
      limit?: number
      offset?: number
    } = {}): Promise<Mailing[]> {
      const raw = await apiCall<unknown>(opts, "Mailing", "get", params)
      return raw.map((v) => MailingSchema.parse(v))
    },

    async getById(id: number, select?: (keyof Mailing)[]): Promise<Mailing> {
      const results = await apiCall<unknown>(opts, "Mailing", "get", {
        where: [["id", "=", id]],
        ...(select ? { select } : {}),
      })
      if (!results.length) throw new Error(`Mailing ${id} not found`)
      return MailingSchema.parse(results[0])
    },

    async create(values: Partial<Omit<Mailing, "id">>): Promise<Mailing> {
      const results = await apiCall<unknown>(opts, "Mailing", "create", { values })
      if (!results.length) throw new Error("No value returned from Mailing.create")
      return MailingSchema.parse(results[0])
    },

    async update(id: number, values: Partial<Omit<Mailing, "id">>): Promise<Mailing> {
      const results = await apiCall<unknown>(opts, "Mailing", "update", {
        where: [["id", "=", id]],
        values,
      })
      if (!results.length) throw new Error("No value returned from Mailing.update")
      return MailingSchema.parse(results[0])
    },

    async delete(id: number): Promise<void> {
      await apiCall<unknown>(opts, "Mailing", "delete", {
        where: [["id", "=", id]],
      })
    },
  }
}

export function createMailingComponentClient(opts: CiviCRMClientOptions) {
  return {
    async get(params: {
      where?: WhereClause
      select?: (keyof MailingComponent)[]
      orderBy?: OrderByClause
      limit?: number
      offset?: number
    } = {}): Promise<MailingComponent[]> {
      const raw = await apiCall<unknown>(opts, "MailingComponent", "get", params)
      return raw.map((v) => MailingComponentSchema.parse(v))
    },

    async getById(id: number, select?: (keyof MailingComponent)[]): Promise<MailingComponent> {
      const results = await apiCall<unknown>(opts, "MailingComponent", "get", {
        where: [["id", "=", id]],
        ...(select ? { select } : {}),
      })
      if (!results.length) throw new Error(`MailingComponent ${id} not found`)
      return MailingComponentSchema.parse(results[0])
    },

    async create(values: Partial<Omit<MailingComponent, "id">>): Promise<MailingComponent> {
      const results = await apiCall<unknown>(opts, "MailingComponent", "create", { values })
      if (!results.length) throw new Error("No value returned from MailingComponent.create")
      return MailingComponentSchema.parse(results[0])
    },

    async update(id: number, values: Partial<Omit<MailingComponent, "id">>): Promise<MailingComponent> {
      const results = await apiCall<unknown>(opts, "MailingComponent", "update", {
        where: [["id", "=", id]],
        values,
      })
      if (!results.length) throw new Error("No value returned from MailingComponent.update")
      return MailingComponentSchema.parse(results[0])
    },

    async delete(id: number): Promise<void> {
      await apiCall<unknown>(opts, "MailingComponent", "delete", {
        where: [["id", "=", id]],
      })
    },
  }
}

export function createMailingEventBounceClient(opts: CiviCRMClientOptions) {
  return {
    async get(params: {
      where?: WhereClause
      select?: (keyof MailingEventBounce)[]
      orderBy?: OrderByClause
      limit?: number
      offset?: number
    } = {}): Promise<MailingEventBounce[]> {
      const raw = await apiCall<unknown>(opts, "MailingEventBounce", "get", params)
      return raw.map((v) => MailingEventBounceSchema.parse(v))
    },

    async getById(id: number, select?: (keyof MailingEventBounce)[]): Promise<MailingEventBounce> {
      const results = await apiCall<unknown>(opts, "MailingEventBounce", "get", {
        where: [["id", "=", id]],
        ...(select ? { select } : {}),
      })
      if (!results.length) throw new Error(`MailingEventBounce ${id} not found`)
      return MailingEventBounceSchema.parse(results[0])
    },

    async create(values: Partial<Omit<MailingEventBounce, "id">>): Promise<MailingEventBounce> {
      const results = await apiCall<unknown>(opts, "MailingEventBounce", "create", { values })
      if (!results.length) throw new Error("No value returned from MailingEventBounce.create")
      return MailingEventBounceSchema.parse(results[0])
    },

    async update(id: number, values: Partial<Omit<MailingEventBounce, "id">>): Promise<MailingEventBounce> {
      const results = await apiCall<unknown>(opts, "MailingEventBounce", "update", {
        where: [["id", "=", id]],
        values,
      })
      if (!results.length) throw new Error("No value returned from MailingEventBounce.update")
      return MailingEventBounceSchema.parse(results[0])
    },

    async delete(id: number): Promise<void> {
      await apiCall<unknown>(opts, "MailingEventBounce", "delete", {
        where: [["id", "=", id]],
      })
    },
  }
}

export function createMailingEventConfirmClient(opts: CiviCRMClientOptions) {
  return {
    async get(params: {
      where?: WhereClause
      select?: (keyof MailingEventConfirm)[]
      orderBy?: OrderByClause
      limit?: number
      offset?: number
    } = {}): Promise<MailingEventConfirm[]> {
      const raw = await apiCall<unknown>(opts, "MailingEventConfirm", "get", params)
      return raw.map((v) => MailingEventConfirmSchema.parse(v))
    },

    async getById(id: number, select?: (keyof MailingEventConfirm)[]): Promise<MailingEventConfirm> {
      const results = await apiCall<unknown>(opts, "MailingEventConfirm", "get", {
        where: [["id", "=", id]],
        ...(select ? { select } : {}),
      })
      if (!results.length) throw new Error(`MailingEventConfirm ${id} not found`)
      return MailingEventConfirmSchema.parse(results[0])
    },

    async create(values: Partial<Omit<MailingEventConfirm, "id">>): Promise<MailingEventConfirm> {
      const results = await apiCall<unknown>(opts, "MailingEventConfirm", "create", { values })
      if (!results.length) throw new Error("No value returned from MailingEventConfirm.create")
      return MailingEventConfirmSchema.parse(results[0])
    },

    async update(id: number, values: Partial<Omit<MailingEventConfirm, "id">>): Promise<MailingEventConfirm> {
      const results = await apiCall<unknown>(opts, "MailingEventConfirm", "update", {
        where: [["id", "=", id]],
        values,
      })
      if (!results.length) throw new Error("No value returned from MailingEventConfirm.update")
      return MailingEventConfirmSchema.parse(results[0])
    },

    async delete(id: number): Promise<void> {
      await apiCall<unknown>(opts, "MailingEventConfirm", "delete", {
        where: [["id", "=", id]],
      })
    },
  }
}

export function createMailingEventDeliveredClient(opts: CiviCRMClientOptions) {
  return {
    async get(params: {
      where?: WhereClause
      select?: (keyof MailingEventDelivered)[]
      orderBy?: OrderByClause
      limit?: number
      offset?: number
    } = {}): Promise<MailingEventDelivered[]> {
      const raw = await apiCall<unknown>(opts, "MailingEventDelivered", "get", params)
      return raw.map((v) => MailingEventDeliveredSchema.parse(v))
    },

    async getById(id: number, select?: (keyof MailingEventDelivered)[]): Promise<MailingEventDelivered> {
      const results = await apiCall<unknown>(opts, "MailingEventDelivered", "get", {
        where: [["id", "=", id]],
        ...(select ? { select } : {}),
      })
      if (!results.length) throw new Error(`MailingEventDelivered ${id} not found`)
      return MailingEventDeliveredSchema.parse(results[0])
    },

    async create(values: Partial<Omit<MailingEventDelivered, "id">>): Promise<MailingEventDelivered> {
      const results = await apiCall<unknown>(opts, "MailingEventDelivered", "create", { values })
      if (!results.length) throw new Error("No value returned from MailingEventDelivered.create")
      return MailingEventDeliveredSchema.parse(results[0])
    },

    async update(id: number, values: Partial<Omit<MailingEventDelivered, "id">>): Promise<MailingEventDelivered> {
      const results = await apiCall<unknown>(opts, "MailingEventDelivered", "update", {
        where: [["id", "=", id]],
        values,
      })
      if (!results.length) throw new Error("No value returned from MailingEventDelivered.update")
      return MailingEventDeliveredSchema.parse(results[0])
    },

    async delete(id: number): Promise<void> {
      await apiCall<unknown>(opts, "MailingEventDelivered", "delete", {
        where: [["id", "=", id]],
      })
    },
  }
}

export function createMailingEventOpenedClient(opts: CiviCRMClientOptions) {
  return {
    async get(params: {
      where?: WhereClause
      select?: (keyof MailingEventOpened)[]
      orderBy?: OrderByClause
      limit?: number
      offset?: number
    } = {}): Promise<MailingEventOpened[]> {
      const raw = await apiCall<unknown>(opts, "MailingEventOpened", "get", params)
      return raw.map((v) => MailingEventOpenedSchema.parse(v))
    },

    async getById(id: number, select?: (keyof MailingEventOpened)[]): Promise<MailingEventOpened> {
      const results = await apiCall<unknown>(opts, "MailingEventOpened", "get", {
        where: [["id", "=", id]],
        ...(select ? { select } : {}),
      })
      if (!results.length) throw new Error(`MailingEventOpened ${id} not found`)
      return MailingEventOpenedSchema.parse(results[0])
    },

    async create(values: Partial<Omit<MailingEventOpened, "id">>): Promise<MailingEventOpened> {
      const results = await apiCall<unknown>(opts, "MailingEventOpened", "create", { values })
      if (!results.length) throw new Error("No value returned from MailingEventOpened.create")
      return MailingEventOpenedSchema.parse(results[0])
    },

    async update(id: number, values: Partial<Omit<MailingEventOpened, "id">>): Promise<MailingEventOpened> {
      const results = await apiCall<unknown>(opts, "MailingEventOpened", "update", {
        where: [["id", "=", id]],
        values,
      })
      if (!results.length) throw new Error("No value returned from MailingEventOpened.update")
      return MailingEventOpenedSchema.parse(results[0])
    },

    async delete(id: number): Promise<void> {
      await apiCall<unknown>(opts, "MailingEventOpened", "delete", {
        where: [["id", "=", id]],
      })
    },
  }
}

export function createMailingEventQueueClient(opts: CiviCRMClientOptions) {
  return {
    async get(params: {
      where?: WhereClause
      select?: (keyof MailingEventQueue)[]
      orderBy?: OrderByClause
      limit?: number
      offset?: number
    } = {}): Promise<MailingEventQueue[]> {
      const raw = await apiCall<unknown>(opts, "MailingEventQueue", "get", params)
      return raw.map((v) => MailingEventQueueSchema.parse(v))
    },

    async getById(id: number, select?: (keyof MailingEventQueue)[]): Promise<MailingEventQueue> {
      const results = await apiCall<unknown>(opts, "MailingEventQueue", "get", {
        where: [["id", "=", id]],
        ...(select ? { select } : {}),
      })
      if (!results.length) throw new Error(`MailingEventQueue ${id} not found`)
      return MailingEventQueueSchema.parse(results[0])
    },

    async create(values: Partial<Omit<MailingEventQueue, "id">>): Promise<MailingEventQueue> {
      const results = await apiCall<unknown>(opts, "MailingEventQueue", "create", { values })
      if (!results.length) throw new Error("No value returned from MailingEventQueue.create")
      return MailingEventQueueSchema.parse(results[0])
    },

    async update(id: number, values: Partial<Omit<MailingEventQueue, "id">>): Promise<MailingEventQueue> {
      const results = await apiCall<unknown>(opts, "MailingEventQueue", "update", {
        where: [["id", "=", id]],
        values,
      })
      if (!results.length) throw new Error("No value returned from MailingEventQueue.update")
      return MailingEventQueueSchema.parse(results[0])
    },

    async delete(id: number): Promise<void> {
      await apiCall<unknown>(opts, "MailingEventQueue", "delete", {
        where: [["id", "=", id]],
      })
    },
  }
}

export function createMailingEventReplyClient(opts: CiviCRMClientOptions) {
  return {
    async get(params: {
      where?: WhereClause
      select?: (keyof MailingEventReply)[]
      orderBy?: OrderByClause
      limit?: number
      offset?: number
    } = {}): Promise<MailingEventReply[]> {
      const raw = await apiCall<unknown>(opts, "MailingEventReply", "get", params)
      return raw.map((v) => MailingEventReplySchema.parse(v))
    },

    async getById(id: number, select?: (keyof MailingEventReply)[]): Promise<MailingEventReply> {
      const results = await apiCall<unknown>(opts, "MailingEventReply", "get", {
        where: [["id", "=", id]],
        ...(select ? { select } : {}),
      })
      if (!results.length) throw new Error(`MailingEventReply ${id} not found`)
      return MailingEventReplySchema.parse(results[0])
    },

    async create(values: Partial<Omit<MailingEventReply, "id">>): Promise<MailingEventReply> {
      const results = await apiCall<unknown>(opts, "MailingEventReply", "create", { values })
      if (!results.length) throw new Error("No value returned from MailingEventReply.create")
      return MailingEventReplySchema.parse(results[0])
    },

    async update(id: number, values: Partial<Omit<MailingEventReply, "id">>): Promise<MailingEventReply> {
      const results = await apiCall<unknown>(opts, "MailingEventReply", "update", {
        where: [["id", "=", id]],
        values,
      })
      if (!results.length) throw new Error("No value returned from MailingEventReply.update")
      return MailingEventReplySchema.parse(results[0])
    },

    async delete(id: number): Promise<void> {
      await apiCall<unknown>(opts, "MailingEventReply", "delete", {
        where: [["id", "=", id]],
      })
    },
  }
}

export function createMailingEventSubscribeClient(opts: CiviCRMClientOptions) {
  return {
    async get(params: {
      where?: WhereClause
      select?: (keyof MailingEventSubscribe)[]
      orderBy?: OrderByClause
      limit?: number
      offset?: number
    } = {}): Promise<MailingEventSubscribe[]> {
      const raw = await apiCall<unknown>(opts, "MailingEventSubscribe", "get", params)
      return raw.map((v) => MailingEventSubscribeSchema.parse(v))
    },

    async getById(id: number, select?: (keyof MailingEventSubscribe)[]): Promise<MailingEventSubscribe> {
      const results = await apiCall<unknown>(opts, "MailingEventSubscribe", "get", {
        where: [["id", "=", id]],
        ...(select ? { select } : {}),
      })
      if (!results.length) throw new Error(`MailingEventSubscribe ${id} not found`)
      return MailingEventSubscribeSchema.parse(results[0])
    },

    async create(values: Partial<Omit<MailingEventSubscribe, "id">>): Promise<MailingEventSubscribe> {
      const results = await apiCall<unknown>(opts, "MailingEventSubscribe", "create", { values })
      if (!results.length) throw new Error("No value returned from MailingEventSubscribe.create")
      return MailingEventSubscribeSchema.parse(results[0])
    },

    async update(id: number, values: Partial<Omit<MailingEventSubscribe, "id">>): Promise<MailingEventSubscribe> {
      const results = await apiCall<unknown>(opts, "MailingEventSubscribe", "update", {
        where: [["id", "=", id]],
        values,
      })
      if (!results.length) throw new Error("No value returned from MailingEventSubscribe.update")
      return MailingEventSubscribeSchema.parse(results[0])
    },

    async delete(id: number): Promise<void> {
      await apiCall<unknown>(opts, "MailingEventSubscribe", "delete", {
        where: [["id", "=", id]],
      })
    },
  }
}

export function createMailingEventTrackableURLOpenClient(opts: CiviCRMClientOptions) {
  return {
    async get(params: {
      where?: WhereClause
      select?: (keyof MailingEventTrackableURLOpen)[]
      orderBy?: OrderByClause
      limit?: number
      offset?: number
    } = {}): Promise<MailingEventTrackableURLOpen[]> {
      const raw = await apiCall<unknown>(opts, "MailingEventTrackableURLOpen", "get", params)
      return raw.map((v) => MailingEventTrackableURLOpenSchema.parse(v))
    },

    async getById(id: number, select?: (keyof MailingEventTrackableURLOpen)[]): Promise<MailingEventTrackableURLOpen> {
      const results = await apiCall<unknown>(opts, "MailingEventTrackableURLOpen", "get", {
        where: [["id", "=", id]],
        ...(select ? { select } : {}),
      })
      if (!results.length) throw new Error(`MailingEventTrackableURLOpen ${id} not found`)
      return MailingEventTrackableURLOpenSchema.parse(results[0])
    },

    async create(values: Partial<Omit<MailingEventTrackableURLOpen, "id">>): Promise<MailingEventTrackableURLOpen> {
      const results = await apiCall<unknown>(opts, "MailingEventTrackableURLOpen", "create", { values })
      if (!results.length) throw new Error("No value returned from MailingEventTrackableURLOpen.create")
      return MailingEventTrackableURLOpenSchema.parse(results[0])
    },

    async update(id: number, values: Partial<Omit<MailingEventTrackableURLOpen, "id">>): Promise<MailingEventTrackableURLOpen> {
      const results = await apiCall<unknown>(opts, "MailingEventTrackableURLOpen", "update", {
        where: [["id", "=", id]],
        values,
      })
      if (!results.length) throw new Error("No value returned from MailingEventTrackableURLOpen.update")
      return MailingEventTrackableURLOpenSchema.parse(results[0])
    },

    async delete(id: number): Promise<void> {
      await apiCall<unknown>(opts, "MailingEventTrackableURLOpen", "delete", {
        where: [["id", "=", id]],
      })
    },
  }
}

export function createMailingEventUnsubscribeClient(opts: CiviCRMClientOptions) {
  return {
    async get(params: {
      where?: WhereClause
      select?: (keyof MailingEventUnsubscribe)[]
      orderBy?: OrderByClause
      limit?: number
      offset?: number
    } = {}): Promise<MailingEventUnsubscribe[]> {
      const raw = await apiCall<unknown>(opts, "MailingEventUnsubscribe", "get", params)
      return raw.map((v) => MailingEventUnsubscribeSchema.parse(v))
    },

    async getById(id: number, select?: (keyof MailingEventUnsubscribe)[]): Promise<MailingEventUnsubscribe> {
      const results = await apiCall<unknown>(opts, "MailingEventUnsubscribe", "get", {
        where: [["id", "=", id]],
        ...(select ? { select } : {}),
      })
      if (!results.length) throw new Error(`MailingEventUnsubscribe ${id} not found`)
      return MailingEventUnsubscribeSchema.parse(results[0])
    },

    async create(values: Partial<Omit<MailingEventUnsubscribe, "id">>): Promise<MailingEventUnsubscribe> {
      const results = await apiCall<unknown>(opts, "MailingEventUnsubscribe", "create", { values })
      if (!results.length) throw new Error("No value returned from MailingEventUnsubscribe.create")
      return MailingEventUnsubscribeSchema.parse(results[0])
    },

    async update(id: number, values: Partial<Omit<MailingEventUnsubscribe, "id">>): Promise<MailingEventUnsubscribe> {
      const results = await apiCall<unknown>(opts, "MailingEventUnsubscribe", "update", {
        where: [["id", "=", id]],
        values,
      })
      if (!results.length) throw new Error("No value returned from MailingEventUnsubscribe.update")
      return MailingEventUnsubscribeSchema.parse(results[0])
    },

    async delete(id: number): Promise<void> {
      await apiCall<unknown>(opts, "MailingEventUnsubscribe", "delete", {
        where: [["id", "=", id]],
      })
    },
  }
}

export function createMailingGroupClient(opts: CiviCRMClientOptions) {
  return {
    async get(params: {
      where?: WhereClause
      select?: (keyof MailingGroup)[]
      orderBy?: OrderByClause
      limit?: number
      offset?: number
    } = {}): Promise<MailingGroup[]> {
      const raw = await apiCall<unknown>(opts, "MailingGroup", "get", params)
      return raw.map((v) => MailingGroupSchema.parse(v))
    },

    async getById(id: number, select?: (keyof MailingGroup)[]): Promise<MailingGroup> {
      const results = await apiCall<unknown>(opts, "MailingGroup", "get", {
        where: [["id", "=", id]],
        ...(select ? { select } : {}),
      })
      if (!results.length) throw new Error(`MailingGroup ${id} not found`)
      return MailingGroupSchema.parse(results[0])
    },

    async create(values: Partial<Omit<MailingGroup, "id">>): Promise<MailingGroup> {
      const results = await apiCall<unknown>(opts, "MailingGroup", "create", { values })
      if (!results.length) throw new Error("No value returned from MailingGroup.create")
      return MailingGroupSchema.parse(results[0])
    },

    async update(id: number, values: Partial<Omit<MailingGroup, "id">>): Promise<MailingGroup> {
      const results = await apiCall<unknown>(opts, "MailingGroup", "update", {
        where: [["id", "=", id]],
        values,
      })
      if (!results.length) throw new Error("No value returned from MailingGroup.update")
      return MailingGroupSchema.parse(results[0])
    },

    async delete(id: number): Promise<void> {
      await apiCall<unknown>(opts, "MailingGroup", "delete", {
        where: [["id", "=", id]],
      })
    },
  }
}

export function createMailingJobClient(opts: CiviCRMClientOptions) {
  return {
    async get(params: {
      where?: WhereClause
      select?: (keyof MailingJob)[]
      orderBy?: OrderByClause
      limit?: number
      offset?: number
    } = {}): Promise<MailingJob[]> {
      const raw = await apiCall<unknown>(opts, "MailingJob", "get", params)
      return raw.map((v) => MailingJobSchema.parse(v))
    },

    async getById(id: number, select?: (keyof MailingJob)[]): Promise<MailingJob> {
      const results = await apiCall<unknown>(opts, "MailingJob", "get", {
        where: [["id", "=", id]],
        ...(select ? { select } : {}),
      })
      if (!results.length) throw new Error(`MailingJob ${id} not found`)
      return MailingJobSchema.parse(results[0])
    },

    async create(values: Partial<Omit<MailingJob, "id">>): Promise<MailingJob> {
      const results = await apiCall<unknown>(opts, "MailingJob", "create", { values })
      if (!results.length) throw new Error("No value returned from MailingJob.create")
      return MailingJobSchema.parse(results[0])
    },

    async update(id: number, values: Partial<Omit<MailingJob, "id">>): Promise<MailingJob> {
      const results = await apiCall<unknown>(opts, "MailingJob", "update", {
        where: [["id", "=", id]],
        values,
      })
      if (!results.length) throw new Error("No value returned from MailingJob.update")
      return MailingJobSchema.parse(results[0])
    },

    async delete(id: number): Promise<void> {
      await apiCall<unknown>(opts, "MailingJob", "delete", {
        where: [["id", "=", id]],
      })
    },
  }
}

export function createMailingTrackableURLClient(opts: CiviCRMClientOptions) {
  return {
    async get(params: {
      where?: WhereClause
      select?: (keyof MailingTrackableURL)[]
      orderBy?: OrderByClause
      limit?: number
      offset?: number
    } = {}): Promise<MailingTrackableURL[]> {
      const raw = await apiCall<unknown>(opts, "MailingTrackableURL", "get", params)
      return raw.map((v) => MailingTrackableURLSchema.parse(v))
    },

    async getById(id: number, select?: (keyof MailingTrackableURL)[]): Promise<MailingTrackableURL> {
      const results = await apiCall<unknown>(opts, "MailingTrackableURL", "get", {
        where: [["id", "=", id]],
        ...(select ? { select } : {}),
      })
      if (!results.length) throw new Error(`MailingTrackableURL ${id} not found`)
      return MailingTrackableURLSchema.parse(results[0])
    },

    async create(values: Partial<Omit<MailingTrackableURL, "id">>): Promise<MailingTrackableURL> {
      const results = await apiCall<unknown>(opts, "MailingTrackableURL", "create", { values })
      if (!results.length) throw new Error("No value returned from MailingTrackableURL.create")
      return MailingTrackableURLSchema.parse(results[0])
    },

    async update(id: number, values: Partial<Omit<MailingTrackableURL, "id">>): Promise<MailingTrackableURL> {
      const results = await apiCall<unknown>(opts, "MailingTrackableURL", "update", {
        where: [["id", "=", id]],
        values,
      })
      if (!results.length) throw new Error("No value returned from MailingTrackableURL.update")
      return MailingTrackableURLSchema.parse(results[0])
    },

    async delete(id: number): Promise<void> {
      await apiCall<unknown>(opts, "MailingTrackableURL", "delete", {
        where: [["id", "=", id]],
      })
    },
  }
}

export function createManagedClient(opts: CiviCRMClientOptions) {
  return {
    async get(params: {
      where?: WhereClause
      select?: (keyof Managed)[]
      orderBy?: OrderByClause
      limit?: number
      offset?: number
    } = {}): Promise<Managed[]> {
      const raw = await apiCall<unknown>(opts, "Managed", "get", params)
      return raw.map((v) => ManagedSchema.parse(v))
    },

    async getById(id: number, select?: (keyof Managed)[]): Promise<Managed> {
      const results = await apiCall<unknown>(opts, "Managed", "get", {
        where: [["id", "=", id]],
        ...(select ? { select } : {}),
      })
      if (!results.length) throw new Error(`Managed ${id} not found`)
      return ManagedSchema.parse(results[0])
    },

    async create(values: Partial<Omit<Managed, "id">>): Promise<Managed> {
      const results = await apiCall<unknown>(opts, "Managed", "create", { values })
      if (!results.length) throw new Error("No value returned from Managed.create")
      return ManagedSchema.parse(results[0])
    },

    async update(id: number, values: Partial<Omit<Managed, "id">>): Promise<Managed> {
      const results = await apiCall<unknown>(opts, "Managed", "update", {
        where: [["id", "=", id]],
        values,
      })
      if (!results.length) throw new Error("No value returned from Managed.update")
      return ManagedSchema.parse(results[0])
    },

    async delete(id: number): Promise<void> {
      await apiCall<unknown>(opts, "Managed", "delete", {
        where: [["id", "=", id]],
      })
    },
  }
}

export function createMappingClient(opts: CiviCRMClientOptions) {
  return {
    async get(params: {
      where?: WhereClause
      select?: (keyof Mapping)[]
      orderBy?: OrderByClause
      limit?: number
      offset?: number
    } = {}): Promise<Mapping[]> {
      const raw = await apiCall<unknown>(opts, "Mapping", "get", params)
      return raw.map((v) => MappingSchema.parse(v))
    },

    async getById(id: number, select?: (keyof Mapping)[]): Promise<Mapping> {
      const results = await apiCall<unknown>(opts, "Mapping", "get", {
        where: [["id", "=", id]],
        ...(select ? { select } : {}),
      })
      if (!results.length) throw new Error(`Mapping ${id} not found`)
      return MappingSchema.parse(results[0])
    },

    async create(values: Partial<Omit<Mapping, "id">>): Promise<Mapping> {
      const results = await apiCall<unknown>(opts, "Mapping", "create", { values })
      if (!results.length) throw new Error("No value returned from Mapping.create")
      return MappingSchema.parse(results[0])
    },

    async update(id: number, values: Partial<Omit<Mapping, "id">>): Promise<Mapping> {
      const results = await apiCall<unknown>(opts, "Mapping", "update", {
        where: [["id", "=", id]],
        values,
      })
      if (!results.length) throw new Error("No value returned from Mapping.update")
      return MappingSchema.parse(results[0])
    },

    async delete(id: number): Promise<void> {
      await apiCall<unknown>(opts, "Mapping", "delete", {
        where: [["id", "=", id]],
      })
    },
  }
}

export function createMappingFieldClient(opts: CiviCRMClientOptions) {
  return {
    async get(params: {
      where?: WhereClause
      select?: (keyof MappingField)[]
      orderBy?: OrderByClause
      limit?: number
      offset?: number
    } = {}): Promise<MappingField[]> {
      const raw = await apiCall<unknown>(opts, "MappingField", "get", params)
      return raw.map((v) => MappingFieldSchema.parse(v))
    },

    async getById(id: number, select?: (keyof MappingField)[]): Promise<MappingField> {
      const results = await apiCall<unknown>(opts, "MappingField", "get", {
        where: [["id", "=", id]],
        ...(select ? { select } : {}),
      })
      if (!results.length) throw new Error(`MappingField ${id} not found`)
      return MappingFieldSchema.parse(results[0])
    },

    async create(values: Partial<Omit<MappingField, "id">>): Promise<MappingField> {
      const results = await apiCall<unknown>(opts, "MappingField", "create", { values })
      if (!results.length) throw new Error("No value returned from MappingField.create")
      return MappingFieldSchema.parse(results[0])
    },

    async update(id: number, values: Partial<Omit<MappingField, "id">>): Promise<MappingField> {
      const results = await apiCall<unknown>(opts, "MappingField", "update", {
        where: [["id", "=", id]],
        values,
      })
      if (!results.length) throw new Error("No value returned from MappingField.update")
      return MappingFieldSchema.parse(results[0])
    },

    async delete(id: number): Promise<void> {
      await apiCall<unknown>(opts, "MappingField", "delete", {
        where: [["id", "=", id]],
      })
    },
  }
}

export function createMembershipClient(opts: CiviCRMClientOptions) {
  return {
    async get(params: {
      where?: WhereClause
      select?: (keyof Membership)[]
      orderBy?: OrderByClause
      limit?: number
      offset?: number
    } = {}): Promise<Membership[]> {
      const raw = await apiCall<unknown>(opts, "Membership", "get", params)
      return raw.map((v) => MembershipSchema.parse(v))
    },

    async getById(id: number, select?: (keyof Membership)[]): Promise<Membership> {
      const results = await apiCall<unknown>(opts, "Membership", "get", {
        where: [["id", "=", id]],
        ...(select ? { select } : {}),
      })
      if (!results.length) throw new Error(`Membership ${id} not found`)
      return MembershipSchema.parse(results[0])
    },

    async create(values: Partial<Omit<Membership, "id">>): Promise<Membership> {
      const results = await apiCall<unknown>(opts, "Membership", "create", { values })
      if (!results.length) throw new Error("No value returned from Membership.create")
      return MembershipSchema.parse(results[0])
    },

    async update(id: number, values: Partial<Omit<Membership, "id">>): Promise<Membership> {
      const results = await apiCall<unknown>(opts, "Membership", "update", {
        where: [["id", "=", id]],
        values,
      })
      if (!results.length) throw new Error("No value returned from Membership.update")
      return MembershipSchema.parse(results[0])
    },

    async delete(id: number): Promise<void> {
      await apiCall<unknown>(opts, "Membership", "delete", {
        where: [["id", "=", id]],
      })
    },
  }
}

export function createMembershipBlockClient(opts: CiviCRMClientOptions) {
  return {
    async get(params: {
      where?: WhereClause
      select?: (keyof MembershipBlock)[]
      orderBy?: OrderByClause
      limit?: number
      offset?: number
    } = {}): Promise<MembershipBlock[]> {
      const raw = await apiCall<unknown>(opts, "MembershipBlock", "get", params)
      return raw.map((v) => MembershipBlockSchema.parse(v))
    },

    async getById(id: number, select?: (keyof MembershipBlock)[]): Promise<MembershipBlock> {
      const results = await apiCall<unknown>(opts, "MembershipBlock", "get", {
        where: [["id", "=", id]],
        ...(select ? { select } : {}),
      })
      if (!results.length) throw new Error(`MembershipBlock ${id} not found`)
      return MembershipBlockSchema.parse(results[0])
    },

    async create(values: Partial<Omit<MembershipBlock, "id">>): Promise<MembershipBlock> {
      const results = await apiCall<unknown>(opts, "MembershipBlock", "create", { values })
      if (!results.length) throw new Error("No value returned from MembershipBlock.create")
      return MembershipBlockSchema.parse(results[0])
    },

    async update(id: number, values: Partial<Omit<MembershipBlock, "id">>): Promise<MembershipBlock> {
      const results = await apiCall<unknown>(opts, "MembershipBlock", "update", {
        where: [["id", "=", id]],
        values,
      })
      if (!results.length) throw new Error("No value returned from MembershipBlock.update")
      return MembershipBlockSchema.parse(results[0])
    },

    async delete(id: number): Promise<void> {
      await apiCall<unknown>(opts, "MembershipBlock", "delete", {
        where: [["id", "=", id]],
      })
    },
  }
}

export function createMembershipLogClient(opts: CiviCRMClientOptions) {
  return {
    async get(params: {
      where?: WhereClause
      select?: (keyof MembershipLog)[]
      orderBy?: OrderByClause
      limit?: number
      offset?: number
    } = {}): Promise<MembershipLog[]> {
      const raw = await apiCall<unknown>(opts, "MembershipLog", "get", params)
      return raw.map((v) => MembershipLogSchema.parse(v))
    },

    async getById(id: number, select?: (keyof MembershipLog)[]): Promise<MembershipLog> {
      const results = await apiCall<unknown>(opts, "MembershipLog", "get", {
        where: [["id", "=", id]],
        ...(select ? { select } : {}),
      })
      if (!results.length) throw new Error(`MembershipLog ${id} not found`)
      return MembershipLogSchema.parse(results[0])
    },

    async create(values: Partial<Omit<MembershipLog, "id">>): Promise<MembershipLog> {
      const results = await apiCall<unknown>(opts, "MembershipLog", "create", { values })
      if (!results.length) throw new Error("No value returned from MembershipLog.create")
      return MembershipLogSchema.parse(results[0])
    },

    async update(id: number, values: Partial<Omit<MembershipLog, "id">>): Promise<MembershipLog> {
      const results = await apiCall<unknown>(opts, "MembershipLog", "update", {
        where: [["id", "=", id]],
        values,
      })
      if (!results.length) throw new Error("No value returned from MembershipLog.update")
      return MembershipLogSchema.parse(results[0])
    },

    async delete(id: number): Promise<void> {
      await apiCall<unknown>(opts, "MembershipLog", "delete", {
        where: [["id", "=", id]],
      })
    },
  }
}

export function createMembershipStatusClient(opts: CiviCRMClientOptions) {
  return {
    async get(params: {
      where?: WhereClause
      select?: (keyof MembershipStatus)[]
      orderBy?: OrderByClause
      limit?: number
      offset?: number
    } = {}): Promise<MembershipStatus[]> {
      const raw = await apiCall<unknown>(opts, "MembershipStatus", "get", params)
      return raw.map((v) => MembershipStatusSchema.parse(v))
    },

    async getById(id: number, select?: (keyof MembershipStatus)[]): Promise<MembershipStatus> {
      const results = await apiCall<unknown>(opts, "MembershipStatus", "get", {
        where: [["id", "=", id]],
        ...(select ? { select } : {}),
      })
      if (!results.length) throw new Error(`MembershipStatus ${id} not found`)
      return MembershipStatusSchema.parse(results[0])
    },

    async create(values: Partial<Omit<MembershipStatus, "id">>): Promise<MembershipStatus> {
      const results = await apiCall<unknown>(opts, "MembershipStatus", "create", { values })
      if (!results.length) throw new Error("No value returned from MembershipStatus.create")
      return MembershipStatusSchema.parse(results[0])
    },

    async update(id: number, values: Partial<Omit<MembershipStatus, "id">>): Promise<MembershipStatus> {
      const results = await apiCall<unknown>(opts, "MembershipStatus", "update", {
        where: [["id", "=", id]],
        values,
      })
      if (!results.length) throw new Error("No value returned from MembershipStatus.update")
      return MembershipStatusSchema.parse(results[0])
    },

    async delete(id: number): Promise<void> {
      await apiCall<unknown>(opts, "MembershipStatus", "delete", {
        where: [["id", "=", id]],
      })
    },
  }
}

export function createMembershipTypeClient(opts: CiviCRMClientOptions) {
  return {
    async get(params: {
      where?: WhereClause
      select?: (keyof MembershipType)[]
      orderBy?: OrderByClause
      limit?: number
      offset?: number
    } = {}): Promise<MembershipType[]> {
      const raw = await apiCall<unknown>(opts, "MembershipType", "get", params)
      return raw.map((v) => MembershipTypeSchema.parse(v))
    },

    async getById(id: number, select?: (keyof MembershipType)[]): Promise<MembershipType> {
      const results = await apiCall<unknown>(opts, "MembershipType", "get", {
        where: [["id", "=", id]],
        ...(select ? { select } : {}),
      })
      if (!results.length) throw new Error(`MembershipType ${id} not found`)
      return MembershipTypeSchema.parse(results[0])
    },

    async create(values: Partial<Omit<MembershipType, "id">>): Promise<MembershipType> {
      const results = await apiCall<unknown>(opts, "MembershipType", "create", { values })
      if (!results.length) throw new Error("No value returned from MembershipType.create")
      return MembershipTypeSchema.parse(results[0])
    },

    async update(id: number, values: Partial<Omit<MembershipType, "id">>): Promise<MembershipType> {
      const results = await apiCall<unknown>(opts, "MembershipType", "update", {
        where: [["id", "=", id]],
        values,
      })
      if (!results.length) throw new Error("No value returned from MembershipType.update")
      return MembershipTypeSchema.parse(results[0])
    },

    async delete(id: number): Promise<void> {
      await apiCall<unknown>(opts, "MembershipType", "delete", {
        where: [["id", "=", id]],
      })
    },
  }
}

export function createMessageTemplateClient(opts: CiviCRMClientOptions) {
  return {
    async get(params: {
      where?: WhereClause
      select?: (keyof MessageTemplate)[]
      orderBy?: OrderByClause
      limit?: number
      offset?: number
    } = {}): Promise<MessageTemplate[]> {
      const raw = await apiCall<unknown>(opts, "MessageTemplate", "get", params)
      return raw.map((v) => MessageTemplateSchema.parse(v))
    },

    async getById(id: number, select?: (keyof MessageTemplate)[]): Promise<MessageTemplate> {
      const results = await apiCall<unknown>(opts, "MessageTemplate", "get", {
        where: [["id", "=", id]],
        ...(select ? { select } : {}),
      })
      if (!results.length) throw new Error(`MessageTemplate ${id} not found`)
      return MessageTemplateSchema.parse(results[0])
    },

    async create(values: Partial<Omit<MessageTemplate, "id">>): Promise<MessageTemplate> {
      const results = await apiCall<unknown>(opts, "MessageTemplate", "create", { values })
      if (!results.length) throw new Error("No value returned from MessageTemplate.create")
      return MessageTemplateSchema.parse(results[0])
    },

    async update(id: number, values: Partial<Omit<MessageTemplate, "id">>): Promise<MessageTemplate> {
      const results = await apiCall<unknown>(opts, "MessageTemplate", "update", {
        where: [["id", "=", id]],
        values,
      })
      if (!results.length) throw new Error("No value returned from MessageTemplate.update")
      return MessageTemplateSchema.parse(results[0])
    },

    async delete(id: number): Promise<void> {
      await apiCall<unknown>(opts, "MessageTemplate", "delete", {
        where: [["id", "=", id]],
      })
    },
  }
}

export function createMosaicoTemplateClient(opts: CiviCRMClientOptions) {
  return {
    async get(params: {
      where?: WhereClause
      select?: (keyof MosaicoTemplate)[]
      orderBy?: OrderByClause
      limit?: number
      offset?: number
    } = {}): Promise<MosaicoTemplate[]> {
      const raw = await apiCall<unknown>(opts, "MosaicoTemplate", "get", params)
      return raw.map((v) => MosaicoTemplateSchema.parse(v))
    },

    async getById(id: number, select?: (keyof MosaicoTemplate)[]): Promise<MosaicoTemplate> {
      const results = await apiCall<unknown>(opts, "MosaicoTemplate", "get", {
        where: [["id", "=", id]],
        ...(select ? { select } : {}),
      })
      if (!results.length) throw new Error(`MosaicoTemplate ${id} not found`)
      return MosaicoTemplateSchema.parse(results[0])
    },

    async create(values: Partial<Omit<MosaicoTemplate, "id">>): Promise<MosaicoTemplate> {
      const results = await apiCall<unknown>(opts, "MosaicoTemplate", "create", { values })
      if (!results.length) throw new Error("No value returned from MosaicoTemplate.create")
      return MosaicoTemplateSchema.parse(results[0])
    },

    async update(id: number, values: Partial<Omit<MosaicoTemplate, "id">>): Promise<MosaicoTemplate> {
      const results = await apiCall<unknown>(opts, "MosaicoTemplate", "update", {
        where: [["id", "=", id]],
        values,
      })
      if (!results.length) throw new Error("No value returned from MosaicoTemplate.update")
      return MosaicoTemplateSchema.parse(results[0])
    },

    async delete(id: number): Promise<void> {
      await apiCall<unknown>(opts, "MosaicoTemplate", "delete", {
        where: [["id", "=", id]],
      })
    },
  }
}

export function createNavigationClient(opts: CiviCRMClientOptions) {
  return {
    async get(params: {
      where?: WhereClause
      select?: (keyof Navigation)[]
      orderBy?: OrderByClause
      limit?: number
      offset?: number
    } = {}): Promise<Navigation[]> {
      const raw = await apiCall<unknown>(opts, "Navigation", "get", params)
      return raw.map((v) => NavigationSchema.parse(v))
    },

    async getById(id: number, select?: (keyof Navigation)[]): Promise<Navigation> {
      const results = await apiCall<unknown>(opts, "Navigation", "get", {
        where: [["id", "=", id]],
        ...(select ? { select } : {}),
      })
      if (!results.length) throw new Error(`Navigation ${id} not found`)
      return NavigationSchema.parse(results[0])
    },

    async create(values: Partial<Omit<Navigation, "id">>): Promise<Navigation> {
      const results = await apiCall<unknown>(opts, "Navigation", "create", { values })
      if (!results.length) throw new Error("No value returned from Navigation.create")
      return NavigationSchema.parse(results[0])
    },

    async update(id: number, values: Partial<Omit<Navigation, "id">>): Promise<Navigation> {
      const results = await apiCall<unknown>(opts, "Navigation", "update", {
        where: [["id", "=", id]],
        values,
      })
      if (!results.length) throw new Error("No value returned from Navigation.update")
      return NavigationSchema.parse(results[0])
    },

    async delete(id: number): Promise<void> {
      await apiCall<unknown>(opts, "Navigation", "delete", {
        where: [["id", "=", id]],
      })
    },
  }
}

export function createNoteClient(opts: CiviCRMClientOptions) {
  return {
    async get(params: {
      where?: WhereClause
      select?: (keyof Note)[]
      orderBy?: OrderByClause
      limit?: number
      offset?: number
    } = {}): Promise<Note[]> {
      const raw = await apiCall<unknown>(opts, "Note", "get", params)
      return raw.map((v) => NoteSchema.parse(v))
    },

    async getById(id: number, select?: (keyof Note)[]): Promise<Note> {
      const results = await apiCall<unknown>(opts, "Note", "get", {
        where: [["id", "=", id]],
        ...(select ? { select } : {}),
      })
      if (!results.length) throw new Error(`Note ${id} not found`)
      return NoteSchema.parse(results[0])
    },

    async create(values: Partial<Omit<Note, "id">>): Promise<Note> {
      const results = await apiCall<unknown>(opts, "Note", "create", { values })
      if (!results.length) throw new Error("No value returned from Note.create")
      return NoteSchema.parse(results[0])
    },

    async update(id: number, values: Partial<Omit<Note, "id">>): Promise<Note> {
      const results = await apiCall<unknown>(opts, "Note", "update", {
        where: [["id", "=", id]],
        values,
      })
      if (!results.length) throw new Error("No value returned from Note.update")
      return NoteSchema.parse(results[0])
    },

    async delete(id: number): Promise<void> {
      await apiCall<unknown>(opts, "Note", "delete", {
        where: [["id", "=", id]],
      })
    },
  }
}

export function createOpenIDClient(opts: CiviCRMClientOptions) {
  return {
    async get(params: {
      where?: WhereClause
      select?: (keyof OpenID)[]
      orderBy?: OrderByClause
      limit?: number
      offset?: number
    } = {}): Promise<OpenID[]> {
      const raw = await apiCall<unknown>(opts, "OpenID", "get", params)
      return raw.map((v) => OpenIDSchema.parse(v))
    },

    async getById(id: number, select?: (keyof OpenID)[]): Promise<OpenID> {
      const results = await apiCall<unknown>(opts, "OpenID", "get", {
        where: [["id", "=", id]],
        ...(select ? { select } : {}),
      })
      if (!results.length) throw new Error(`OpenID ${id} not found`)
      return OpenIDSchema.parse(results[0])
    },

    async create(values: Partial<Omit<OpenID, "id">>): Promise<OpenID> {
      const results = await apiCall<unknown>(opts, "OpenID", "create", { values })
      if (!results.length) throw new Error("No value returned from OpenID.create")
      return OpenIDSchema.parse(results[0])
    },

    async update(id: number, values: Partial<Omit<OpenID, "id">>): Promise<OpenID> {
      const results = await apiCall<unknown>(opts, "OpenID", "update", {
        where: [["id", "=", id]],
        values,
      })
      if (!results.length) throw new Error("No value returned from OpenID.update")
      return OpenIDSchema.parse(results[0])
    },

    async delete(id: number): Promise<void> {
      await apiCall<unknown>(opts, "OpenID", "delete", {
        where: [["id", "=", id]],
      })
    },
  }
}

export function createOptionGroupClient(opts: CiviCRMClientOptions) {
  return {
    async get(params: {
      where?: WhereClause
      select?: (keyof OptionGroup)[]
      orderBy?: OrderByClause
      limit?: number
      offset?: number
    } = {}): Promise<OptionGroup[]> {
      const raw = await apiCall<unknown>(opts, "OptionGroup", "get", params)
      return raw.map((v) => OptionGroupSchema.parse(v))
    },

    async getById(id: number, select?: (keyof OptionGroup)[]): Promise<OptionGroup> {
      const results = await apiCall<unknown>(opts, "OptionGroup", "get", {
        where: [["id", "=", id]],
        ...(select ? { select } : {}),
      })
      if (!results.length) throw new Error(`OptionGroup ${id} not found`)
      return OptionGroupSchema.parse(results[0])
    },

    async create(values: Partial<Omit<OptionGroup, "id">>): Promise<OptionGroup> {
      const results = await apiCall<unknown>(opts, "OptionGroup", "create", { values })
      if (!results.length) throw new Error("No value returned from OptionGroup.create")
      return OptionGroupSchema.parse(results[0])
    },

    async update(id: number, values: Partial<Omit<OptionGroup, "id">>): Promise<OptionGroup> {
      const results = await apiCall<unknown>(opts, "OptionGroup", "update", {
        where: [["id", "=", id]],
        values,
      })
      if (!results.length) throw new Error("No value returned from OptionGroup.update")
      return OptionGroupSchema.parse(results[0])
    },

    async delete(id: number): Promise<void> {
      await apiCall<unknown>(opts, "OptionGroup", "delete", {
        where: [["id", "=", id]],
      })
    },
  }
}

export function createOptionValueClient(opts: CiviCRMClientOptions) {
  return {
    async get(params: {
      where?: WhereClause
      select?: (keyof OptionValue)[]
      orderBy?: OrderByClause
      limit?: number
      offset?: number
    } = {}): Promise<OptionValue[]> {
      const raw = await apiCall<unknown>(opts, "OptionValue", "get", params)
      return raw.map((v) => OptionValueSchema.parse(v))
    },

    async getById(id: number, select?: (keyof OptionValue)[]): Promise<OptionValue> {
      const results = await apiCall<unknown>(opts, "OptionValue", "get", {
        where: [["id", "=", id]],
        ...(select ? { select } : {}),
      })
      if (!results.length) throw new Error(`OptionValue ${id} not found`)
      return OptionValueSchema.parse(results[0])
    },

    async create(values: Partial<Omit<OptionValue, "id">>): Promise<OptionValue> {
      const results = await apiCall<unknown>(opts, "OptionValue", "create", { values })
      if (!results.length) throw new Error("No value returned from OptionValue.create")
      return OptionValueSchema.parse(results[0])
    },

    async update(id: number, values: Partial<Omit<OptionValue, "id">>): Promise<OptionValue> {
      const results = await apiCall<unknown>(opts, "OptionValue", "update", {
        where: [["id", "=", id]],
        values,
      })
      if (!results.length) throw new Error("No value returned from OptionValue.update")
      return OptionValueSchema.parse(results[0])
    },

    async delete(id: number): Promise<void> {
      await apiCall<unknown>(opts, "OptionValue", "delete", {
        where: [["id", "=", id]],
      })
    },
  }
}

export function createOrganizationClient(opts: CiviCRMClientOptions) {
  return {
    async get(params: {
      where?: WhereClause
      select?: (keyof Organization)[]
      orderBy?: OrderByClause
      limit?: number
      offset?: number
    } = {}): Promise<Organization[]> {
      const raw = await apiCall<unknown>(opts, "Organization", "get", params)
      return raw.map((v) => OrganizationSchema.parse(v))
    },

    async getById(id: number, select?: (keyof Organization)[]): Promise<Organization> {
      const results = await apiCall<unknown>(opts, "Organization", "get", {
        where: [["id", "=", id]],
        ...(select ? { select } : {}),
      })
      if (!results.length) throw new Error(`Organization ${id} not found`)
      return OrganizationSchema.parse(results[0])
    },

    async create(values: Partial<Omit<Organization, "id">>): Promise<Organization> {
      const results = await apiCall<unknown>(opts, "Organization", "create", { values })
      if (!results.length) throw new Error("No value returned from Organization.create")
      return OrganizationSchema.parse(results[0])
    },

    async update(id: number, values: Partial<Omit<Organization, "id">>): Promise<Organization> {
      const results = await apiCall<unknown>(opts, "Organization", "update", {
        where: [["id", "=", id]],
        values,
      })
      if (!results.length) throw new Error("No value returned from Organization.update")
      return OrganizationSchema.parse(results[0])
    },

    async delete(id: number): Promise<void> {
      await apiCall<unknown>(opts, "Organization", "delete", {
        where: [["id", "=", id]],
      })
    },
  }
}

export function createPCPClient(opts: CiviCRMClientOptions) {
  return {
    async get(params: {
      where?: WhereClause
      select?: (keyof PCP)[]
      orderBy?: OrderByClause
      limit?: number
      offset?: number
    } = {}): Promise<PCP[]> {
      const raw = await apiCall<unknown>(opts, "PCP", "get", params)
      return raw.map((v) => PCPSchema.parse(v))
    },

    async getById(id: number, select?: (keyof PCP)[]): Promise<PCP> {
      const results = await apiCall<unknown>(opts, "PCP", "get", {
        where: [["id", "=", id]],
        ...(select ? { select } : {}),
      })
      if (!results.length) throw new Error(`PCP ${id} not found`)
      return PCPSchema.parse(results[0])
    },

    async create(values: Partial<Omit<PCP, "id">>): Promise<PCP> {
      const results = await apiCall<unknown>(opts, "PCP", "create", { values })
      if (!results.length) throw new Error("No value returned from PCP.create")
      return PCPSchema.parse(results[0])
    },

    async update(id: number, values: Partial<Omit<PCP, "id">>): Promise<PCP> {
      const results = await apiCall<unknown>(opts, "PCP", "update", {
        where: [["id", "=", id]],
        values,
      })
      if (!results.length) throw new Error("No value returned from PCP.update")
      return PCPSchema.parse(results[0])
    },

    async delete(id: number): Promise<void> {
      await apiCall<unknown>(opts, "PCP", "delete", {
        where: [["id", "=", id]],
      })
    },
  }
}

export function createPCPBlockClient(opts: CiviCRMClientOptions) {
  return {
    async get(params: {
      where?: WhereClause
      select?: (keyof PCPBlock)[]
      orderBy?: OrderByClause
      limit?: number
      offset?: number
    } = {}): Promise<PCPBlock[]> {
      const raw = await apiCall<unknown>(opts, "PCPBlock", "get", params)
      return raw.map((v) => PCPBlockSchema.parse(v))
    },

    async getById(id: number, select?: (keyof PCPBlock)[]): Promise<PCPBlock> {
      const results = await apiCall<unknown>(opts, "PCPBlock", "get", {
        where: [["id", "=", id]],
        ...(select ? { select } : {}),
      })
      if (!results.length) throw new Error(`PCPBlock ${id} not found`)
      return PCPBlockSchema.parse(results[0])
    },

    async create(values: Partial<Omit<PCPBlock, "id">>): Promise<PCPBlock> {
      const results = await apiCall<unknown>(opts, "PCPBlock", "create", { values })
      if (!results.length) throw new Error("No value returned from PCPBlock.create")
      return PCPBlockSchema.parse(results[0])
    },

    async update(id: number, values: Partial<Omit<PCPBlock, "id">>): Promise<PCPBlock> {
      const results = await apiCall<unknown>(opts, "PCPBlock", "update", {
        where: [["id", "=", id]],
        values,
      })
      if (!results.length) throw new Error("No value returned from PCPBlock.update")
      return PCPBlockSchema.parse(results[0])
    },

    async delete(id: number): Promise<void> {
      await apiCall<unknown>(opts, "PCPBlock", "delete", {
        where: [["id", "=", id]],
      })
    },
  }
}

export function createParticipantClient(opts: CiviCRMClientOptions) {
  return {
    async get(params: {
      where?: WhereClause
      select?: (keyof Participant)[]
      orderBy?: OrderByClause
      limit?: number
      offset?: number
    } = {}): Promise<Participant[]> {
      const raw = await apiCall<unknown>(opts, "Participant", "get", params)
      return raw.map((v) => ParticipantSchema.parse(v))
    },

    async getById(id: number, select?: (keyof Participant)[]): Promise<Participant> {
      const results = await apiCall<unknown>(opts, "Participant", "get", {
        where: [["id", "=", id]],
        ...(select ? { select } : {}),
      })
      if (!results.length) throw new Error(`Participant ${id} not found`)
      return ParticipantSchema.parse(results[0])
    },

    async create(values: Partial<Omit<Participant, "id">>): Promise<Participant> {
      const results = await apiCall<unknown>(opts, "Participant", "create", { values })
      if (!results.length) throw new Error("No value returned from Participant.create")
      return ParticipantSchema.parse(results[0])
    },

    async update(id: number, values: Partial<Omit<Participant, "id">>): Promise<Participant> {
      const results = await apiCall<unknown>(opts, "Participant", "update", {
        where: [["id", "=", id]],
        values,
      })
      if (!results.length) throw new Error("No value returned from Participant.update")
      return ParticipantSchema.parse(results[0])
    },

    async delete(id: number): Promise<void> {
      await apiCall<unknown>(opts, "Participant", "delete", {
        where: [["id", "=", id]],
      })
    },
  }
}

export function createParticipantStatusTypeClient(opts: CiviCRMClientOptions) {
  return {
    async get(params: {
      where?: WhereClause
      select?: (keyof ParticipantStatusType)[]
      orderBy?: OrderByClause
      limit?: number
      offset?: number
    } = {}): Promise<ParticipantStatusType[]> {
      const raw = await apiCall<unknown>(opts, "ParticipantStatusType", "get", params)
      return raw.map((v) => ParticipantStatusTypeSchema.parse(v))
    },

    async getById(id: number, select?: (keyof ParticipantStatusType)[]): Promise<ParticipantStatusType> {
      const results = await apiCall<unknown>(opts, "ParticipantStatusType", "get", {
        where: [["id", "=", id]],
        ...(select ? { select } : {}),
      })
      if (!results.length) throw new Error(`ParticipantStatusType ${id} not found`)
      return ParticipantStatusTypeSchema.parse(results[0])
    },

    async create(values: Partial<Omit<ParticipantStatusType, "id">>): Promise<ParticipantStatusType> {
      const results = await apiCall<unknown>(opts, "ParticipantStatusType", "create", { values })
      if (!results.length) throw new Error("No value returned from ParticipantStatusType.create")
      return ParticipantStatusTypeSchema.parse(results[0])
    },

    async update(id: number, values: Partial<Omit<ParticipantStatusType, "id">>): Promise<ParticipantStatusType> {
      const results = await apiCall<unknown>(opts, "ParticipantStatusType", "update", {
        where: [["id", "=", id]],
        values,
      })
      if (!results.length) throw new Error("No value returned from ParticipantStatusType.update")
      return ParticipantStatusTypeSchema.parse(results[0])
    },

    async delete(id: number): Promise<void> {
      await apiCall<unknown>(opts, "ParticipantStatusType", "delete", {
        where: [["id", "=", id]],
      })
    },
  }
}

export function createPaymentClient(opts: CiviCRMClientOptions) {
  return {
    async get(params: {
      where?: WhereClause
      select?: (keyof Payment)[]
      orderBy?: OrderByClause
      limit?: number
      offset?: number
    } = {}): Promise<Payment[]> {
      const raw = await apiCall<unknown>(opts, "Payment", "get", params)
      return raw.map((v) => PaymentSchema.parse(v))
    },

    async getById(id: number, select?: (keyof Payment)[]): Promise<Payment> {
      const results = await apiCall<unknown>(opts, "Payment", "get", {
        where: [["id", "=", id]],
        ...(select ? { select } : {}),
      })
      if (!results.length) throw new Error(`Payment ${id} not found`)
      return PaymentSchema.parse(results[0])
    },

    async create(values: Partial<Omit<Payment, "id">>): Promise<Payment> {
      const results = await apiCall<unknown>(opts, "Payment", "create", { values })
      if (!results.length) throw new Error("No value returned from Payment.create")
      return PaymentSchema.parse(results[0])
    },

    async update(id: number, values: Partial<Omit<Payment, "id">>): Promise<Payment> {
      const results = await apiCall<unknown>(opts, "Payment", "update", {
        where: [["id", "=", id]],
        values,
      })
      if (!results.length) throw new Error("No value returned from Payment.update")
      return PaymentSchema.parse(results[0])
    },

    async delete(id: number): Promise<void> {
      await apiCall<unknown>(opts, "Payment", "delete", {
        where: [["id", "=", id]],
      })
    },
  }
}

export function createPaymentProcessorClient(opts: CiviCRMClientOptions) {
  return {
    async get(params: {
      where?: WhereClause
      select?: (keyof PaymentProcessor)[]
      orderBy?: OrderByClause
      limit?: number
      offset?: number
    } = {}): Promise<PaymentProcessor[]> {
      const raw = await apiCall<unknown>(opts, "PaymentProcessor", "get", params)
      return raw.map((v) => PaymentProcessorSchema.parse(v))
    },

    async getById(id: number, select?: (keyof PaymentProcessor)[]): Promise<PaymentProcessor> {
      const results = await apiCall<unknown>(opts, "PaymentProcessor", "get", {
        where: [["id", "=", id]],
        ...(select ? { select } : {}),
      })
      if (!results.length) throw new Error(`PaymentProcessor ${id} not found`)
      return PaymentProcessorSchema.parse(results[0])
    },

    async create(values: Partial<Omit<PaymentProcessor, "id">>): Promise<PaymentProcessor> {
      const results = await apiCall<unknown>(opts, "PaymentProcessor", "create", { values })
      if (!results.length) throw new Error("No value returned from PaymentProcessor.create")
      return PaymentProcessorSchema.parse(results[0])
    },

    async update(id: number, values: Partial<Omit<PaymentProcessor, "id">>): Promise<PaymentProcessor> {
      const results = await apiCall<unknown>(opts, "PaymentProcessor", "update", {
        where: [["id", "=", id]],
        values,
      })
      if (!results.length) throw new Error("No value returned from PaymentProcessor.update")
      return PaymentProcessorSchema.parse(results[0])
    },

    async delete(id: number): Promise<void> {
      await apiCall<unknown>(opts, "PaymentProcessor", "delete", {
        where: [["id", "=", id]],
      })
    },
  }
}

export function createPaymentProcessorTypeClient(opts: CiviCRMClientOptions) {
  return {
    async get(params: {
      where?: WhereClause
      select?: (keyof PaymentProcessorType)[]
      orderBy?: OrderByClause
      limit?: number
      offset?: number
    } = {}): Promise<PaymentProcessorType[]> {
      const raw = await apiCall<unknown>(opts, "PaymentProcessorType", "get", params)
      return raw.map((v) => PaymentProcessorTypeSchema.parse(v))
    },

    async getById(id: number, select?: (keyof PaymentProcessorType)[]): Promise<PaymentProcessorType> {
      const results = await apiCall<unknown>(opts, "PaymentProcessorType", "get", {
        where: [["id", "=", id]],
        ...(select ? { select } : {}),
      })
      if (!results.length) throw new Error(`PaymentProcessorType ${id} not found`)
      return PaymentProcessorTypeSchema.parse(results[0])
    },

    async create(values: Partial<Omit<PaymentProcessorType, "id">>): Promise<PaymentProcessorType> {
      const results = await apiCall<unknown>(opts, "PaymentProcessorType", "create", { values })
      if (!results.length) throw new Error("No value returned from PaymentProcessorType.create")
      return PaymentProcessorTypeSchema.parse(results[0])
    },

    async update(id: number, values: Partial<Omit<PaymentProcessorType, "id">>): Promise<PaymentProcessorType> {
      const results = await apiCall<unknown>(opts, "PaymentProcessorType", "update", {
        where: [["id", "=", id]],
        values,
      })
      if (!results.length) throw new Error("No value returned from PaymentProcessorType.update")
      return PaymentProcessorTypeSchema.parse(results[0])
    },

    async delete(id: number): Promise<void> {
      await apiCall<unknown>(opts, "PaymentProcessorType", "delete", {
        where: [["id", "=", id]],
      })
    },
  }
}

export function createPaymentTokenClient(opts: CiviCRMClientOptions) {
  return {
    async get(params: {
      where?: WhereClause
      select?: (keyof PaymentToken)[]
      orderBy?: OrderByClause
      limit?: number
      offset?: number
    } = {}): Promise<PaymentToken[]> {
      const raw = await apiCall<unknown>(opts, "PaymentToken", "get", params)
      return raw.map((v) => PaymentTokenSchema.parse(v))
    },

    async getById(id: number, select?: (keyof PaymentToken)[]): Promise<PaymentToken> {
      const results = await apiCall<unknown>(opts, "PaymentToken", "get", {
        where: [["id", "=", id]],
        ...(select ? { select } : {}),
      })
      if (!results.length) throw new Error(`PaymentToken ${id} not found`)
      return PaymentTokenSchema.parse(results[0])
    },

    async create(values: Partial<Omit<PaymentToken, "id">>): Promise<PaymentToken> {
      const results = await apiCall<unknown>(opts, "PaymentToken", "create", { values })
      if (!results.length) throw new Error("No value returned from PaymentToken.create")
      return PaymentTokenSchema.parse(results[0])
    },

    async update(id: number, values: Partial<Omit<PaymentToken, "id">>): Promise<PaymentToken> {
      const results = await apiCall<unknown>(opts, "PaymentToken", "update", {
        where: [["id", "=", id]],
        values,
      })
      if (!results.length) throw new Error("No value returned from PaymentToken.update")
      return PaymentTokenSchema.parse(results[0])
    },

    async delete(id: number): Promise<void> {
      await apiCall<unknown>(opts, "PaymentToken", "delete", {
        where: [["id", "=", id]],
      })
    },
  }
}

export function createPermissionClient(opts: CiviCRMClientOptions) {
  return {
    async get(params: {
      where?: WhereClause
      select?: (keyof Permission)[]
      orderBy?: OrderByClause
      limit?: number
      offset?: number
    } = {}): Promise<Permission[]> {
      const raw = await apiCall<unknown>(opts, "Permission", "get", params)
      return raw.map((v) => PermissionSchema.parse(v))
    },

    async getById(id: number, select?: (keyof Permission)[]): Promise<Permission> {
      const results = await apiCall<unknown>(opts, "Permission", "get", {
        where: [["id", "=", id]],
        ...(select ? { select } : {}),
      })
      if (!results.length) throw new Error(`Permission ${id} not found`)
      return PermissionSchema.parse(results[0])
    },

    async create(values: Partial<Omit<Permission, "id">>): Promise<Permission> {
      const results = await apiCall<unknown>(opts, "Permission", "create", { values })
      if (!results.length) throw new Error("No value returned from Permission.create")
      return PermissionSchema.parse(results[0])
    },

    async update(id: number, values: Partial<Omit<Permission, "id">>): Promise<Permission> {
      const results = await apiCall<unknown>(opts, "Permission", "update", {
        where: [["id", "=", id]],
        values,
      })
      if (!results.length) throw new Error("No value returned from Permission.update")
      return PermissionSchema.parse(results[0])
    },

    async delete(id: number): Promise<void> {
      await apiCall<unknown>(opts, "Permission", "delete", {
        where: [["id", "=", id]],
      })
    },
  }
}

export function createPhoneClient(opts: CiviCRMClientOptions) {
  return {
    async get(params: {
      where?: WhereClause
      select?: (keyof Phone)[]
      orderBy?: OrderByClause
      limit?: number
      offset?: number
    } = {}): Promise<Phone[]> {
      const raw = await apiCall<unknown>(opts, "Phone", "get", params)
      return raw.map((v) => PhoneSchema.parse(v))
    },

    async getById(id: number, select?: (keyof Phone)[]): Promise<Phone> {
      const results = await apiCall<unknown>(opts, "Phone", "get", {
        where: [["id", "=", id]],
        ...(select ? { select } : {}),
      })
      if (!results.length) throw new Error(`Phone ${id} not found`)
      return PhoneSchema.parse(results[0])
    },

    async create(values: Partial<Omit<Phone, "id">>): Promise<Phone> {
      const results = await apiCall<unknown>(opts, "Phone", "create", { values })
      if (!results.length) throw new Error("No value returned from Phone.create")
      return PhoneSchema.parse(results[0])
    },

    async update(id: number, values: Partial<Omit<Phone, "id">>): Promise<Phone> {
      const results = await apiCall<unknown>(opts, "Phone", "update", {
        where: [["id", "=", id]],
        values,
      })
      if (!results.length) throw new Error("No value returned from Phone.update")
      return PhoneSchema.parse(results[0])
    },

    async delete(id: number): Promise<void> {
      await apiCall<unknown>(opts, "Phone", "delete", {
        where: [["id", "=", id]],
      })
    },
  }
}

export function createPreferencesDateClient(opts: CiviCRMClientOptions) {
  return {
    async get(params: {
      where?: WhereClause
      select?: (keyof PreferencesDate)[]
      orderBy?: OrderByClause
      limit?: number
      offset?: number
    } = {}): Promise<PreferencesDate[]> {
      const raw = await apiCall<unknown>(opts, "PreferencesDate", "get", params)
      return raw.map((v) => PreferencesDateSchema.parse(v))
    },

    async getById(id: number, select?: (keyof PreferencesDate)[]): Promise<PreferencesDate> {
      const results = await apiCall<unknown>(opts, "PreferencesDate", "get", {
        where: [["id", "=", id]],
        ...(select ? { select } : {}),
      })
      if (!results.length) throw new Error(`PreferencesDate ${id} not found`)
      return PreferencesDateSchema.parse(results[0])
    },

    async create(values: Partial<Omit<PreferencesDate, "id">>): Promise<PreferencesDate> {
      const results = await apiCall<unknown>(opts, "PreferencesDate", "create", { values })
      if (!results.length) throw new Error("No value returned from PreferencesDate.create")
      return PreferencesDateSchema.parse(results[0])
    },

    async update(id: number, values: Partial<Omit<PreferencesDate, "id">>): Promise<PreferencesDate> {
      const results = await apiCall<unknown>(opts, "PreferencesDate", "update", {
        where: [["id", "=", id]],
        values,
      })
      if (!results.length) throw new Error("No value returned from PreferencesDate.update")
      return PreferencesDateSchema.parse(results[0])
    },

    async delete(id: number): Promise<void> {
      await apiCall<unknown>(opts, "PreferencesDate", "delete", {
        where: [["id", "=", id]],
      })
    },
  }
}

export function createPremiumClient(opts: CiviCRMClientOptions) {
  return {
    async get(params: {
      where?: WhereClause
      select?: (keyof Premium)[]
      orderBy?: OrderByClause
      limit?: number
      offset?: number
    } = {}): Promise<Premium[]> {
      const raw = await apiCall<unknown>(opts, "Premium", "get", params)
      return raw.map((v) => PremiumSchema.parse(v))
    },

    async getById(id: number, select?: (keyof Premium)[]): Promise<Premium> {
      const results = await apiCall<unknown>(opts, "Premium", "get", {
        where: [["id", "=", id]],
        ...(select ? { select } : {}),
      })
      if (!results.length) throw new Error(`Premium ${id} not found`)
      return PremiumSchema.parse(results[0])
    },

    async create(values: Partial<Omit<Premium, "id">>): Promise<Premium> {
      const results = await apiCall<unknown>(opts, "Premium", "create", { values })
      if (!results.length) throw new Error("No value returned from Premium.create")
      return PremiumSchema.parse(results[0])
    },

    async update(id: number, values: Partial<Omit<Premium, "id">>): Promise<Premium> {
      const results = await apiCall<unknown>(opts, "Premium", "update", {
        where: [["id", "=", id]],
        values,
      })
      if (!results.length) throw new Error("No value returned from Premium.update")
      return PremiumSchema.parse(results[0])
    },

    async delete(id: number): Promise<void> {
      await apiCall<unknown>(opts, "Premium", "delete", {
        where: [["id", "=", id]],
      })
    },
  }
}

export function createPremiumsProductClient(opts: CiviCRMClientOptions) {
  return {
    async get(params: {
      where?: WhereClause
      select?: (keyof PremiumsProduct)[]
      orderBy?: OrderByClause
      limit?: number
      offset?: number
    } = {}): Promise<PremiumsProduct[]> {
      const raw = await apiCall<unknown>(opts, "PremiumsProduct", "get", params)
      return raw.map((v) => PremiumsProductSchema.parse(v))
    },

    async getById(id: number, select?: (keyof PremiumsProduct)[]): Promise<PremiumsProduct> {
      const results = await apiCall<unknown>(opts, "PremiumsProduct", "get", {
        where: [["id", "=", id]],
        ...(select ? { select } : {}),
      })
      if (!results.length) throw new Error(`PremiumsProduct ${id} not found`)
      return PremiumsProductSchema.parse(results[0])
    },

    async create(values: Partial<Omit<PremiumsProduct, "id">>): Promise<PremiumsProduct> {
      const results = await apiCall<unknown>(opts, "PremiumsProduct", "create", { values })
      if (!results.length) throw new Error("No value returned from PremiumsProduct.create")
      return PremiumsProductSchema.parse(results[0])
    },

    async update(id: number, values: Partial<Omit<PremiumsProduct, "id">>): Promise<PremiumsProduct> {
      const results = await apiCall<unknown>(opts, "PremiumsProduct", "update", {
        where: [["id", "=", id]],
        values,
      })
      if (!results.length) throw new Error("No value returned from PremiumsProduct.update")
      return PremiumsProductSchema.parse(results[0])
    },

    async delete(id: number): Promise<void> {
      await apiCall<unknown>(opts, "PremiumsProduct", "delete", {
        where: [["id", "=", id]],
      })
    },
  }
}

export function createPriceFieldClient(opts: CiviCRMClientOptions) {
  return {
    async get(params: {
      where?: WhereClause
      select?: (keyof PriceField)[]
      orderBy?: OrderByClause
      limit?: number
      offset?: number
    } = {}): Promise<PriceField[]> {
      const raw = await apiCall<unknown>(opts, "PriceField", "get", params)
      return raw.map((v) => PriceFieldSchema.parse(v))
    },

    async getById(id: number, select?: (keyof PriceField)[]): Promise<PriceField> {
      const results = await apiCall<unknown>(opts, "PriceField", "get", {
        where: [["id", "=", id]],
        ...(select ? { select } : {}),
      })
      if (!results.length) throw new Error(`PriceField ${id} not found`)
      return PriceFieldSchema.parse(results[0])
    },

    async create(values: Partial<Omit<PriceField, "id">>): Promise<PriceField> {
      const results = await apiCall<unknown>(opts, "PriceField", "create", { values })
      if (!results.length) throw new Error("No value returned from PriceField.create")
      return PriceFieldSchema.parse(results[0])
    },

    async update(id: number, values: Partial<Omit<PriceField, "id">>): Promise<PriceField> {
      const results = await apiCall<unknown>(opts, "PriceField", "update", {
        where: [["id", "=", id]],
        values,
      })
      if (!results.length) throw new Error("No value returned from PriceField.update")
      return PriceFieldSchema.parse(results[0])
    },

    async delete(id: number): Promise<void> {
      await apiCall<unknown>(opts, "PriceField", "delete", {
        where: [["id", "=", id]],
      })
    },
  }
}

export function createPriceFieldValueClient(opts: CiviCRMClientOptions) {
  return {
    async get(params: {
      where?: WhereClause
      select?: (keyof PriceFieldValue)[]
      orderBy?: OrderByClause
      limit?: number
      offset?: number
    } = {}): Promise<PriceFieldValue[]> {
      const raw = await apiCall<unknown>(opts, "PriceFieldValue", "get", params)
      return raw.map((v) => PriceFieldValueSchema.parse(v))
    },

    async getById(id: number, select?: (keyof PriceFieldValue)[]): Promise<PriceFieldValue> {
      const results = await apiCall<unknown>(opts, "PriceFieldValue", "get", {
        where: [["id", "=", id]],
        ...(select ? { select } : {}),
      })
      if (!results.length) throw new Error(`PriceFieldValue ${id} not found`)
      return PriceFieldValueSchema.parse(results[0])
    },

    async create(values: Partial<Omit<PriceFieldValue, "id">>): Promise<PriceFieldValue> {
      const results = await apiCall<unknown>(opts, "PriceFieldValue", "create", { values })
      if (!results.length) throw new Error("No value returned from PriceFieldValue.create")
      return PriceFieldValueSchema.parse(results[0])
    },

    async update(id: number, values: Partial<Omit<PriceFieldValue, "id">>): Promise<PriceFieldValue> {
      const results = await apiCall<unknown>(opts, "PriceFieldValue", "update", {
        where: [["id", "=", id]],
        values,
      })
      if (!results.length) throw new Error("No value returned from PriceFieldValue.update")
      return PriceFieldValueSchema.parse(results[0])
    },

    async delete(id: number): Promise<void> {
      await apiCall<unknown>(opts, "PriceFieldValue", "delete", {
        where: [["id", "=", id]],
      })
    },
  }
}

export function createPriceSetClient(opts: CiviCRMClientOptions) {
  return {
    async get(params: {
      where?: WhereClause
      select?: (keyof PriceSet)[]
      orderBy?: OrderByClause
      limit?: number
      offset?: number
    } = {}): Promise<PriceSet[]> {
      const raw = await apiCall<unknown>(opts, "PriceSet", "get", params)
      return raw.map((v) => PriceSetSchema.parse(v))
    },

    async getById(id: number, select?: (keyof PriceSet)[]): Promise<PriceSet> {
      const results = await apiCall<unknown>(opts, "PriceSet", "get", {
        where: [["id", "=", id]],
        ...(select ? { select } : {}),
      })
      if (!results.length) throw new Error(`PriceSet ${id} not found`)
      return PriceSetSchema.parse(results[0])
    },

    async create(values: Partial<Omit<PriceSet, "id">>): Promise<PriceSet> {
      const results = await apiCall<unknown>(opts, "PriceSet", "create", { values })
      if (!results.length) throw new Error("No value returned from PriceSet.create")
      return PriceSetSchema.parse(results[0])
    },

    async update(id: number, values: Partial<Omit<PriceSet, "id">>): Promise<PriceSet> {
      const results = await apiCall<unknown>(opts, "PriceSet", "update", {
        where: [["id", "=", id]],
        values,
      })
      if (!results.length) throw new Error("No value returned from PriceSet.update")
      return PriceSetSchema.parse(results[0])
    },

    async delete(id: number): Promise<void> {
      await apiCall<unknown>(opts, "PriceSet", "delete", {
        where: [["id", "=", id]],
      })
    },
  }
}

export function createPriceSetEntityClient(opts: CiviCRMClientOptions) {
  return {
    async get(params: {
      where?: WhereClause
      select?: (keyof PriceSetEntity)[]
      orderBy?: OrderByClause
      limit?: number
      offset?: number
    } = {}): Promise<PriceSetEntity[]> {
      const raw = await apiCall<unknown>(opts, "PriceSetEntity", "get", params)
      return raw.map((v) => PriceSetEntitySchema.parse(v))
    },

    async getById(id: number, select?: (keyof PriceSetEntity)[]): Promise<PriceSetEntity> {
      const results = await apiCall<unknown>(opts, "PriceSetEntity", "get", {
        where: [["id", "=", id]],
        ...(select ? { select } : {}),
      })
      if (!results.length) throw new Error(`PriceSetEntity ${id} not found`)
      return PriceSetEntitySchema.parse(results[0])
    },

    async create(values: Partial<Omit<PriceSetEntity, "id">>): Promise<PriceSetEntity> {
      const results = await apiCall<unknown>(opts, "PriceSetEntity", "create", { values })
      if (!results.length) throw new Error("No value returned from PriceSetEntity.create")
      return PriceSetEntitySchema.parse(results[0])
    },

    async update(id: number, values: Partial<Omit<PriceSetEntity, "id">>): Promise<PriceSetEntity> {
      const results = await apiCall<unknown>(opts, "PriceSetEntity", "update", {
        where: [["id", "=", id]],
        values,
      })
      if (!results.length) throw new Error("No value returned from PriceSetEntity.update")
      return PriceSetEntitySchema.parse(results[0])
    },

    async delete(id: number): Promise<void> {
      await apiCall<unknown>(opts, "PriceSetEntity", "delete", {
        where: [["id", "=", id]],
      })
    },
  }
}

export function createPrintLabelClient(opts: CiviCRMClientOptions) {
  return {
    async get(params: {
      where?: WhereClause
      select?: (keyof PrintLabel)[]
      orderBy?: OrderByClause
      limit?: number
      offset?: number
    } = {}): Promise<PrintLabel[]> {
      const raw = await apiCall<unknown>(opts, "PrintLabel", "get", params)
      return raw.map((v) => PrintLabelSchema.parse(v))
    },

    async getById(id: number, select?: (keyof PrintLabel)[]): Promise<PrintLabel> {
      const results = await apiCall<unknown>(opts, "PrintLabel", "get", {
        where: [["id", "=", id]],
        ...(select ? { select } : {}),
      })
      if (!results.length) throw new Error(`PrintLabel ${id} not found`)
      return PrintLabelSchema.parse(results[0])
    },

    async create(values: Partial<Omit<PrintLabel, "id">>): Promise<PrintLabel> {
      const results = await apiCall<unknown>(opts, "PrintLabel", "create", { values })
      if (!results.length) throw new Error("No value returned from PrintLabel.create")
      return PrintLabelSchema.parse(results[0])
    },

    async update(id: number, values: Partial<Omit<PrintLabel, "id">>): Promise<PrintLabel> {
      const results = await apiCall<unknown>(opts, "PrintLabel", "update", {
        where: [["id", "=", id]],
        values,
      })
      if (!results.length) throw new Error("No value returned from PrintLabel.update")
      return PrintLabelSchema.parse(results[0])
    },

    async delete(id: number): Promise<void> {
      await apiCall<unknown>(opts, "PrintLabel", "delete", {
        where: [["id", "=", id]],
      })
    },
  }
}

export function createProductClient(opts: CiviCRMClientOptions) {
  return {
    async get(params: {
      where?: WhereClause
      select?: (keyof Product)[]
      orderBy?: OrderByClause
      limit?: number
      offset?: number
    } = {}): Promise<Product[]> {
      const raw = await apiCall<unknown>(opts, "Product", "get", params)
      return raw.map((v) => ProductSchema.parse(v))
    },

    async getById(id: number, select?: (keyof Product)[]): Promise<Product> {
      const results = await apiCall<unknown>(opts, "Product", "get", {
        where: [["id", "=", id]],
        ...(select ? { select } : {}),
      })
      if (!results.length) throw new Error(`Product ${id} not found`)
      return ProductSchema.parse(results[0])
    },

    async create(values: Partial<Omit<Product, "id">>): Promise<Product> {
      const results = await apiCall<unknown>(opts, "Product", "create", { values })
      if (!results.length) throw new Error("No value returned from Product.create")
      return ProductSchema.parse(results[0])
    },

    async update(id: number, values: Partial<Omit<Product, "id">>): Promise<Product> {
      const results = await apiCall<unknown>(opts, "Product", "update", {
        where: [["id", "=", id]],
        values,
      })
      if (!results.length) throw new Error("No value returned from Product.update")
      return ProductSchema.parse(results[0])
    },

    async delete(id: number): Promise<void> {
      await apiCall<unknown>(opts, "Product", "delete", {
        where: [["id", "=", id]],
      })
    },
  }
}

export function createQueueClient(opts: CiviCRMClientOptions) {
  return {
    async get(params: {
      where?: WhereClause
      select?: (keyof Queue)[]
      orderBy?: OrderByClause
      limit?: number
      offset?: number
    } = {}): Promise<Queue[]> {
      const raw = await apiCall<unknown>(opts, "Queue", "get", params)
      return raw.map((v) => QueueSchema.parse(v))
    },

    async getById(id: number, select?: (keyof Queue)[]): Promise<Queue> {
      const results = await apiCall<unknown>(opts, "Queue", "get", {
        where: [["id", "=", id]],
        ...(select ? { select } : {}),
      })
      if (!results.length) throw new Error(`Queue ${id} not found`)
      return QueueSchema.parse(results[0])
    },

    async create(values: Partial<Omit<Queue, "id">>): Promise<Queue> {
      const results = await apiCall<unknown>(opts, "Queue", "create", { values })
      if (!results.length) throw new Error("No value returned from Queue.create")
      return QueueSchema.parse(results[0])
    },

    async update(id: number, values: Partial<Omit<Queue, "id">>): Promise<Queue> {
      const results = await apiCall<unknown>(opts, "Queue", "update", {
        where: [["id", "=", id]],
        values,
      })
      if (!results.length) throw new Error("No value returned from Queue.update")
      return QueueSchema.parse(results[0])
    },

    async delete(id: number): Promise<void> {
      await apiCall<unknown>(opts, "Queue", "delete", {
        where: [["id", "=", id]],
      })
    },
  }
}

export function createQueueItemClient(opts: CiviCRMClientOptions) {
  return {
    async get(params: {
      where?: WhereClause
      select?: (keyof QueueItem)[]
      orderBy?: OrderByClause
      limit?: number
      offset?: number
    } = {}): Promise<QueueItem[]> {
      const raw = await apiCall<unknown>(opts, "QueueItem", "get", params)
      return raw.map((v) => QueueItemSchema.parse(v))
    },

    async getById(id: number, select?: (keyof QueueItem)[]): Promise<QueueItem> {
      const results = await apiCall<unknown>(opts, "QueueItem", "get", {
        where: [["id", "=", id]],
        ...(select ? { select } : {}),
      })
      if (!results.length) throw new Error(`QueueItem ${id} not found`)
      return QueueItemSchema.parse(results[0])
    },

    async create(values: Partial<Omit<QueueItem, "id">>): Promise<QueueItem> {
      const results = await apiCall<unknown>(opts, "QueueItem", "create", { values })
      if (!results.length) throw new Error("No value returned from QueueItem.create")
      return QueueItemSchema.parse(results[0])
    },

    async update(id: number, values: Partial<Omit<QueueItem, "id">>): Promise<QueueItem> {
      const results = await apiCall<unknown>(opts, "QueueItem", "update", {
        where: [["id", "=", id]],
        values,
      })
      if (!results.length) throw new Error("No value returned from QueueItem.update")
      return QueueItemSchema.parse(results[0])
    },

    async delete(id: number): Promise<void> {
      await apiCall<unknown>(opts, "QueueItem", "delete", {
        where: [["id", "=", id]],
      })
    },
  }
}

export function createRecentItemClient(opts: CiviCRMClientOptions) {
  return {
    async get(params: {
      where?: WhereClause
      select?: (keyof RecentItem)[]
      orderBy?: OrderByClause
      limit?: number
      offset?: number
    } = {}): Promise<RecentItem[]> {
      const raw = await apiCall<unknown>(opts, "RecentItem", "get", params)
      return raw.map((v) => RecentItemSchema.parse(v))
    },

    async getById(id: number, select?: (keyof RecentItem)[]): Promise<RecentItem> {
      const results = await apiCall<unknown>(opts, "RecentItem", "get", {
        where: [["id", "=", id]],
        ...(select ? { select } : {}),
      })
      if (!results.length) throw new Error(`RecentItem ${id} not found`)
      return RecentItemSchema.parse(results[0])
    },

    async create(values: Partial<Omit<RecentItem, "id">>): Promise<RecentItem> {
      const results = await apiCall<unknown>(opts, "RecentItem", "create", { values })
      if (!results.length) throw new Error("No value returned from RecentItem.create")
      return RecentItemSchema.parse(results[0])
    },

    async update(id: number, values: Partial<Omit<RecentItem, "id">>): Promise<RecentItem> {
      const results = await apiCall<unknown>(opts, "RecentItem", "update", {
        where: [["id", "=", id]],
        values,
      })
      if (!results.length) throw new Error("No value returned from RecentItem.update")
      return RecentItemSchema.parse(results[0])
    },

    async delete(id: number): Promise<void> {
      await apiCall<unknown>(opts, "RecentItem", "delete", {
        where: [["id", "=", id]],
      })
    },
  }
}

export function createRelationshipClient(opts: CiviCRMClientOptions) {
  return {
    async get(params: {
      where?: WhereClause
      select?: (keyof Relationship)[]
      orderBy?: OrderByClause
      limit?: number
      offset?: number
    } = {}): Promise<Relationship[]> {
      const raw = await apiCall<unknown>(opts, "Relationship", "get", params)
      return raw.map((v) => RelationshipSchema.parse(v))
    },

    async getById(id: number, select?: (keyof Relationship)[]): Promise<Relationship> {
      const results = await apiCall<unknown>(opts, "Relationship", "get", {
        where: [["id", "=", id]],
        ...(select ? { select } : {}),
      })
      if (!results.length) throw new Error(`Relationship ${id} not found`)
      return RelationshipSchema.parse(results[0])
    },

    async create(values: Partial<Omit<Relationship, "id">>): Promise<Relationship> {
      const results = await apiCall<unknown>(opts, "Relationship", "create", { values })
      if (!results.length) throw new Error("No value returned from Relationship.create")
      return RelationshipSchema.parse(results[0])
    },

    async update(id: number, values: Partial<Omit<Relationship, "id">>): Promise<Relationship> {
      const results = await apiCall<unknown>(opts, "Relationship", "update", {
        where: [["id", "=", id]],
        values,
      })
      if (!results.length) throw new Error("No value returned from Relationship.update")
      return RelationshipSchema.parse(results[0])
    },

    async delete(id: number): Promise<void> {
      await apiCall<unknown>(opts, "Relationship", "delete", {
        where: [["id", "=", id]],
      })
    },
  }
}

export function createRelationshipCacheClient(opts: CiviCRMClientOptions) {
  return {
    async get(params: {
      where?: WhereClause
      select?: (keyof RelationshipCache)[]
      orderBy?: OrderByClause
      limit?: number
      offset?: number
    } = {}): Promise<RelationshipCache[]> {
      const raw = await apiCall<unknown>(opts, "RelationshipCache", "get", params)
      return raw.map((v) => RelationshipCacheSchema.parse(v))
    },

    async getById(id: number, select?: (keyof RelationshipCache)[]): Promise<RelationshipCache> {
      const results = await apiCall<unknown>(opts, "RelationshipCache", "get", {
        where: [["id", "=", id]],
        ...(select ? { select } : {}),
      })
      if (!results.length) throw new Error(`RelationshipCache ${id} not found`)
      return RelationshipCacheSchema.parse(results[0])
    },

    async create(values: Partial<Omit<RelationshipCache, "id">>): Promise<RelationshipCache> {
      const results = await apiCall<unknown>(opts, "RelationshipCache", "create", { values })
      if (!results.length) throw new Error("No value returned from RelationshipCache.create")
      return RelationshipCacheSchema.parse(results[0])
    },

    async update(id: number, values: Partial<Omit<RelationshipCache, "id">>): Promise<RelationshipCache> {
      const results = await apiCall<unknown>(opts, "RelationshipCache", "update", {
        where: [["id", "=", id]],
        values,
      })
      if (!results.length) throw new Error("No value returned from RelationshipCache.update")
      return RelationshipCacheSchema.parse(results[0])
    },

    async delete(id: number): Promise<void> {
      await apiCall<unknown>(opts, "RelationshipCache", "delete", {
        where: [["id", "=", id]],
      })
    },
  }
}

export function createRelationshipTypeClient(opts: CiviCRMClientOptions) {
  return {
    async get(params: {
      where?: WhereClause
      select?: (keyof RelationshipType)[]
      orderBy?: OrderByClause
      limit?: number
      offset?: number
    } = {}): Promise<RelationshipType[]> {
      const raw = await apiCall<unknown>(opts, "RelationshipType", "get", params)
      return raw.map((v) => RelationshipTypeSchema.parse(v))
    },

    async getById(id: number, select?: (keyof RelationshipType)[]): Promise<RelationshipType> {
      const results = await apiCall<unknown>(opts, "RelationshipType", "get", {
        where: [["id", "=", id]],
        ...(select ? { select } : {}),
      })
      if (!results.length) throw new Error(`RelationshipType ${id} not found`)
      return RelationshipTypeSchema.parse(results[0])
    },

    async create(values: Partial<Omit<RelationshipType, "id">>): Promise<RelationshipType> {
      const results = await apiCall<unknown>(opts, "RelationshipType", "create", { values })
      if (!results.length) throw new Error("No value returned from RelationshipType.create")
      return RelationshipTypeSchema.parse(results[0])
    },

    async update(id: number, values: Partial<Omit<RelationshipType, "id">>): Promise<RelationshipType> {
      const results = await apiCall<unknown>(opts, "RelationshipType", "update", {
        where: [["id", "=", id]],
        values,
      })
      if (!results.length) throw new Error("No value returned from RelationshipType.update")
      return RelationshipTypeSchema.parse(results[0])
    },

    async delete(id: number): Promise<void> {
      await apiCall<unknown>(opts, "RelationshipType", "delete", {
        where: [["id", "=", id]],
      })
    },
  }
}

export function createReportInstanceClient(opts: CiviCRMClientOptions) {
  return {
    async get(params: {
      where?: WhereClause
      select?: (keyof ReportInstance)[]
      orderBy?: OrderByClause
      limit?: number
      offset?: number
    } = {}): Promise<ReportInstance[]> {
      const raw = await apiCall<unknown>(opts, "ReportInstance", "get", params)
      return raw.map((v) => ReportInstanceSchema.parse(v))
    },

    async getById(id: number, select?: (keyof ReportInstance)[]): Promise<ReportInstance> {
      const results = await apiCall<unknown>(opts, "ReportInstance", "get", {
        where: [["id", "=", id]],
        ...(select ? { select } : {}),
      })
      if (!results.length) throw new Error(`ReportInstance ${id} not found`)
      return ReportInstanceSchema.parse(results[0])
    },

    async create(values: Partial<Omit<ReportInstance, "id">>): Promise<ReportInstance> {
      const results = await apiCall<unknown>(opts, "ReportInstance", "create", { values })
      if (!results.length) throw new Error("No value returned from ReportInstance.create")
      return ReportInstanceSchema.parse(results[0])
    },

    async update(id: number, values: Partial<Omit<ReportInstance, "id">>): Promise<ReportInstance> {
      const results = await apiCall<unknown>(opts, "ReportInstance", "update", {
        where: [["id", "=", id]],
        values,
      })
      if (!results.length) throw new Error("No value returned from ReportInstance.update")
      return ReportInstanceSchema.parse(results[0])
    },

    async delete(id: number): Promise<void> {
      await apiCall<unknown>(opts, "ReportInstance", "delete", {
        where: [["id", "=", id]],
      })
    },
  }
}

export function createRiverleaStreamClient(opts: CiviCRMClientOptions) {
  return {
    async get(params: {
      where?: WhereClause
      select?: (keyof RiverleaStream)[]
      orderBy?: OrderByClause
      limit?: number
      offset?: number
    } = {}): Promise<RiverleaStream[]> {
      const raw = await apiCall<unknown>(opts, "RiverleaStream", "get", params)
      return raw.map((v) => RiverleaStreamSchema.parse(v))
    },

    async getById(id: number, select?: (keyof RiverleaStream)[]): Promise<RiverleaStream> {
      const results = await apiCall<unknown>(opts, "RiverleaStream", "get", {
        where: [["id", "=", id]],
        ...(select ? { select } : {}),
      })
      if (!results.length) throw new Error(`RiverleaStream ${id} not found`)
      return RiverleaStreamSchema.parse(results[0])
    },

    async create(values: Partial<Omit<RiverleaStream, "id">>): Promise<RiverleaStream> {
      const results = await apiCall<unknown>(opts, "RiverleaStream", "create", { values })
      if (!results.length) throw new Error("No value returned from RiverleaStream.create")
      return RiverleaStreamSchema.parse(results[0])
    },

    async update(id: number, values: Partial<Omit<RiverleaStream, "id">>): Promise<RiverleaStream> {
      const results = await apiCall<unknown>(opts, "RiverleaStream", "update", {
        where: [["id", "=", id]],
        values,
      })
      if (!results.length) throw new Error("No value returned from RiverleaStream.update")
      return RiverleaStreamSchema.parse(results[0])
    },

    async delete(id: number): Promise<void> {
      await apiCall<unknown>(opts, "RiverleaStream", "delete", {
        where: [["id", "=", id]],
      })
    },
  }
}

export function createRoleClient(opts: CiviCRMClientOptions) {
  return {
    async get(params: {
      where?: WhereClause
      select?: (keyof Role)[]
      orderBy?: OrderByClause
      limit?: number
      offset?: number
    } = {}): Promise<Role[]> {
      const raw = await apiCall<unknown>(opts, "Role", "get", params)
      return raw.map((v) => RoleSchema.parse(v))
    },

    async getById(id: number, select?: (keyof Role)[]): Promise<Role> {
      const results = await apiCall<unknown>(opts, "Role", "get", {
        where: [["id", "=", id]],
        ...(select ? { select } : {}),
      })
      if (!results.length) throw new Error(`Role ${id} not found`)
      return RoleSchema.parse(results[0])
    },

    async create(values: Partial<Omit<Role, "id">>): Promise<Role> {
      const results = await apiCall<unknown>(opts, "Role", "create", { values })
      if (!results.length) throw new Error("No value returned from Role.create")
      return RoleSchema.parse(results[0])
    },

    async update(id: number, values: Partial<Omit<Role, "id">>): Promise<Role> {
      const results = await apiCall<unknown>(opts, "Role", "update", {
        where: [["id", "=", id]],
        values,
      })
      if (!results.length) throw new Error("No value returned from Role.update")
      return RoleSchema.parse(results[0])
    },

    async delete(id: number): Promise<void> {
      await apiCall<unknown>(opts, "Role", "delete", {
        where: [["id", "=", id]],
      })
    },
  }
}

export function createRolePermissionClient(opts: CiviCRMClientOptions) {
  return {
    async get(params: {
      where?: WhereClause
      select?: (keyof RolePermission)[]
      orderBy?: OrderByClause
      limit?: number
      offset?: number
    } = {}): Promise<RolePermission[]> {
      const raw = await apiCall<unknown>(opts, "RolePermission", "get", params)
      return raw.map((v) => RolePermissionSchema.parse(v))
    },

    async getById(id: number, select?: (keyof RolePermission)[]): Promise<RolePermission> {
      const results = await apiCall<unknown>(opts, "RolePermission", "get", {
        where: [["id", "=", id]],
        ...(select ? { select } : {}),
      })
      if (!results.length) throw new Error(`RolePermission ${id} not found`)
      return RolePermissionSchema.parse(results[0])
    },

    async create(values: Partial<Omit<RolePermission, "id">>): Promise<RolePermission> {
      const results = await apiCall<unknown>(opts, "RolePermission", "create", { values })
      if (!results.length) throw new Error("No value returned from RolePermission.create")
      return RolePermissionSchema.parse(results[0])
    },

    async update(id: number, values: Partial<Omit<RolePermission, "id">>): Promise<RolePermission> {
      const results = await apiCall<unknown>(opts, "RolePermission", "update", {
        where: [["id", "=", id]],
        values,
      })
      if (!results.length) throw new Error("No value returned from RolePermission.update")
      return RolePermissionSchema.parse(results[0])
    },

    async delete(id: number): Promise<void> {
      await apiCall<unknown>(opts, "RolePermission", "delete", {
        where: [["id", "=", id]],
      })
    },
  }
}

export function createRouteClient(opts: CiviCRMClientOptions) {
  return {
    async get(params: {
      where?: WhereClause
      select?: (keyof Route)[]
      orderBy?: OrderByClause
      limit?: number
      offset?: number
    } = {}): Promise<Route[]> {
      const raw = await apiCall<unknown>(opts, "Route", "get", params)
      return raw.map((v) => RouteSchema.parse(v))
    },

    async getById(id: number, select?: (keyof Route)[]): Promise<Route> {
      const results = await apiCall<unknown>(opts, "Route", "get", {
        where: [["id", "=", id]],
        ...(select ? { select } : {}),
      })
      if (!results.length) throw new Error(`Route ${id} not found`)
      return RouteSchema.parse(results[0])
    },

    async create(values: Partial<Omit<Route, "id">>): Promise<Route> {
      const results = await apiCall<unknown>(opts, "Route", "create", { values })
      if (!results.length) throw new Error("No value returned from Route.create")
      return RouteSchema.parse(results[0])
    },

    async update(id: number, values: Partial<Omit<Route, "id">>): Promise<Route> {
      const results = await apiCall<unknown>(opts, "Route", "update", {
        where: [["id", "=", id]],
        values,
      })
      if (!results.length) throw new Error("No value returned from Route.update")
      return RouteSchema.parse(results[0])
    },

    async delete(id: number): Promise<void> {
      await apiCall<unknown>(opts, "Route", "delete", {
        where: [["id", "=", id]],
      })
    },
  }
}

export function createSavedSearchClient(opts: CiviCRMClientOptions) {
  return {
    async get(params: {
      where?: WhereClause
      select?: (keyof SavedSearch)[]
      orderBy?: OrderByClause
      limit?: number
      offset?: number
    } = {}): Promise<SavedSearch[]> {
      const raw = await apiCall<unknown>(opts, "SavedSearch", "get", params)
      return raw.map((v) => SavedSearchSchema.parse(v))
    },

    async getById(id: number, select?: (keyof SavedSearch)[]): Promise<SavedSearch> {
      const results = await apiCall<unknown>(opts, "SavedSearch", "get", {
        where: [["id", "=", id]],
        ...(select ? { select } : {}),
      })
      if (!results.length) throw new Error(`SavedSearch ${id} not found`)
      return SavedSearchSchema.parse(results[0])
    },

    async create(values: Partial<Omit<SavedSearch, "id">>): Promise<SavedSearch> {
      const results = await apiCall<unknown>(opts, "SavedSearch", "create", { values })
      if (!results.length) throw new Error("No value returned from SavedSearch.create")
      return SavedSearchSchema.parse(results[0])
    },

    async update(id: number, values: Partial<Omit<SavedSearch, "id">>): Promise<SavedSearch> {
      const results = await apiCall<unknown>(opts, "SavedSearch", "update", {
        where: [["id", "=", id]],
        values,
      })
      if (!results.length) throw new Error("No value returned from SavedSearch.update")
      return SavedSearchSchema.parse(results[0])
    },

    async delete(id: number): Promise<void> {
      await apiCall<unknown>(opts, "SavedSearch", "delete", {
        where: [["id", "=", id]],
      })
    },
  }
}

export function createSearchDisplayClient(opts: CiviCRMClientOptions) {
  return {
    async get(params: {
      where?: WhereClause
      select?: (keyof SearchDisplay)[]
      orderBy?: OrderByClause
      limit?: number
      offset?: number
    } = {}): Promise<SearchDisplay[]> {
      const raw = await apiCall<unknown>(opts, "SearchDisplay", "get", params)
      return raw.map((v) => SearchDisplaySchema.parse(v))
    },

    async getById(id: number, select?: (keyof SearchDisplay)[]): Promise<SearchDisplay> {
      const results = await apiCall<unknown>(opts, "SearchDisplay", "get", {
        where: [["id", "=", id]],
        ...(select ? { select } : {}),
      })
      if (!results.length) throw new Error(`SearchDisplay ${id} not found`)
      return SearchDisplaySchema.parse(results[0])
    },

    async create(values: Partial<Omit<SearchDisplay, "id">>): Promise<SearchDisplay> {
      const results = await apiCall<unknown>(opts, "SearchDisplay", "create", { values })
      if (!results.length) throw new Error("No value returned from SearchDisplay.create")
      return SearchDisplaySchema.parse(results[0])
    },

    async update(id: number, values: Partial<Omit<SearchDisplay, "id">>): Promise<SearchDisplay> {
      const results = await apiCall<unknown>(opts, "SearchDisplay", "update", {
        where: [["id", "=", id]],
        values,
      })
      if (!results.length) throw new Error("No value returned from SearchDisplay.update")
      return SearchDisplaySchema.parse(results[0])
    },

    async delete(id: number): Promise<void> {
      await apiCall<unknown>(opts, "SearchDisplay", "delete", {
        where: [["id", "=", id]],
      })
    },
  }
}

export function createSearchParamSetClient(opts: CiviCRMClientOptions) {
  return {
    async get(params: {
      where?: WhereClause
      select?: (keyof SearchParamSet)[]
      orderBy?: OrderByClause
      limit?: number
      offset?: number
    } = {}): Promise<SearchParamSet[]> {
      const raw = await apiCall<unknown>(opts, "SearchParamSet", "get", params)
      return raw.map((v) => SearchParamSetSchema.parse(v))
    },

    async getById(id: number, select?: (keyof SearchParamSet)[]): Promise<SearchParamSet> {
      const results = await apiCall<unknown>(opts, "SearchParamSet", "get", {
        where: [["id", "=", id]],
        ...(select ? { select } : {}),
      })
      if (!results.length) throw new Error(`SearchParamSet ${id} not found`)
      return SearchParamSetSchema.parse(results[0])
    },

    async create(values: Partial<Omit<SearchParamSet, "id">>): Promise<SearchParamSet> {
      const results = await apiCall<unknown>(opts, "SearchParamSet", "create", { values })
      if (!results.length) throw new Error("No value returned from SearchParamSet.create")
      return SearchParamSetSchema.parse(results[0])
    },

    async update(id: number, values: Partial<Omit<SearchParamSet, "id">>): Promise<SearchParamSet> {
      const results = await apiCall<unknown>(opts, "SearchParamSet", "update", {
        where: [["id", "=", id]],
        values,
      })
      if (!results.length) throw new Error("No value returned from SearchParamSet.update")
      return SearchParamSetSchema.parse(results[0])
    },

    async delete(id: number): Promise<void> {
      await apiCall<unknown>(opts, "SearchParamSet", "delete", {
        where: [["id", "=", id]],
      })
    },
  }
}

export function createSearchSegmentClient(opts: CiviCRMClientOptions) {
  return {
    async get(params: {
      where?: WhereClause
      select?: (keyof SearchSegment)[]
      orderBy?: OrderByClause
      limit?: number
      offset?: number
    } = {}): Promise<SearchSegment[]> {
      const raw = await apiCall<unknown>(opts, "SearchSegment", "get", params)
      return raw.map((v) => SearchSegmentSchema.parse(v))
    },

    async getById(id: number, select?: (keyof SearchSegment)[]): Promise<SearchSegment> {
      const results = await apiCall<unknown>(opts, "SearchSegment", "get", {
        where: [["id", "=", id]],
        ...(select ? { select } : {}),
      })
      if (!results.length) throw new Error(`SearchSegment ${id} not found`)
      return SearchSegmentSchema.parse(results[0])
    },

    async create(values: Partial<Omit<SearchSegment, "id">>): Promise<SearchSegment> {
      const results = await apiCall<unknown>(opts, "SearchSegment", "create", { values })
      if (!results.length) throw new Error("No value returned from SearchSegment.create")
      return SearchSegmentSchema.parse(results[0])
    },

    async update(id: number, values: Partial<Omit<SearchSegment, "id">>): Promise<SearchSegment> {
      const results = await apiCall<unknown>(opts, "SearchSegment", "update", {
        where: [["id", "=", id]],
        values,
      })
      if (!results.length) throw new Error("No value returned from SearchSegment.update")
      return SearchSegmentSchema.parse(results[0])
    },

    async delete(id: number): Promise<void> {
      await apiCall<unknown>(opts, "SearchSegment", "delete", {
        where: [["id", "=", id]],
      })
    },
  }
}

export function createSessionClient(opts: CiviCRMClientOptions) {
  return {
    async get(params: {
      where?: WhereClause
      select?: (keyof Session)[]
      orderBy?: OrderByClause
      limit?: number
      offset?: number
    } = {}): Promise<Session[]> {
      const raw = await apiCall<unknown>(opts, "Session", "get", params)
      return raw.map((v) => SessionSchema.parse(v))
    },

    async getById(id: number, select?: (keyof Session)[]): Promise<Session> {
      const results = await apiCall<unknown>(opts, "Session", "get", {
        where: [["id", "=", id]],
        ...(select ? { select } : {}),
      })
      if (!results.length) throw new Error(`Session ${id} not found`)
      return SessionSchema.parse(results[0])
    },

    async create(values: Partial<Omit<Session, "id">>): Promise<Session> {
      const results = await apiCall<unknown>(opts, "Session", "create", { values })
      if (!results.length) throw new Error("No value returned from Session.create")
      return SessionSchema.parse(results[0])
    },

    async update(id: number, values: Partial<Omit<Session, "id">>): Promise<Session> {
      const results = await apiCall<unknown>(opts, "Session", "update", {
        where: [["id", "=", id]],
        values,
      })
      if (!results.length) throw new Error("No value returned from Session.update")
      return SessionSchema.parse(results[0])
    },

    async delete(id: number): Promise<void> {
      await apiCall<unknown>(opts, "Session", "delete", {
        where: [["id", "=", id]],
      })
    },
  }
}

export function createSettingClient(opts: CiviCRMClientOptions) {
  return {
    async get(params: {
      where?: WhereClause
      select?: (keyof Setting)[]
      orderBy?: OrderByClause
      limit?: number
      offset?: number
    } = {}): Promise<Setting[]> {
      const raw = await apiCall<unknown>(opts, "Setting", "get", params)
      return raw.map((v) => SettingSchema.parse(v))
    },

    async getById(id: number, select?: (keyof Setting)[]): Promise<Setting> {
      const results = await apiCall<unknown>(opts, "Setting", "get", {
        where: [["id", "=", id]],
        ...(select ? { select } : {}),
      })
      if (!results.length) throw new Error(`Setting ${id} not found`)
      return SettingSchema.parse(results[0])
    },

    async create(values: Partial<Omit<Setting, "id">>): Promise<Setting> {
      const results = await apiCall<unknown>(opts, "Setting", "create", { values })
      if (!results.length) throw new Error("No value returned from Setting.create")
      return SettingSchema.parse(results[0])
    },

    async update(id: number, values: Partial<Omit<Setting, "id">>): Promise<Setting> {
      const results = await apiCall<unknown>(opts, "Setting", "update", {
        where: [["id", "=", id]],
        values,
      })
      if (!results.length) throw new Error("No value returned from Setting.update")
      return SettingSchema.parse(results[0])
    },

    async delete(id: number): Promise<void> {
      await apiCall<unknown>(opts, "Setting", "delete", {
        where: [["id", "=", id]],
      })
    },
  }
}

export function createSiteEmailAddressClient(opts: CiviCRMClientOptions) {
  return {
    async get(params: {
      where?: WhereClause
      select?: (keyof SiteEmailAddress)[]
      orderBy?: OrderByClause
      limit?: number
      offset?: number
    } = {}): Promise<SiteEmailAddress[]> {
      const raw = await apiCall<unknown>(opts, "SiteEmailAddress", "get", params)
      return raw.map((v) => SiteEmailAddressSchema.parse(v))
    },

    async getById(id: number, select?: (keyof SiteEmailAddress)[]): Promise<SiteEmailAddress> {
      const results = await apiCall<unknown>(opts, "SiteEmailAddress", "get", {
        where: [["id", "=", id]],
        ...(select ? { select } : {}),
      })
      if (!results.length) throw new Error(`SiteEmailAddress ${id} not found`)
      return SiteEmailAddressSchema.parse(results[0])
    },

    async create(values: Partial<Omit<SiteEmailAddress, "id">>): Promise<SiteEmailAddress> {
      const results = await apiCall<unknown>(opts, "SiteEmailAddress", "create", { values })
      if (!results.length) throw new Error("No value returned from SiteEmailAddress.create")
      return SiteEmailAddressSchema.parse(results[0])
    },

    async update(id: number, values: Partial<Omit<SiteEmailAddress, "id">>): Promise<SiteEmailAddress> {
      const results = await apiCall<unknown>(opts, "SiteEmailAddress", "update", {
        where: [["id", "=", id]],
        values,
      })
      if (!results.length) throw new Error("No value returned from SiteEmailAddress.update")
      return SiteEmailAddressSchema.parse(results[0])
    },

    async delete(id: number): Promise<void> {
      await apiCall<unknown>(opts, "SiteEmailAddress", "delete", {
        where: [["id", "=", id]],
      })
    },
  }
}

export function createSiteTokenClient(opts: CiviCRMClientOptions) {
  return {
    async get(params: {
      where?: WhereClause
      select?: (keyof SiteToken)[]
      orderBy?: OrderByClause
      limit?: number
      offset?: number
    } = {}): Promise<SiteToken[]> {
      const raw = await apiCall<unknown>(opts, "SiteToken", "get", params)
      return raw.map((v) => SiteTokenSchema.parse(v))
    },

    async getById(id: number, select?: (keyof SiteToken)[]): Promise<SiteToken> {
      const results = await apiCall<unknown>(opts, "SiteToken", "get", {
        where: [["id", "=", id]],
        ...(select ? { select } : {}),
      })
      if (!results.length) throw new Error(`SiteToken ${id} not found`)
      return SiteTokenSchema.parse(results[0])
    },

    async create(values: Partial<Omit<SiteToken, "id">>): Promise<SiteToken> {
      const results = await apiCall<unknown>(opts, "SiteToken", "create", { values })
      if (!results.length) throw new Error("No value returned from SiteToken.create")
      return SiteTokenSchema.parse(results[0])
    },

    async update(id: number, values: Partial<Omit<SiteToken, "id">>): Promise<SiteToken> {
      const results = await apiCall<unknown>(opts, "SiteToken", "update", {
        where: [["id", "=", id]],
        values,
      })
      if (!results.length) throw new Error("No value returned from SiteToken.update")
      return SiteTokenSchema.parse(results[0])
    },

    async delete(id: number): Promise<void> {
      await apiCall<unknown>(opts, "SiteToken", "delete", {
        where: [["id", "=", id]],
      })
    },
  }
}

export function createSmsProviderClient(opts: CiviCRMClientOptions) {
  return {
    async get(params: {
      where?: WhereClause
      select?: (keyof SmsProvider)[]
      orderBy?: OrderByClause
      limit?: number
      offset?: number
    } = {}): Promise<SmsProvider[]> {
      const raw = await apiCall<unknown>(opts, "SmsProvider", "get", params)
      return raw.map((v) => SmsProviderSchema.parse(v))
    },

    async getById(id: number, select?: (keyof SmsProvider)[]): Promise<SmsProvider> {
      const results = await apiCall<unknown>(opts, "SmsProvider", "get", {
        where: [["id", "=", id]],
        ...(select ? { select } : {}),
      })
      if (!results.length) throw new Error(`SmsProvider ${id} not found`)
      return SmsProviderSchema.parse(results[0])
    },

    async create(values: Partial<Omit<SmsProvider, "id">>): Promise<SmsProvider> {
      const results = await apiCall<unknown>(opts, "SmsProvider", "create", { values })
      if (!results.length) throw new Error("No value returned from SmsProvider.create")
      return SmsProviderSchema.parse(results[0])
    },

    async update(id: number, values: Partial<Omit<SmsProvider, "id">>): Promise<SmsProvider> {
      const results = await apiCall<unknown>(opts, "SmsProvider", "update", {
        where: [["id", "=", id]],
        values,
      })
      if (!results.length) throw new Error("No value returned from SmsProvider.update")
      return SmsProviderSchema.parse(results[0])
    },

    async delete(id: number): Promise<void> {
      await apiCall<unknown>(opts, "SmsProvider", "delete", {
        where: [["id", "=", id]],
      })
    },
  }
}

export function createStateProvinceClient(opts: CiviCRMClientOptions) {
  return {
    async get(params: {
      where?: WhereClause
      select?: (keyof StateProvince)[]
      orderBy?: OrderByClause
      limit?: number
      offset?: number
    } = {}): Promise<StateProvince[]> {
      const raw = await apiCall<unknown>(opts, "StateProvince", "get", params)
      return raw.map((v) => StateProvinceSchema.parse(v))
    },

    async getById(id: number, select?: (keyof StateProvince)[]): Promise<StateProvince> {
      const results = await apiCall<unknown>(opts, "StateProvince", "get", {
        where: [["id", "=", id]],
        ...(select ? { select } : {}),
      })
      if (!results.length) throw new Error(`StateProvince ${id} not found`)
      return StateProvinceSchema.parse(results[0])
    },

    async create(values: Partial<Omit<StateProvince, "id">>): Promise<StateProvince> {
      const results = await apiCall<unknown>(opts, "StateProvince", "create", { values })
      if (!results.length) throw new Error("No value returned from StateProvince.create")
      return StateProvinceSchema.parse(results[0])
    },

    async update(id: number, values: Partial<Omit<StateProvince, "id">>): Promise<StateProvince> {
      const results = await apiCall<unknown>(opts, "StateProvince", "update", {
        where: [["id", "=", id]],
        values,
      })
      if (!results.length) throw new Error("No value returned from StateProvince.update")
      return StateProvinceSchema.parse(results[0])
    },

    async delete(id: number): Promise<void> {
      await apiCall<unknown>(opts, "StateProvince", "delete", {
        where: [["id", "=", id]],
      })
    },
  }
}

export function createStatusPreferenceClient(opts: CiviCRMClientOptions) {
  return {
    async get(params: {
      where?: WhereClause
      select?: (keyof StatusPreference)[]
      orderBy?: OrderByClause
      limit?: number
      offset?: number
    } = {}): Promise<StatusPreference[]> {
      const raw = await apiCall<unknown>(opts, "StatusPreference", "get", params)
      return raw.map((v) => StatusPreferenceSchema.parse(v))
    },

    async getById(id: number, select?: (keyof StatusPreference)[]): Promise<StatusPreference> {
      const results = await apiCall<unknown>(opts, "StatusPreference", "get", {
        where: [["id", "=", id]],
        ...(select ? { select } : {}),
      })
      if (!results.length) throw new Error(`StatusPreference ${id} not found`)
      return StatusPreferenceSchema.parse(results[0])
    },

    async create(values: Partial<Omit<StatusPreference, "id">>): Promise<StatusPreference> {
      const results = await apiCall<unknown>(opts, "StatusPreference", "create", { values })
      if (!results.length) throw new Error("No value returned from StatusPreference.create")
      return StatusPreferenceSchema.parse(results[0])
    },

    async update(id: number, values: Partial<Omit<StatusPreference, "id">>): Promise<StatusPreference> {
      const results = await apiCall<unknown>(opts, "StatusPreference", "update", {
        where: [["id", "=", id]],
        values,
      })
      if (!results.length) throw new Error("No value returned from StatusPreference.update")
      return StatusPreferenceSchema.parse(results[0])
    },

    async delete(id: number): Promise<void> {
      await apiCall<unknown>(opts, "StatusPreference", "delete", {
        where: [["id", "=", id]],
      })
    },
  }
}

export function createSubscriptionHistoryClient(opts: CiviCRMClientOptions) {
  return {
    async get(params: {
      where?: WhereClause
      select?: (keyof SubscriptionHistory)[]
      orderBy?: OrderByClause
      limit?: number
      offset?: number
    } = {}): Promise<SubscriptionHistory[]> {
      const raw = await apiCall<unknown>(opts, "SubscriptionHistory", "get", params)
      return raw.map((v) => SubscriptionHistorySchema.parse(v))
    },

    async getById(id: number, select?: (keyof SubscriptionHistory)[]): Promise<SubscriptionHistory> {
      const results = await apiCall<unknown>(opts, "SubscriptionHistory", "get", {
        where: [["id", "=", id]],
        ...(select ? { select } : {}),
      })
      if (!results.length) throw new Error(`SubscriptionHistory ${id} not found`)
      return SubscriptionHistorySchema.parse(results[0])
    },

    async create(values: Partial<Omit<SubscriptionHistory, "id">>): Promise<SubscriptionHistory> {
      const results = await apiCall<unknown>(opts, "SubscriptionHistory", "create", { values })
      if (!results.length) throw new Error("No value returned from SubscriptionHistory.create")
      return SubscriptionHistorySchema.parse(results[0])
    },

    async update(id: number, values: Partial<Omit<SubscriptionHistory, "id">>): Promise<SubscriptionHistory> {
      const results = await apiCall<unknown>(opts, "SubscriptionHistory", "update", {
        where: [["id", "=", id]],
        values,
      })
      if (!results.length) throw new Error("No value returned from SubscriptionHistory.update")
      return SubscriptionHistorySchema.parse(results[0])
    },

    async delete(id: number): Promise<void> {
      await apiCall<unknown>(opts, "SubscriptionHistory", "delete", {
        where: [["id", "=", id]],
      })
    },
  }
}

export function createTagClient(opts: CiviCRMClientOptions) {
  return {
    async get(params: {
      where?: WhereClause
      select?: (keyof Tag)[]
      orderBy?: OrderByClause
      limit?: number
      offset?: number
    } = {}): Promise<Tag[]> {
      const raw = await apiCall<unknown>(opts, "Tag", "get", params)
      return raw.map((v) => TagSchema.parse(v))
    },

    async getById(id: number, select?: (keyof Tag)[]): Promise<Tag> {
      const results = await apiCall<unknown>(opts, "Tag", "get", {
        where: [["id", "=", id]],
        ...(select ? { select } : {}),
      })
      if (!results.length) throw new Error(`Tag ${id} not found`)
      return TagSchema.parse(results[0])
    },

    async create(values: Partial<Omit<Tag, "id">>): Promise<Tag> {
      const results = await apiCall<unknown>(opts, "Tag", "create", { values })
      if (!results.length) throw new Error("No value returned from Tag.create")
      return TagSchema.parse(results[0])
    },

    async update(id: number, values: Partial<Omit<Tag, "id">>): Promise<Tag> {
      const results = await apiCall<unknown>(opts, "Tag", "update", {
        where: [["id", "=", id]],
        values,
      })
      if (!results.length) throw new Error("No value returned from Tag.update")
      return TagSchema.parse(results[0])
    },

    async delete(id: number): Promise<void> {
      await apiCall<unknown>(opts, "Tag", "delete", {
        where: [["id", "=", id]],
      })
    },
  }
}

export function createTotpClient(opts: CiviCRMClientOptions) {
  return {
    async get(params: {
      where?: WhereClause
      select?: (keyof Totp)[]
      orderBy?: OrderByClause
      limit?: number
      offset?: number
    } = {}): Promise<Totp[]> {
      const raw = await apiCall<unknown>(opts, "Totp", "get", params)
      return raw.map((v) => TotpSchema.parse(v))
    },

    async getById(id: number, select?: (keyof Totp)[]): Promise<Totp> {
      const results = await apiCall<unknown>(opts, "Totp", "get", {
        where: [["id", "=", id]],
        ...(select ? { select } : {}),
      })
      if (!results.length) throw new Error(`Totp ${id} not found`)
      return TotpSchema.parse(results[0])
    },

    async create(values: Partial<Omit<Totp, "id">>): Promise<Totp> {
      const results = await apiCall<unknown>(opts, "Totp", "create", { values })
      if (!results.length) throw new Error("No value returned from Totp.create")
      return TotpSchema.parse(results[0])
    },

    async update(id: number, values: Partial<Omit<Totp, "id">>): Promise<Totp> {
      const results = await apiCall<unknown>(opts, "Totp", "update", {
        where: [["id", "=", id]],
        values,
      })
      if (!results.length) throw new Error("No value returned from Totp.update")
      return TotpSchema.parse(results[0])
    },

    async delete(id: number): Promise<void> {
      await apiCall<unknown>(opts, "Totp", "delete", {
        where: [["id", "=", id]],
      })
    },
  }
}

export function createTranslationClient(opts: CiviCRMClientOptions) {
  return {
    async get(params: {
      where?: WhereClause
      select?: (keyof Translation)[]
      orderBy?: OrderByClause
      limit?: number
      offset?: number
    } = {}): Promise<Translation[]> {
      const raw = await apiCall<unknown>(opts, "Translation", "get", params)
      return raw.map((v) => TranslationSchema.parse(v))
    },

    async getById(id: number, select?: (keyof Translation)[]): Promise<Translation> {
      const results = await apiCall<unknown>(opts, "Translation", "get", {
        where: [["id", "=", id]],
        ...(select ? { select } : {}),
      })
      if (!results.length) throw new Error(`Translation ${id} not found`)
      return TranslationSchema.parse(results[0])
    },

    async create(values: Partial<Omit<Translation, "id">>): Promise<Translation> {
      const results = await apiCall<unknown>(opts, "Translation", "create", { values })
      if (!results.length) throw new Error("No value returned from Translation.create")
      return TranslationSchema.parse(results[0])
    },

    async update(id: number, values: Partial<Omit<Translation, "id">>): Promise<Translation> {
      const results = await apiCall<unknown>(opts, "Translation", "update", {
        where: [["id", "=", id]],
        values,
      })
      if (!results.length) throw new Error("No value returned from Translation.update")
      return TranslationSchema.parse(results[0])
    },

    async delete(id: number): Promise<void> {
      await apiCall<unknown>(opts, "Translation", "delete", {
        where: [["id", "=", id]],
      })
    },
  }
}

export function createTranslationSourceClient(opts: CiviCRMClientOptions) {
  return {
    async get(params: {
      where?: WhereClause
      select?: (keyof TranslationSource)[]
      orderBy?: OrderByClause
      limit?: number
      offset?: number
    } = {}): Promise<TranslationSource[]> {
      const raw = await apiCall<unknown>(opts, "TranslationSource", "get", params)
      return raw.map((v) => TranslationSourceSchema.parse(v))
    },

    async getById(id: number, select?: (keyof TranslationSource)[]): Promise<TranslationSource> {
      const results = await apiCall<unknown>(opts, "TranslationSource", "get", {
        where: [["id", "=", id]],
        ...(select ? { select } : {}),
      })
      if (!results.length) throw new Error(`TranslationSource ${id} not found`)
      return TranslationSourceSchema.parse(results[0])
    },

    async create(values: Partial<Omit<TranslationSource, "id">>): Promise<TranslationSource> {
      const results = await apiCall<unknown>(opts, "TranslationSource", "create", { values })
      if (!results.length) throw new Error("No value returned from TranslationSource.create")
      return TranslationSourceSchema.parse(results[0])
    },

    async update(id: number, values: Partial<Omit<TranslationSource, "id">>): Promise<TranslationSource> {
      const results = await apiCall<unknown>(opts, "TranslationSource", "update", {
        where: [["id", "=", id]],
        values,
      })
      if (!results.length) throw new Error("No value returned from TranslationSource.update")
      return TranslationSourceSchema.parse(results[0])
    },

    async delete(id: number): Promise<void> {
      await apiCall<unknown>(opts, "TranslationSource", "delete", {
        where: [["id", "=", id]],
      })
    },
  }
}

export function createUFFieldClient(opts: CiviCRMClientOptions) {
  return {
    async get(params: {
      where?: WhereClause
      select?: (keyof UFField)[]
      orderBy?: OrderByClause
      limit?: number
      offset?: number
    } = {}): Promise<UFField[]> {
      const raw = await apiCall<unknown>(opts, "UFField", "get", params)
      return raw.map((v) => UFFieldSchema.parse(v))
    },

    async getById(id: number, select?: (keyof UFField)[]): Promise<UFField> {
      const results = await apiCall<unknown>(opts, "UFField", "get", {
        where: [["id", "=", id]],
        ...(select ? { select } : {}),
      })
      if (!results.length) throw new Error(`UFField ${id} not found`)
      return UFFieldSchema.parse(results[0])
    },

    async create(values: Partial<Omit<UFField, "id">>): Promise<UFField> {
      const results = await apiCall<unknown>(opts, "UFField", "create", { values })
      if (!results.length) throw new Error("No value returned from UFField.create")
      return UFFieldSchema.parse(results[0])
    },

    async update(id: number, values: Partial<Omit<UFField, "id">>): Promise<UFField> {
      const results = await apiCall<unknown>(opts, "UFField", "update", {
        where: [["id", "=", id]],
        values,
      })
      if (!results.length) throw new Error("No value returned from UFField.update")
      return UFFieldSchema.parse(results[0])
    },

    async delete(id: number): Promise<void> {
      await apiCall<unknown>(opts, "UFField", "delete", {
        where: [["id", "=", id]],
      })
    },
  }
}

export function createUFGroupClient(opts: CiviCRMClientOptions) {
  return {
    async get(params: {
      where?: WhereClause
      select?: (keyof UFGroup)[]
      orderBy?: OrderByClause
      limit?: number
      offset?: number
    } = {}): Promise<UFGroup[]> {
      const raw = await apiCall<unknown>(opts, "UFGroup", "get", params)
      return raw.map((v) => UFGroupSchema.parse(v))
    },

    async getById(id: number, select?: (keyof UFGroup)[]): Promise<UFGroup> {
      const results = await apiCall<unknown>(opts, "UFGroup", "get", {
        where: [["id", "=", id]],
        ...(select ? { select } : {}),
      })
      if (!results.length) throw new Error(`UFGroup ${id} not found`)
      return UFGroupSchema.parse(results[0])
    },

    async create(values: Partial<Omit<UFGroup, "id">>): Promise<UFGroup> {
      const results = await apiCall<unknown>(opts, "UFGroup", "create", { values })
      if (!results.length) throw new Error("No value returned from UFGroup.create")
      return UFGroupSchema.parse(results[0])
    },

    async update(id: number, values: Partial<Omit<UFGroup, "id">>): Promise<UFGroup> {
      const results = await apiCall<unknown>(opts, "UFGroup", "update", {
        where: [["id", "=", id]],
        values,
      })
      if (!results.length) throw new Error("No value returned from UFGroup.update")
      return UFGroupSchema.parse(results[0])
    },

    async delete(id: number): Promise<void> {
      await apiCall<unknown>(opts, "UFGroup", "delete", {
        where: [["id", "=", id]],
      })
    },
  }
}

export function createUFJoinClient(opts: CiviCRMClientOptions) {
  return {
    async get(params: {
      where?: WhereClause
      select?: (keyof UFJoin)[]
      orderBy?: OrderByClause
      limit?: number
      offset?: number
    } = {}): Promise<UFJoin[]> {
      const raw = await apiCall<unknown>(opts, "UFJoin", "get", params)
      return raw.map((v) => UFJoinSchema.parse(v))
    },

    async getById(id: number, select?: (keyof UFJoin)[]): Promise<UFJoin> {
      const results = await apiCall<unknown>(opts, "UFJoin", "get", {
        where: [["id", "=", id]],
        ...(select ? { select } : {}),
      })
      if (!results.length) throw new Error(`UFJoin ${id} not found`)
      return UFJoinSchema.parse(results[0])
    },

    async create(values: Partial<Omit<UFJoin, "id">>): Promise<UFJoin> {
      const results = await apiCall<unknown>(opts, "UFJoin", "create", { values })
      if (!results.length) throw new Error("No value returned from UFJoin.create")
      return UFJoinSchema.parse(results[0])
    },

    async update(id: number, values: Partial<Omit<UFJoin, "id">>): Promise<UFJoin> {
      const results = await apiCall<unknown>(opts, "UFJoin", "update", {
        where: [["id", "=", id]],
        values,
      })
      if (!results.length) throw new Error("No value returned from UFJoin.update")
      return UFJoinSchema.parse(results[0])
    },

    async delete(id: number): Promise<void> {
      await apiCall<unknown>(opts, "UFJoin", "delete", {
        where: [["id", "=", id]],
      })
    },
  }
}

export function createUFMatchClient(opts: CiviCRMClientOptions) {
  return {
    async get(params: {
      where?: WhereClause
      select?: (keyof UFMatch)[]
      orderBy?: OrderByClause
      limit?: number
      offset?: number
    } = {}): Promise<UFMatch[]> {
      const raw = await apiCall<unknown>(opts, "UFMatch", "get", params)
      return raw.map((v) => UFMatchSchema.parse(v))
    },

    async getById(id: number, select?: (keyof UFMatch)[]): Promise<UFMatch> {
      const results = await apiCall<unknown>(opts, "UFMatch", "get", {
        where: [["id", "=", id]],
        ...(select ? { select } : {}),
      })
      if (!results.length) throw new Error(`UFMatch ${id} not found`)
      return UFMatchSchema.parse(results[0])
    },

    async create(values: Partial<Omit<UFMatch, "id">>): Promise<UFMatch> {
      const results = await apiCall<unknown>(opts, "UFMatch", "create", { values })
      if (!results.length) throw new Error("No value returned from UFMatch.create")
      return UFMatchSchema.parse(results[0])
    },

    async update(id: number, values: Partial<Omit<UFMatch, "id">>): Promise<UFMatch> {
      const results = await apiCall<unknown>(opts, "UFMatch", "update", {
        where: [["id", "=", id]],
        values,
      })
      if (!results.length) throw new Error("No value returned from UFMatch.update")
      return UFMatchSchema.parse(results[0])
    },

    async delete(id: number): Promise<void> {
      await apiCall<unknown>(opts, "UFMatch", "delete", {
        where: [["id", "=", id]],
      })
    },
  }
}

export function createUserClient(opts: CiviCRMClientOptions) {
  return {
    async get(params: {
      where?: WhereClause
      select?: (keyof User)[]
      orderBy?: OrderByClause
      limit?: number
      offset?: number
    } = {}): Promise<User[]> {
      const raw = await apiCall<unknown>(opts, "User", "get", params)
      return raw.map((v) => UserSchema.parse(v))
    },

    async getById(id: number, select?: (keyof User)[]): Promise<User> {
      const results = await apiCall<unknown>(opts, "User", "get", {
        where: [["id", "=", id]],
        ...(select ? { select } : {}),
      })
      if (!results.length) throw new Error(`User ${id} not found`)
      return UserSchema.parse(results[0])
    },

    async create(values: Partial<Omit<User, "id">>): Promise<User> {
      const results = await apiCall<unknown>(opts, "User", "create", { values })
      if (!results.length) throw new Error("No value returned from User.create")
      return UserSchema.parse(results[0])
    },

    async update(id: number, values: Partial<Omit<User, "id">>): Promise<User> {
      const results = await apiCall<unknown>(opts, "User", "update", {
        where: [["id", "=", id]],
        values,
      })
      if (!results.length) throw new Error("No value returned from User.update")
      return UserSchema.parse(results[0])
    },

    async delete(id: number): Promise<void> {
      await apiCall<unknown>(opts, "User", "delete", {
        where: [["id", "=", id]],
      })
    },
  }
}

export function createUserJobClient(opts: CiviCRMClientOptions) {
  return {
    async get(params: {
      where?: WhereClause
      select?: (keyof UserJob)[]
      orderBy?: OrderByClause
      limit?: number
      offset?: number
    } = {}): Promise<UserJob[]> {
      const raw = await apiCall<unknown>(opts, "UserJob", "get", params)
      return raw.map((v) => UserJobSchema.parse(v))
    },

    async getById(id: number, select?: (keyof UserJob)[]): Promise<UserJob> {
      const results = await apiCall<unknown>(opts, "UserJob", "get", {
        where: [["id", "=", id]],
        ...(select ? { select } : {}),
      })
      if (!results.length) throw new Error(`UserJob ${id} not found`)
      return UserJobSchema.parse(results[0])
    },

    async create(values: Partial<Omit<UserJob, "id">>): Promise<UserJob> {
      const results = await apiCall<unknown>(opts, "UserJob", "create", { values })
      if (!results.length) throw new Error("No value returned from UserJob.create")
      return UserJobSchema.parse(results[0])
    },

    async update(id: number, values: Partial<Omit<UserJob, "id">>): Promise<UserJob> {
      const results = await apiCall<unknown>(opts, "UserJob", "update", {
        where: [["id", "=", id]],
        values,
      })
      if (!results.length) throw new Error("No value returned from UserJob.update")
      return UserJobSchema.parse(results[0])
    },

    async delete(id: number): Promise<void> {
      await apiCall<unknown>(opts, "UserJob", "delete", {
        where: [["id", "=", id]],
      })
    },
  }
}

export function createUserRoleClient(opts: CiviCRMClientOptions) {
  return {
    async get(params: {
      where?: WhereClause
      select?: (keyof UserRole)[]
      orderBy?: OrderByClause
      limit?: number
      offset?: number
    } = {}): Promise<UserRole[]> {
      const raw = await apiCall<unknown>(opts, "UserRole", "get", params)
      return raw.map((v) => UserRoleSchema.parse(v))
    },

    async getById(id: number, select?: (keyof UserRole)[]): Promise<UserRole> {
      const results = await apiCall<unknown>(opts, "UserRole", "get", {
        where: [["id", "=", id]],
        ...(select ? { select } : {}),
      })
      if (!results.length) throw new Error(`UserRole ${id} not found`)
      return UserRoleSchema.parse(results[0])
    },

    async create(values: Partial<Omit<UserRole, "id">>): Promise<UserRole> {
      const results = await apiCall<unknown>(opts, "UserRole", "create", { values })
      if (!results.length) throw new Error("No value returned from UserRole.create")
      return UserRoleSchema.parse(results[0])
    },

    async update(id: number, values: Partial<Omit<UserRole, "id">>): Promise<UserRole> {
      const results = await apiCall<unknown>(opts, "UserRole", "update", {
        where: [["id", "=", id]],
        values,
      })
      if (!results.length) throw new Error("No value returned from UserRole.update")
      return UserRoleSchema.parse(results[0])
    },

    async delete(id: number): Promise<void> {
      await apiCall<unknown>(opts, "UserRole", "delete", {
        where: [["id", "=", id]],
      })
    },
  }
}

export function createWebsiteClient(opts: CiviCRMClientOptions) {
  return {
    async get(params: {
      where?: WhereClause
      select?: (keyof Website)[]
      orderBy?: OrderByClause
      limit?: number
      offset?: number
    } = {}): Promise<Website[]> {
      const raw = await apiCall<unknown>(opts, "Website", "get", params)
      return raw.map((v) => WebsiteSchema.parse(v))
    },

    async getById(id: number, select?: (keyof Website)[]): Promise<Website> {
      const results = await apiCall<unknown>(opts, "Website", "get", {
        where: [["id", "=", id]],
        ...(select ? { select } : {}),
      })
      if (!results.length) throw new Error(`Website ${id} not found`)
      return WebsiteSchema.parse(results[0])
    },

    async create(values: Partial<Omit<Website, "id">>): Promise<Website> {
      const results = await apiCall<unknown>(opts, "Website", "create", { values })
      if (!results.length) throw new Error("No value returned from Website.create")
      return WebsiteSchema.parse(results[0])
    },

    async update(id: number, values: Partial<Omit<Website, "id">>): Promise<Website> {
      const results = await apiCall<unknown>(opts, "Website", "update", {
        where: [["id", "=", id]],
        values,
      })
      if (!results.length) throw new Error("No value returned from Website.update")
      return WebsiteSchema.parse(results[0])
    },

    async delete(id: number): Promise<void> {
      await apiCall<unknown>(opts, "Website", "delete", {
        where: [["id", "=", id]],
      })
    },
  }
}

export function createWordReplacementClient(opts: CiviCRMClientOptions) {
  return {
    async get(params: {
      where?: WhereClause
      select?: (keyof WordReplacement)[]
      orderBy?: OrderByClause
      limit?: number
      offset?: number
    } = {}): Promise<WordReplacement[]> {
      const raw = await apiCall<unknown>(opts, "WordReplacement", "get", params)
      return raw.map((v) => WordReplacementSchema.parse(v))
    },

    async getById(id: number, select?: (keyof WordReplacement)[]): Promise<WordReplacement> {
      const results = await apiCall<unknown>(opts, "WordReplacement", "get", {
        where: [["id", "=", id]],
        ...(select ? { select } : {}),
      })
      if (!results.length) throw new Error(`WordReplacement ${id} not found`)
      return WordReplacementSchema.parse(results[0])
    },

    async create(values: Partial<Omit<WordReplacement, "id">>): Promise<WordReplacement> {
      const results = await apiCall<unknown>(opts, "WordReplacement", "create", { values })
      if (!results.length) throw new Error("No value returned from WordReplacement.create")
      return WordReplacementSchema.parse(results[0])
    },

    async update(id: number, values: Partial<Omit<WordReplacement, "id">>): Promise<WordReplacement> {
      const results = await apiCall<unknown>(opts, "WordReplacement", "update", {
        where: [["id", "=", id]],
        values,
      })
      if (!results.length) throw new Error("No value returned from WordReplacement.update")
      return WordReplacementSchema.parse(results[0])
    },

    async delete(id: number): Promise<void> {
      await apiCall<unknown>(opts, "WordReplacement", "delete", {
        where: [["id", "=", id]],
      })
    },
  }
}

export function createWorkflowMessageClient(opts: CiviCRMClientOptions) {
  return {
    async get(params: {
      where?: WhereClause
      select?: (keyof WorkflowMessage)[]
      orderBy?: OrderByClause
      limit?: number
      offset?: number
    } = {}): Promise<WorkflowMessage[]> {
      const raw = await apiCall<unknown>(opts, "WorkflowMessage", "get", params)
      return raw.map((v) => WorkflowMessageSchema.parse(v))
    },

    async getById(id: number, select?: (keyof WorkflowMessage)[]): Promise<WorkflowMessage> {
      const results = await apiCall<unknown>(opts, "WorkflowMessage", "get", {
        where: [["id", "=", id]],
        ...(select ? { select } : {}),
      })
      if (!results.length) throw new Error(`WorkflowMessage ${id} not found`)
      return WorkflowMessageSchema.parse(results[0])
    },

    async create(values: Partial<Omit<WorkflowMessage, "id">>): Promise<WorkflowMessage> {
      const results = await apiCall<unknown>(opts, "WorkflowMessage", "create", { values })
      if (!results.length) throw new Error("No value returned from WorkflowMessage.create")
      return WorkflowMessageSchema.parse(results[0])
    },

    async update(id: number, values: Partial<Omit<WorkflowMessage, "id">>): Promise<WorkflowMessage> {
      const results = await apiCall<unknown>(opts, "WorkflowMessage", "update", {
        where: [["id", "=", id]],
        values,
      })
      if (!results.length) throw new Error("No value returned from WorkflowMessage.update")
      return WorkflowMessageSchema.parse(results[0])
    },

    async delete(id: number): Promise<void> {
      await apiCall<unknown>(opts, "WorkflowMessage", "delete", {
        where: [["id", "=", id]],
      })
    },
  }
}

export function createWorldRegionClient(opts: CiviCRMClientOptions) {
  return {
    async get(params: {
      where?: WhereClause
      select?: (keyof WorldRegion)[]
      orderBy?: OrderByClause
      limit?: number
      offset?: number
    } = {}): Promise<WorldRegion[]> {
      const raw = await apiCall<unknown>(opts, "WorldRegion", "get", params)
      return raw.map((v) => WorldRegionSchema.parse(v))
    },

    async getById(id: number, select?: (keyof WorldRegion)[]): Promise<WorldRegion> {
      const results = await apiCall<unknown>(opts, "WorldRegion", "get", {
        where: [["id", "=", id]],
        ...(select ? { select } : {}),
      })
      if (!results.length) throw new Error(`WorldRegion ${id} not found`)
      return WorldRegionSchema.parse(results[0])
    },

    async create(values: Partial<Omit<WorldRegion, "id">>): Promise<WorldRegion> {
      const results = await apiCall<unknown>(opts, "WorldRegion", "create", { values })
      if (!results.length) throw new Error("No value returned from WorldRegion.create")
      return WorldRegionSchema.parse(results[0])
    },

    async update(id: number, values: Partial<Omit<WorldRegion, "id">>): Promise<WorldRegion> {
      const results = await apiCall<unknown>(opts, "WorldRegion", "update", {
        where: [["id", "=", id]],
        values,
      })
      if (!results.length) throw new Error("No value returned from WorldRegion.update")
      return WorldRegionSchema.parse(results[0])
    },

    async delete(id: number): Promise<void> {
      await apiCall<unknown>(opts, "WorldRegion", "delete", {
        where: [["id", "=", id]],
      })
    },
  }
}

export function createCiviCRMClient(opts: CiviCRMClientOptions) {
  return {
    ACL: createACLClient(opts),
    ACLEntityRole: createACLEntityRoleClient(opts),
    ActionSchedule: createActionScheduleClient(opts),
    Activity: createActivityClient(opts),
    ActivityContact: createActivityContactClient(opts),
    Address: createAddressClient(opts),
    Afform: createAfformClient(opts),
    AfformBehavior: createAfformBehaviorClient(opts),
    AfformSubmission: createAfformSubmissionClient(opts),
    Batch: createBatchClient(opts),
    BouncePattern: createBouncePatternClient(opts),
    BounceType: createBounceTypeClient(opts),
    Case: createCaseClient(opts),
    CaseActivity: createCaseActivityClient(opts),
    CaseContact: createCaseContactClient(opts),
    CaseType: createCaseTypeClient(opts),
    Contact: createContactClient(opts),
    ContactType: createContactTypeClient(opts),
    Contribution: createContributionClient(opts),
    ContributionPage: createContributionPageClient(opts),
    ContributionProduct: createContributionProductClient(opts),
    ContributionRecur: createContributionRecurClient(opts),
    ContributionSoft: createContributionSoftClient(opts),
    Country: createCountryClient(opts),
    County: createCountyClient(opts),
    CustomField: createCustomFieldClient(opts),
    CustomGroup: createCustomGroupClient(opts),
    Dashboard: createDashboardClient(opts),
    DashboardContact: createDashboardContactClient(opts),
    DedupeException: createDedupeExceptionClient(opts),
    DedupeRule: createDedupeRuleClient(opts),
    DedupeRuleGroup: createDedupeRuleGroupClient(opts),
    Discount: createDiscountClient(opts),
    Domain: createDomainClient(opts),
    Email: createEmailClient(opts),
    Entity: createEntityClient(opts),
    EntityBatch: createEntityBatchClient(opts),
    EntityFile: createEntityFileClient(opts),
    EntityFinancialAccount: createEntityFinancialAccountClient(opts),
    EntityFinancialTrxn: createEntityFinancialTrxnClient(opts),
    EntityTag: createEntityTagClient(opts),
    Event: createEventClient(opts),
    ExampleData: createExampleDataClient(opts),
    Extension: createExtensionClient(opts),
    File: createFileClient(opts),
    FinancialAccount: createFinancialAccountClient(opts),
    FinancialItem: createFinancialItemClient(opts),
    FinancialTrxn: createFinancialTrxnClient(opts),
    FinancialType: createFinancialTypeClient(opts),
    Group: createGroupClient(opts),
    GroupContact: createGroupContactClient(opts),
    GroupNesting: createGroupNestingClient(opts),
    GroupOrganization: createGroupOrganizationClient(opts),
    GroupSubscription: createGroupSubscriptionClient(opts),
    Household: createHouseholdClient(opts),
    IM: createIMClient(opts),
    Individual: createIndividualClient(opts),
    Job: createJobClient(opts),
    JobLog: createJobLogClient(opts),
    LineItem: createLineItemClient(opts),
    LocBlock: createLocBlockClient(opts),
    LocationType: createLocationTypeClient(opts),
    Log: createLogClient(opts),
    MailSettings: createMailSettingsClient(opts),
    Mailing: createMailingClient(opts),
    MailingComponent: createMailingComponentClient(opts),
    MailingEventBounce: createMailingEventBounceClient(opts),
    MailingEventConfirm: createMailingEventConfirmClient(opts),
    MailingEventDelivered: createMailingEventDeliveredClient(opts),
    MailingEventOpened: createMailingEventOpenedClient(opts),
    MailingEventQueue: createMailingEventQueueClient(opts),
    MailingEventReply: createMailingEventReplyClient(opts),
    MailingEventSubscribe: createMailingEventSubscribeClient(opts),
    MailingEventTrackableURLOpen: createMailingEventTrackableURLOpenClient(opts),
    MailingEventUnsubscribe: createMailingEventUnsubscribeClient(opts),
    MailingGroup: createMailingGroupClient(opts),
    MailingJob: createMailingJobClient(opts),
    MailingTrackableURL: createMailingTrackableURLClient(opts),
    Managed: createManagedClient(opts),
    Mapping: createMappingClient(opts),
    MappingField: createMappingFieldClient(opts),
    Membership: createMembershipClient(opts),
    MembershipBlock: createMembershipBlockClient(opts),
    MembershipLog: createMembershipLogClient(opts),
    MembershipStatus: createMembershipStatusClient(opts),
    MembershipType: createMembershipTypeClient(opts),
    MessageTemplate: createMessageTemplateClient(opts),
    MosaicoTemplate: createMosaicoTemplateClient(opts),
    Navigation: createNavigationClient(opts),
    Note: createNoteClient(opts),
    OpenID: createOpenIDClient(opts),
    OptionGroup: createOptionGroupClient(opts),
    OptionValue: createOptionValueClient(opts),
    Organization: createOrganizationClient(opts),
    PCP: createPCPClient(opts),
    PCPBlock: createPCPBlockClient(opts),
    Participant: createParticipantClient(opts),
    ParticipantStatusType: createParticipantStatusTypeClient(opts),
    Payment: createPaymentClient(opts),
    PaymentProcessor: createPaymentProcessorClient(opts),
    PaymentProcessorType: createPaymentProcessorTypeClient(opts),
    PaymentToken: createPaymentTokenClient(opts),
    Permission: createPermissionClient(opts),
    Phone: createPhoneClient(opts),
    PreferencesDate: createPreferencesDateClient(opts),
    Premium: createPremiumClient(opts),
    PremiumsProduct: createPremiumsProductClient(opts),
    PriceField: createPriceFieldClient(opts),
    PriceFieldValue: createPriceFieldValueClient(opts),
    PriceSet: createPriceSetClient(opts),
    PriceSetEntity: createPriceSetEntityClient(opts),
    PrintLabel: createPrintLabelClient(opts),
    Product: createProductClient(opts),
    Queue: createQueueClient(opts),
    QueueItem: createQueueItemClient(opts),
    RecentItem: createRecentItemClient(opts),
    Relationship: createRelationshipClient(opts),
    RelationshipCache: createRelationshipCacheClient(opts),
    RelationshipType: createRelationshipTypeClient(opts),
    ReportInstance: createReportInstanceClient(opts),
    RiverleaStream: createRiverleaStreamClient(opts),
    Role: createRoleClient(opts),
    RolePermission: createRolePermissionClient(opts),
    Route: createRouteClient(opts),
    SavedSearch: createSavedSearchClient(opts),
    SearchDisplay: createSearchDisplayClient(opts),
    SearchParamSet: createSearchParamSetClient(opts),
    SearchSegment: createSearchSegmentClient(opts),
    Session: createSessionClient(opts),
    Setting: createSettingClient(opts),
    SiteEmailAddress: createSiteEmailAddressClient(opts),
    SiteToken: createSiteTokenClient(opts),
    SmsProvider: createSmsProviderClient(opts),
    StateProvince: createStateProvinceClient(opts),
    StatusPreference: createStatusPreferenceClient(opts),
    SubscriptionHistory: createSubscriptionHistoryClient(opts),
    Tag: createTagClient(opts),
    Totp: createTotpClient(opts),
    Translation: createTranslationClient(opts),
    TranslationSource: createTranslationSourceClient(opts),
    UFField: createUFFieldClient(opts),
    UFGroup: createUFGroupClient(opts),
    UFJoin: createUFJoinClient(opts),
    UFMatch: createUFMatchClient(opts),
    User: createUserClient(opts),
    UserJob: createUserJobClient(opts),
    UserRole: createUserRoleClient(opts),
    Website: createWebsiteClient(opts),
    WordReplacement: createWordReplacementClient(opts),
    WorkflowMessage: createWorkflowMessageClient(opts),
    WorldRegion: createWorldRegionClient(opts),
  }
}