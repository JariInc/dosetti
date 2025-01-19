DELETE FROM tenant;

DELETE FROM prescription;

DELETE FROM serving;

INSERT INTO
    tenant (id, key)
VALUES
    (1, 'foobar');

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
        'daily',
        '2024-11-29T08:00:00.000000000Z',
        'Panadol',
        '500 mg'
    );
