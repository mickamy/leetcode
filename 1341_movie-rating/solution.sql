select results
from (
    select
        u.name results
    from MovieRating r
    join Users u on r.user_id = u.user_id
    group by 1
    order by count(r.movie_id) desc, 1
    limit 1
) users
union all (
    select m.title results
    from Movies m
    join MovieRating r on m.movie_id = r.movie_id
    where r.created_at >= '2020-02-01' and r.created_at < '2020-03-01'
    group by m.title
    order by avg(r.rating) desc, m.title
    limit 1
)
;
