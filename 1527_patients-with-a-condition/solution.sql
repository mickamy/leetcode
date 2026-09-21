select patient_id, patient_name, conditions
from Patients
where regexp_like(conditions, '(^|[[:space:]])DIAB1');
;
