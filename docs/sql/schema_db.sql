-- Card table
CREATE TABLE card (
    "id" UUID PRIMARY KEY NOT NULL,
    "token" VARCHAR(255) NOT NULL,
    "number" VARCHAR(255) NOT NULL,
    "cvv" VARCHAR(10) NOT NULL
);

-- Card table sample data
INSERT INTO card (id, token, number, cvv) VALUES
    (gen_random_uuid(), '7zIzcfXQJId44/NCv1o8hQ==', '4111111111111111', '123'),
    (gen_random_uuid(), 'NBlwenkmWUag2JWUV1y+ug==', '5555555555554444', '456'),
    (gen_random_uuid(), 'IxrFKczK5Iemr0kop6lMCA==', '378282246310005', '789'),
    (gen_random_uuid(), 'O4YVgfzPZgDdwajI9vh6rQ==', '6011111111111117', '012'),
    (gen_random_uuid(), 'YQHRqHNRnwAZw+xaFH+mjw==', '4222222222222', '345');
