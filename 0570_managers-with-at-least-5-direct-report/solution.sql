select boss.name
from Employee boss
join Employee report on report.managerId = boss.id
group by boss.id
having count(report.id) >= 5
;
