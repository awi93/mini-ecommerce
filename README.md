# Answer

##  Notes

Please run this script to run database and migrate the table

```
docker-compose up -d
```

1. For Number 1, all the DDL (Data Definition Language) Located at `migrations/ddl`
2. For Number 2, all the entity model located at `src/models/entities`
3. All CSV file located at `csv`. To Import Data from CSV please use command line script bellow

    To import customer data
    ```shell
    go run main.go csv customer -f ./csv/customer.csv
    ```
    To import employee data
    ```shell
    go run main.go csv employee -f ./csv/employee.csv
    ```
    To import product data
    ```shell
    go run main.go csv product -f ./csv/product.csv
    ```
    To import shipping method data
    ```shell
    go run main.go csv shipping-method -f ./csv/shipping_method.csv
    ```
    To import order data
    ```shell
    go run main.go csv order -f ./csv/order.csv
    ```
    To import order detail data
    ```shell
    go run main.go csv orde-detail -f ./csv/order_detail.csv
    ```

4. Here SQL Script for Question No. 4
    * List of customers located in Irvine city.
        ```SQL
            SELECT
                id,
                first_name,
                last_name,
                email,
                phone_number,
                fax_number,
                billing_address,
                billing_city,
                billing_state,
                billing_zip_code,
                company_name,
                company_website,
                ship_address,
                ship_city,
                ship_state,
                ship_zip_code,
                ship_phone_number
            FROM
                public.customers
            WHERE
                billing_city = 'Irvine city'
        ```
    * List of customers whose order is handled by an employee named Adam Barr.
        ```SQL
            SELECT
                DISTINCT 
                c.id,
                c.first_name,
                c.last_name,
                c.email,
                c.phone_number,
                c.fax_number,
                c.billing_address,
                c.billing_city,
                c.billing_state,
                c.billing_zip_code,
                c.company_name,
                c.company_website,
                c.ship_address,
                c.ship_city,
                c.ship_state,
                c.ship_zip_code,
                c.ship_phone_number
            FROM
                customers c
                INNER JOIN
                orders o 
                ON c.id = o.customer_id
                INNER JOIN
                employees e 
                ON e.id = o.employee_id
            WHERE
                e.first_name = 'Adam' AND e.last_name = 'Barr'
        ```
    * List of products which are ordered by "Contonso, Ltd" Company.
        ```SQL
            SELECT
                DISTINCT
                p.id,
                p.product_name,
                p.unit_price,
                p.in_stock 
            FROM
                products p
                INNER JOIN
                order_details od 
                ON p.id = od.product_id
                INNER JOIN
                orders o
                ON o.id = od.order_id
                INNER JOIN
                customers c 
                ON c.id = o.customer_id
            WHERE
                c.company_name = 'Contonso, Ltd'
        ```
    * List of transactions (orders) which has "UPS Ground" as shipping method.
        ```SQL
            SELECT
                o.id,
                o.customer_id,
                o.employee_id,
                o.order_date,
                o.purchase_order_number,
                o.ship_date,
                o.shipping_method_id,
                o.freight_charge,
                o.taxes,
                o.payment_received,
                o."comment" 
            FROM
                orders o
                INNER JOIN
                shipping_methods sm
                ON o.shipping_method_id = sm.id
            WHERE
                sm.shipping_method  = 'UPS Ground'
        ```
    * List of total cost (including tax and freight charge) for every order sorted by ship date.
        ```SQL
            select
                o.id,
                o.customer_id,
                o.employee_id,
                o.order_date,
                o.purchase_order_number,
                o.ship_date,
                o.shipping_method_id,
                o.freight_charge,
                o.taxes,
                od.subtotal,
                od.subtotal + o.freight_charge + o.taxes as grand_total,
                o.payment_received,
                o."comment" 
            from
                orders o
                inner join
                (
                    select
                        od.order_id,
                        sum(od.quantity * (od.unit_price * ((100-od.discount)/100))) as subtotal
                    from
                        order_details od
                    group by
                        od.order_id
                ) as od
                on
                o.id = od.order_id
            order by
                o.ship_date asc
        ```
5. Create an API (REST) using Golang or Java to display order details with the following

    Please run this script to start http service
    ```shell
        go run main.go http
    ```

    Change Host and ID 

    ``` ssh
    curl --location --request GET '{host}:{port}/order-details/{id}'
    ```

    Find port setting at config.yml
    