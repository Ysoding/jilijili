INSERT INTO
  users (
    user_id,
    name,
    email,
    roles,
    password,
    enabled,
    date_created,
    date_updated
  )
VALUES
  (
    '5cf37266-3473-4006-984f-9325122678b7',
    '西门吹雪',
    'xmchx@test.com',
    '{ADMIN}',
    '$2a$10$H2DpOWjcFuIctX6DA/y/uueCOntRnx6Ec2DVj76EzT6JTwvhDuXz.',
    true,
    '2024-03-24 00:00:00',
    '2019-03-24 00:00:00'
  ),
  (
    '45b5fbd3-755f-4379-8f07-a58d4a30fa2f',
    'User Gopher',
    'user@example.com',
    '{USER}',
    '$2a$10$H2DpOWjcFuIctX6DA/y/uueCOntRnx6Ec2DVj76EzT6JTwvhDuXz.',
    true,
    '2019-03-24 00:00:00',
    '2019-03-24 00:00:00'
  ) ON CONFLICT DO NOTHING;