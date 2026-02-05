# Create a tag for categorizing contacts
resource "civicrm_tag" "priority_donor" {
  name        = "priority_donor"
  label       = "Priority Donor"
  description = "High-value donors requiring special attention"
  color       = "#ffc107"
  used_for    = ["civicrm_contact"]
}
