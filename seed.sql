DELETE FROM tenant;

DELETE FROM prescription;

DELETE FROM serving;

INSERT INTO
    tenant (id, key)
VALUES
    (1, '1Ftr9osUPs0');

INSERT INTO
    medicine (tenant, name, doses_left)
VALUES
    (1, 'Panadol 500mg', 24);

INSERT INTO
    prescription (
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
        'daily',
        '2024-11-29T08:00:00.000000000Z',
        last_insert_rowid (),
        1.0
    );
