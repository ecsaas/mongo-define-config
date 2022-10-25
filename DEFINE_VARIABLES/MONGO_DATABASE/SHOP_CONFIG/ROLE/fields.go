package ROLE

import "github.com/ecsaas/mongo-define-config/DEFINE_VARIABLES/MONGO_DATABASE/DEFAULT_KEY"

const ID = DEFAULT_KEY.ID
const CREATED = DEFAULT_KEY.CREATED
const CREATED_UNIX = DEFAULT_KEY.CREATED_UNIX
const UPDATED = DEFAULT_KEY.UPDATED
const UPDATED_UNIX = DEFAULT_KEY.UPDATED_UNIX
const ROOT_USER_ID = DEFAULT_KEY.ROOT_USER_ID

const KEY_VERIFY_EMAIL = "keyVerifyEmail"

const OWNER = "owner"
const ROLE_EMAIL = "roleEmail"
const ROLE_EMAIL_VERIFIED = "roleEmailVerified"

const GENERAL_ROLES = "generalRoles"
const BILLING_ROLES = "billingRoles"
const MARKETING_AND_CUSTOMER_MANAGEMENT_ROLES = "marketingAndCustomerManagementRoles"
const PAYMENTS_ROLES = "paymentsRoles"
const GIVE_SITE_ACCESS = "giveSiteAccess"
