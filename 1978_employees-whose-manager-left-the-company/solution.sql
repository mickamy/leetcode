select e.employee_id
from Employees e
where e.manager_id not in (select employee_id from Employees)
and e.salary < 30000
order by e.employee_id
;