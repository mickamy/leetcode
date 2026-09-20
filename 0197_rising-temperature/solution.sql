select today.id
from Weather today
join Weather yesterday on today.recordDate = date_add(yesterday.recordDate, interval 1 day)
where today.temperature > yesterday.temperature
;
