select
    boss.employee_id,
    boss.name,
    count(reporter.employee_id) as reports_count,
    round(avg(reporter.age), 0) as average_age
from Employees boss
join (
    select *
    from Employees e
    where e.reports_to is not null
) reporter on boss.employee_id = reporter.reports_to
group by boss.employee_id, boss.name
order by boss.employee_id
;
