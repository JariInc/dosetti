DELETE FROM tenant;

DELETE FROM prescription;

DELETE FROM serving;

INSERT INTO
    tenant (id, key, name)
VALUES
    (1, "foobar", "FooBar");

INSERT INTO
    prescription (
        id,
        tenant,
        interval,
        interval_unit,
        start_at,
        medicine,
        amount
    )
VALUES
    (
        1,
        1,
        1,
        "daily",
        "2024-11-29 08:00:00.000000000Z",
        "Panadol",
        "500 mg"
    );
