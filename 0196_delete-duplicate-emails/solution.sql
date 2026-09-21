with duplicates as (
    select
        Id,
        Email,
        row_number() over (partition by email order by Id) rn
    from Person
)
delete
from Person
where Id not in (
    select Id
    from duplicates
    where rn = 1
)
and Email in (select Email from duplicates)
;
