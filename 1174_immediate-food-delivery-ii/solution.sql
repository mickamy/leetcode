SELECT
    ROUND(AVG(fo.order_date = fo.customer_pref_delivery_date) * 100, 2) AS immediate_percentage
FROM (
         SELECT
             order_date,
             customer_pref_delivery_date,
             ROW_NUMBER() OVER (PARTITION BY customer_id ORDER BY order_date) AS rn
         FROM Delivery
     ) AS fo
WHERE fo.rn = 1;