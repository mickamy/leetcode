select
    earner.Department,
    earner.Employee,
    earner.Salary
from (
    select
        d.name Department,
        e.name Employee,
        e.salary Salary,
        dense_rank() over (partition by d.name order by e.salary desc) rn
    from Employee e
    join Department d on e.departmentId = d.id
) earner
where earner.rn < 4
;
