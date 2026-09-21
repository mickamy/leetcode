select
    e.employee_id, e.department_id
from Employee e
where e.primary_flag = 'Y'
union (
    select e.employee_id, min(e.department_id)
    from Employee e
    group by e.employee_id
    having count(e.employee_id) = 1
)
order by employee_id
;
