select name
from sales_person
where sales_id not in (
    select o.sales_id
    from orders o
    join company c using(com_id)
    where c.name = 'CRIMSON'
)