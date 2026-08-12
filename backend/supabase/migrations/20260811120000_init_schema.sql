-- ============================================================================
-- Doctor Platform — initial schema
-- Generated from the Go/GORM models in internal/*/model.go so that this SQL
-- produces (a superset of) exactly what gorm.AutoMigrate() would create.
-- Safe to run against a fresh Supabase project's SQL editor, or via:
--   supabase db push
-- All statements are idempotent (IF NOT EXISTS) so this can be re-run safely.
-- ============================================================================

-- ----------------------------------------------------------------------------
-- users  (internal/auth.User — this is the table actually used by the app)
-- ----------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS users (
    id                  BIGSERIAL PRIMARY KEY,
    full_name           TEXT,
    email               VARCHAR(255) UNIQUE,
    password            TEXT NOT NULL,
    role                TEXT,
    verification_status TEXT,
    created_at          TIMESTAMPTZ NOT NULL DEFAULT now()
);
CREATE UNIQUE INDEX IF NOT EXISTS idx_users_email ON users (email);

-- ----------------------------------------------------------------------------
-- profiles  (internal/profile.Profile)
-- ----------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS profiles (
    id                BIGSERIAL PRIMARY KEY,
    created_at        TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at        TIMESTAMPTZ NOT NULL DEFAULT now(),
    deleted_at        TIMESTAMPTZ,
    user_id           BIGINT NOT NULL UNIQUE REFERENCES users(id) ON DELETE CASCADE,
    specialization    VARCHAR(100) NOT NULL,
    hospital          VARCHAR(150),
    country           VARCHAR(100),
    city              VARCHAR(100),
    years_experience  INTEGER CHECK (years_experience >= 0),
    license_number    VARCHAR(100) UNIQUE,
    education         VARCHAR(255),
    languages         VARCHAR(255),
    bio               TEXT,
    profile_image_url VARCHAR(255),
    license_verified  BOOLEAN NOT NULL DEFAULT false
);
CREATE INDEX IF NOT EXISTS idx_profiles_deleted_at ON profiles (deleted_at);

-- ----------------------------------------------------------------------------
-- communities  (internal/communities.Community)
-- ----------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS communities (
    id          BIGSERIAL PRIMARY KEY,
    created_at  TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at  TIMESTAMPTZ NOT NULL DEFAULT now(),
    deleted_at  TIMESTAMPTZ,
    name        VARCHAR(100) NOT NULL UNIQUE,
    description TEXT,
    category    VARCHAR(100) NOT NULL,
    creator_id  BIGINT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    banner_url  VARCHAR(255),
    is_private  BOOLEAN NOT NULL DEFAULT false
);
CREATE INDEX IF NOT EXISTS idx_communities_deleted_at ON communities (deleted_at);

-- ----------------------------------------------------------------------------
-- posts  (internal/posts.Post)
-- ----------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS posts (
    id             BIGSERIAL PRIMARY KEY,
    created_at     TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at     TIMESTAMPTZ NOT NULL DEFAULT now(),
    deleted_at     TIMESTAMPTZ,
    community_id   BIGINT NOT NULL REFERENCES communities(id) ON DELETE CASCADE,
    author_id      BIGINT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    title          VARCHAR(255) NOT NULL,
    content        TEXT NOT NULL,
    content_type   VARCHAR(50) NOT NULL DEFAULT 'markdown',
    image_url      TEXT,
    is_anonymous   BOOLEAN NOT NULL DEFAULT false,
    is_edited      BOOLEAN NOT NULL DEFAULT false,
    is_pinned      BOOLEAN NOT NULL DEFAULT false,
    is_locked      BOOLEAN NOT NULL DEFAULT false,
    likes_count    INTEGER NOT NULL DEFAULT 0,
    comments_count INTEGER NOT NULL DEFAULT 0,
    views_count    INTEGER NOT NULL DEFAULT 0
);
CREATE INDEX IF NOT EXISTS idx_posts_community_id ON posts (community_id);
CREATE INDEX IF NOT EXISTS idx_posts_author_id    ON posts (author_id);
CREATE INDEX IF NOT EXISTS idx_posts_deleted_at   ON posts (deleted_at);

-- posts.Attachment
CREATE TABLE IF NOT EXISTS attachments (
    id         BIGSERIAL PRIMARY KEY,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    deleted_at TIMESTAMPTZ,
    post_id    BIGINT NOT NULL REFERENCES posts(id) ON DELETE CASCADE,
    file_name  TEXT,
    file_url   TEXT,
    file_type  TEXT,
    file_size  BIGINT
);
CREATE INDEX IF NOT EXISTS idx_attachments_post_id    ON attachments (post_id);
CREATE INDEX IF NOT EXISTS idx_attachments_deleted_at ON attachments (deleted_at);

-- posts.Tag
CREATE TABLE IF NOT EXISTS tags (
    id         BIGSERIAL PRIMARY KEY,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    deleted_at TIMESTAMPTZ,
    name       VARCHAR(100) UNIQUE
);
CREATE INDEX IF NOT EXISTS idx_tags_deleted_at ON tags (deleted_at);

-- posts <-> tags many2many join table (GORM default name: post_tags)
CREATE TABLE IF NOT EXISTS post_tags (
    post_id BIGINT NOT NULL REFERENCES posts(id) ON DELETE CASCADE,
    tag_id  BIGINT NOT NULL REFERENCES tags(id) ON DELETE CASCADE,
    PRIMARY KEY (post_id, tag_id)
);

-- posts.Poll
CREATE TABLE IF NOT EXISTS polls (
    id         BIGSERIAL PRIMARY KEY,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    deleted_at TIMESTAMPTZ,
    post_id    BIGINT NOT NULL REFERENCES posts(id) ON DELETE CASCADE,
    question   VARCHAR(255) NOT NULL,
    expires_at TIMESTAMPTZ
);
CREATE INDEX IF NOT EXISTS idx_polls_post_id    ON polls (post_id);
CREATE INDEX IF NOT EXISTS idx_polls_deleted_at ON polls (deleted_at);

-- posts.PollOption
CREATE TABLE IF NOT EXISTS poll_options (
    id         BIGSERIAL PRIMARY KEY,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    deleted_at TIMESTAMPTZ,
    poll_id    BIGINT NOT NULL REFERENCES polls(id) ON DELETE CASCADE,
    text       VARCHAR(255) NOT NULL,
    votes      INTEGER NOT NULL DEFAULT 0
);
CREATE INDEX IF NOT EXISTS idx_poll_options_poll_id    ON poll_options (poll_id);
CREATE INDEX IF NOT EXISTS idx_poll_options_deleted_at ON poll_options (deleted_at);

-- ----------------------------------------------------------------------------
-- comments  (internal/comments.Comment)
-- ----------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS comments (
    id         BIGSERIAL PRIMARY KEY,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    deleted_at TIMESTAMPTZ,
    post_id    BIGINT NOT NULL REFERENCES posts(id) ON DELETE CASCADE,
    author_id  BIGINT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    content    TEXT NOT NULL,
    is_edited  BOOLEAN NOT NULL DEFAULT false
);
CREATE INDEX IF NOT EXISTS idx_comments_post_id    ON comments (post_id);
CREATE INDEX IF NOT EXISTS idx_comments_author_id  ON comments (author_id);
CREATE INDEX IF NOT EXISTS idx_comments_deleted_at ON comments (deleted_at);

-- ----------------------------------------------------------------------------
-- reactions  (internal/reactions.Reaction)
-- ----------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS reactions (
    id            BIGSERIAL PRIMARY KEY,
    created_at    TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at    TIMESTAMPTZ NOT NULL DEFAULT now(),
    deleted_at    TIMESTAMPTZ,
    post_id       BIGINT NOT NULL REFERENCES posts(id) ON DELETE CASCADE,
    user_id       BIGINT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    reaction_type VARCHAR(20) NOT NULL
);
CREATE INDEX IF NOT EXISTS idx_reactions_post_id    ON reactions (post_id);
CREATE INDEX IF NOT EXISTS idx_reactions_user_id    ON reactions (user_id);
CREATE INDEX IF NOT EXISTS idx_reactions_deleted_at ON reactions (deleted_at);

-- ----------------------------------------------------------------------------
-- messages  (internal/messages.Message)
-- ----------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS messages (
    id             BIGSERIAL PRIMARY KEY,
    created_at     TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at     TIMESTAMPTZ NOT NULL DEFAULT now(),
    deleted_at     TIMESTAMPTZ,
    sender_id      BIGINT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    receiver_id    BIGINT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    content        TEXT NOT NULL,
    message_type   VARCHAR(20) NOT NULL DEFAULT 'text',
    attachment_url TEXT,
    reply_to_id    BIGINT REFERENCES messages(id) ON DELETE SET NULL,
    is_read        BOOLEAN NOT NULL DEFAULT false,
    is_edited      BOOLEAN NOT NULL DEFAULT false
);
CREATE INDEX IF NOT EXISTS idx_messages_sender_id   ON messages (sender_id);
CREATE INDEX IF NOT EXISTS idx_messages_receiver_id ON messages (receiver_id);
CREATE INDEX IF NOT EXISTS idx_messages_deleted_at  ON messages (deleted_at);

-- ----------------------------------------------------------------------------
-- notifications  (internal/notifications.Notification)
-- ----------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS notifications (
    id             BIGSERIAL PRIMARY KEY,
    created_at     TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at     TIMESTAMPTZ NOT NULL DEFAULT now(),
    deleted_at     TIMESTAMPTZ,
    user_id        BIGINT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    actor_id       BIGINT REFERENCES users(id) ON DELETE SET NULL,
    type           VARCHAR(50) NOT NULL,
    title          VARCHAR(255) NOT NULL,
    message        TEXT NOT NULL,
    reference_id   BIGINT,
    reference_type VARCHAR(50),
    is_read        BOOLEAN NOT NULL DEFAULT false,
    action_url     TEXT
);
CREATE INDEX IF NOT EXISTS idx_notifications_user_id    ON notifications (user_id);
CREATE INDEX IF NOT EXISTS idx_notifications_actor_id   ON notifications (actor_id);
CREATE INDEX IF NOT EXISTS idx_notifications_type       ON notifications (type);
CREATE INDEX IF NOT EXISTS idx_notifications_deleted_at ON notifications (deleted_at);

-- ----------------------------------------------------------------------------
-- hospitals  (internal/hospitals.Hospital) — created before appointments-adjacent
-- tables that reference it
-- ----------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS hospitals (
    id             BIGSERIAL PRIMARY KEY,
    created_at     TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at     TIMESTAMPTZ NOT NULL DEFAULT now(),
    deleted_at     TIMESTAMPTZ,
    name           VARCHAR(255) NOT NULL,
    description    TEXT,
    email          VARCHAR(255) UNIQUE,
    phone          VARCHAR(50),
    website        VARCHAR(255),
    address        VARCHAR(255),
    city           VARCHAR(100),
    state          VARCHAR(100),
    country        VARCHAR(100),
    zip_code       VARCHAR(20),
    license_number VARCHAR(100) UNIQUE,
    is_active      BOOLEAN NOT NULL DEFAULT true
);
CREATE INDEX IF NOT EXISTS idx_hospitals_deleted_at ON hospitals (deleted_at);

-- ----------------------------------------------------------------------------
-- appointments  (internal/appointments.Appointment)
-- ----------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS appointments (
    id                BIGSERIAL PRIMARY KEY,
    created_at        TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at        TIMESTAMPTZ NOT NULL DEFAULT now(),
    deleted_at        TIMESTAMPTZ,
    doctor_id         BIGINT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    patient_id        BIGINT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    appointment_time  TIMESTAMPTZ NOT NULL,
    duration_minutes  INTEGER NOT NULL DEFAULT 30,
    status            VARCHAR(20) NOT NULL DEFAULT 'pending',
    reason            TEXT,
    notes             TEXT,
    meeting_link      VARCHAR(255)
);
CREATE INDEX IF NOT EXISTS idx_appointments_doctor_id  ON appointments (doctor_id);
CREATE INDEX IF NOT EXISTS idx_appointments_patient_id ON appointments (patient_id);
CREATE INDEX IF NOT EXISTS idx_appointments_deleted_at ON appointments (deleted_at);

-- ----------------------------------------------------------------------------
-- medical_records  (internal/medicalrecords.MedicalRecord)
-- ----------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS medical_records (
    id              BIGSERIAL PRIMARY KEY,
    created_at      TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at      TIMESTAMPTZ NOT NULL DEFAULT now(),
    deleted_at      TIMESTAMPTZ,
    patient_id      BIGINT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    doctor_id       BIGINT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    diagnosis       TEXT NOT NULL,
    symptoms        TEXT,
    treatment       TEXT,
    prescription    TEXT,
    notes           TEXT,
    follow_up_date  TIMESTAMPTZ
);
CREATE INDEX IF NOT EXISTS idx_medical_records_patient_id ON medical_records (patient_id);
CREATE INDEX IF NOT EXISTS idx_medical_records_doctor_id  ON medical_records (doctor_id);
CREATE INDEX IF NOT EXISTS idx_medical_records_deleted_at ON medical_records (deleted_at);

-- ----------------------------------------------------------------------------
-- uploads  (internal/uploads.Upload)
-- ----------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS uploads (
    id                 BIGSERIAL PRIMARY KEY,
    created_at         TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at         TIMESTAMPTZ NOT NULL DEFAULT now(),
    deleted_at         TIMESTAMPTZ,
    user_id            BIGINT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    appointment_id     BIGINT REFERENCES appointments(id) ON DELETE SET NULL,
    medical_record_id  BIGINT REFERENCES medical_records(id) ON DELETE SET NULL,
    hospital_id        BIGINT REFERENCES hospitals(id) ON DELETE SET NULL,
    file_name          VARCHAR(255) NOT NULL,
    original_name      VARCHAR(255) NOT NULL,
    file_type          VARCHAR(100) NOT NULL,
    file_size          BIGINT NOT NULL,
    file_path          VARCHAR(500) NOT NULL,
    description        TEXT,
    is_public          BOOLEAN NOT NULL DEFAULT false
);
CREATE INDEX IF NOT EXISTS idx_uploads_user_id           ON uploads (user_id);
CREATE INDEX IF NOT EXISTS idx_uploads_appointment_id    ON uploads (appointment_id);
CREATE INDEX IF NOT EXISTS idx_uploads_medical_record_id ON uploads (medical_record_id);
CREATE INDEX IF NOT EXISTS idx_uploads_hospital_id       ON uploads (hospital_id);
CREATE INDEX IF NOT EXISTS idx_uploads_deleted_at        ON uploads (deleted_at);

-- ----------------------------------------------------------------------------
-- video_consultations  (internal/video_consultations.VideoConsultation)
-- ----------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS video_consultations (
    id             BIGSERIAL PRIMARY KEY,
    created_at     TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at     TIMESTAMPTZ NOT NULL DEFAULT now(),
    deleted_at     TIMESTAMPTZ,
    appointment_id BIGINT NOT NULL UNIQUE REFERENCES appointments(id) ON DELETE CASCADE,
    doctor_id      BIGINT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    patient_id     BIGINT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    room_id        VARCHAR(255) NOT NULL UNIQUE,
    session_key    VARCHAR(255),
    scheduled_at   TIMESTAMPTZ,
    started_at     TIMESTAMPTZ,
    ended_at       TIMESTAMPTZ,
    status         VARCHAR(20) NOT NULL DEFAULT 'scheduled',
    notes          TEXT
);
CREATE INDEX IF NOT EXISTS idx_video_consultations_doctor_id  ON video_consultations (doctor_id);
CREATE INDEX IF NOT EXISTS idx_video_consultations_patient_id ON video_consultations (patient_id);
CREATE INDEX IF NOT EXISTS idx_video_consultations_deleted_at ON video_consultations (deleted_at);

-- ----------------------------------------------------------------------------
-- presences  (internal/presence.Presence)
-- ----------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS presences (
    id            BIGSERIAL PRIMARY KEY,
    created_at    TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at    TIMESTAMPTZ NOT NULL DEFAULT now(),
    deleted_at    TIMESTAMPTZ,
    user_id       BIGINT NOT NULL UNIQUE REFERENCES users(id) ON DELETE CASCADE,
    is_online     BOOLEAN NOT NULL DEFAULT false,
    last_seen     TIMESTAMPTZ,
    connection_id VARCHAR(255)
);
CREATE INDEX IF NOT EXISTS idx_presences_deleted_at ON presences (deleted_at);

-- ----------------------------------------------------------------------------
-- searches  (internal/search.Search)
-- ----------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS searches (
    id           BIGSERIAL PRIMARY KEY,
    created_at   TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at   TIMESTAMPTZ NOT NULL DEFAULT now(),
    deleted_at   TIMESTAMPTZ,
    user_id      BIGINT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    query        VARCHAR(255) NOT NULL,
    type         VARCHAR(30) NOT NULL,
    result_count INTEGER NOT NULL DEFAULT 0
);
CREATE INDEX IF NOT EXISTS idx_searches_user_id    ON searches (user_id);
CREATE INDEX IF NOT EXISTS idx_searches_query      ON searches (query);
CREATE INDEX IF NOT EXISTS idx_searches_deleted_at ON searches (deleted_at);

-- ----------------------------------------------------------------------------
-- admins  (internal/admin.Admin — audit log)
-- ----------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS admins (
    id          BIGSERIAL PRIMARY KEY,
    created_at  TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at  TIMESTAMPTZ NOT NULL DEFAULT now(),
    deleted_at  TIMESTAMPTZ,
    admin_id    BIGINT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    resource_id BIGINT NOT NULL,
    resource    VARCHAR(30) NOT NULL,
    action      VARCHAR(30) NOT NULL,
    description TEXT
);
CREATE INDEX IF NOT EXISTS idx_admins_admin_id    ON admins (admin_id);
CREATE INDEX IF NOT EXISTS idx_admins_resource_id ON admins (resource_id);
CREATE INDEX IF NOT EXISTS idx_admins_resource    ON admins (resource);
CREATE INDEX IF NOT EXISTS idx_admins_action      ON admins (action);
CREATE INDEX IF NOT EXISTS idx_admins_deleted_at  ON admins (deleted_at);

-- ----------------------------------------------------------------------------
-- payments  (internal/payments.Payment)
-- ----------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS payments (
    id                     BIGSERIAL PRIMARY KEY,
    created_at             TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at             TIMESTAMPTZ NOT NULL DEFAULT now(),
    deleted_at             TIMESTAMPTZ,
    appointment_id         BIGINT NOT NULL REFERENCES appointments(id) ON DELETE CASCADE,
    patient_id             BIGINT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    doctor_id              BIGINT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    hospital_id            BIGINT REFERENCES hospitals(id) ON DELETE SET NULL,
    amount                 NUMERIC(12,2) NOT NULL,
    currency               VARCHAR(10) NOT NULL DEFAULT 'KES',
    method                 VARCHAR(30) NOT NULL,
    status                 VARCHAR(30) NOT NULL DEFAULT 'pending',
    transaction_reference  VARCHAR(255) UNIQUE,
    description            TEXT,
    paid_at                TIMESTAMPTZ
);
CREATE INDEX IF NOT EXISTS idx_payments_appointment_id ON payments (appointment_id);
CREATE INDEX IF NOT EXISTS idx_payments_patient_id     ON payments (patient_id);
CREATE INDEX IF NOT EXISTS idx_payments_doctor_id      ON payments (doctor_id);
CREATE INDEX IF NOT EXISTS idx_payments_hospital_id    ON payments (hospital_id);
CREATE INDEX IF NOT EXISTS idx_payments_deleted_at     ON payments (deleted_at);

-- ----------------------------------------------------------------------------
-- doctor_schedules  (internal/doctor_schedules.DoctorSchedule)
-- ----------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS doctor_schedules (
    id                     BIGSERIAL PRIMARY KEY,
    created_at             TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at             TIMESTAMPTZ NOT NULL DEFAULT now(),
    deleted_at             TIMESTAMPTZ,
    doctor_id              BIGINT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    day                    VARCHAR(20) NOT NULL,
    start_time             VARCHAR(5) NOT NULL,
    end_time               VARCHAR(5) NOT NULL,
    break_start            VARCHAR(5),
    break_end              VARCHAR(5),
    consultation_duration  INTEGER NOT NULL DEFAULT 30,
    max_patients           INTEGER NOT NULL DEFAULT 20,
    is_active              BOOLEAN NOT NULL DEFAULT true
);
CREATE INDEX IF NOT EXISTS idx_doctor_schedules_doctor_id  ON doctor_schedules (doctor_id);
CREATE INDEX IF NOT EXISTS idx_doctor_schedules_day        ON doctor_schedules (day);
CREATE INDEX IF NOT EXISTS idx_doctor_schedules_deleted_at ON doctor_schedules (deleted_at);

-- ----------------------------------------------------------------------------
-- availabilities  (internal/availability.Availability)
-- ----------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS availabilities (
    id             BIGSERIAL PRIMARY KEY,
    created_at     TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at     TIMESTAMPTZ NOT NULL DEFAULT now(),
    deleted_at     TIMESTAMPTZ,
    doctor_id      BIGINT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    schedule_id    BIGINT NOT NULL REFERENCES doctor_schedules(id) ON DELETE CASCADE,
    date           DATE NOT NULL,
    start_time     VARCHAR(5) NOT NULL,
    end_time       VARCHAR(5) NOT NULL,
    status         VARCHAR(20) NOT NULL DEFAULT 'available',
    appointment_id BIGINT REFERENCES appointments(id) ON DELETE SET NULL,
    notes          TEXT
);
CREATE INDEX IF NOT EXISTS idx_availabilities_doctor_id      ON availabilities (doctor_id);
CREATE INDEX IF NOT EXISTS idx_availabilities_schedule_id    ON availabilities (schedule_id);
CREATE INDEX IF NOT EXISTS idx_availabilities_date           ON availabilities (date);
CREATE INDEX IF NOT EXISTS idx_availabilities_appointment_id ON availabilities (appointment_id);
CREATE INDEX IF NOT EXISTS idx_availabilities_deleted_at     ON availabilities (deleted_at);

-- ----------------------------------------------------------------------------
-- medical_specialties  (internal/medical_specialties.MedicalSpecialty)
-- ----------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS medical_specialties (
    id          BIGSERIAL PRIMARY KEY,
    created_at  TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at  TIMESTAMPTZ NOT NULL DEFAULT now(),
    deleted_at  TIMESTAMPTZ,
    name        VARCHAR(100) NOT NULL UNIQUE,
    code        VARCHAR(20) NOT NULL UNIQUE,
    description TEXT,
    is_active   BOOLEAN NOT NULL DEFAULT true
);
CREATE INDEX IF NOT EXISTS idx_medical_specialties_deleted_at ON medical_specialties (deleted_at);

-- ----------------------------------------------------------------------------
-- doctor_reviews  (internal/doctor_reviews.DoctorReview)
-- ----------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS doctor_reviews (
    id             BIGSERIAL PRIMARY KEY,
    created_at     TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at     TIMESTAMPTZ NOT NULL DEFAULT now(),
    deleted_at     TIMESTAMPTZ,
    doctor_id      BIGINT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    patient_id     BIGINT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    appointment_id BIGINT NOT NULL UNIQUE REFERENCES appointments(id) ON DELETE CASCADE,
    rating         INTEGER NOT NULL CHECK (rating >= 1 AND rating <= 5),
    title          VARCHAR(150),
    comment        TEXT,
    is_anonymous   BOOLEAN NOT NULL DEFAULT false,
    is_published   BOOLEAN NOT NULL DEFAULT true
);
CREATE INDEX IF NOT EXISTS idx_doctor_reviews_doctor_id  ON doctor_reviews (doctor_id);
CREATE INDEX IF NOT EXISTS idx_doctor_reviews_patient_id ON doctor_reviews (patient_id);
CREATE INDEX IF NOT EXISTS idx_doctor_reviews_deleted_at ON doctor_reviews (deleted_at);

-- ----------------------------------------------------------------------------
-- prescriptions & prescription_items  (internal/prescriptions)
-- ----------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS prescriptions (
    id             BIGSERIAL PRIMARY KEY,
    created_at     TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at     TIMESTAMPTZ NOT NULL DEFAULT now(),
    deleted_at     TIMESTAMPTZ,
    doctor_id      BIGINT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    patient_id     BIGINT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    appointment_id BIGINT NOT NULL REFERENCES appointments(id) ON DELETE CASCADE,
    diagnosis      TEXT NOT NULL,
    notes          TEXT,
    status         VARCHAR(20) NOT NULL DEFAULT 'active',
    issued_at      TIMESTAMPTZ,
    expires_at     TIMESTAMPTZ
);
CREATE INDEX IF NOT EXISTS idx_prescriptions_doctor_id      ON prescriptions (doctor_id);
CREATE INDEX IF NOT EXISTS idx_prescriptions_patient_id     ON prescriptions (patient_id);
CREATE INDEX IF NOT EXISTS idx_prescriptions_appointment_id ON prescriptions (appointment_id);
CREATE INDEX IF NOT EXISTS idx_prescriptions_deleted_at     ON prescriptions (deleted_at);

CREATE TABLE IF NOT EXISTS prescription_items (
    id               BIGSERIAL PRIMARY KEY,
    created_at       TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at       TIMESTAMPTZ NOT NULL DEFAULT now(),
    prescription_id  BIGINT NOT NULL REFERENCES prescriptions(id) ON DELETE CASCADE,
    medication_name  VARCHAR(255) NOT NULL,
    dosage           VARCHAR(100) NOT NULL,
    frequency        VARCHAR(100) NOT NULL,
    duration         VARCHAR(100) NOT NULL,
    instructions     TEXT,
    quantity         INTEGER NOT NULL,
    refills          INTEGER NOT NULL DEFAULT 0
);
CREATE INDEX IF NOT EXISTS idx_prescription_items_prescription_id ON prescription_items (prescription_id);

-- ----------------------------------------------------------------------------
-- vitals  (internal/vitals.Vital)
-- ----------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS vitals (
    id                BIGSERIAL PRIMARY KEY,
    created_at        TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at        TIMESTAMPTZ NOT NULL DEFAULT now(),
    patient_id        BIGINT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    doctor_id         BIGINT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    appointment_id    BIGINT REFERENCES appointments(id) ON DELETE SET NULL,
    temperature       NUMERIC(4,1),
    heart_rate        INTEGER,
    respiratory_rate  INTEGER,
    systolic_bp       INTEGER,
    diastolic_bp      INTEGER,
    oxygen_saturation INTEGER,
    weight            NUMERIC(5,2),
    height            NUMERIC(5,2),
    bmi               NUMERIC(5,2),
    notes             TEXT,
    recorded_at       TIMESTAMPTZ NOT NULL
);
CREATE INDEX IF NOT EXISTS idx_vitals_patient_id     ON vitals (patient_id);
CREATE INDEX IF NOT EXISTS idx_vitals_doctor_id      ON vitals (doctor_id);
CREATE INDEX IF NOT EXISTS idx_vitals_appointment_id ON vitals (appointment_id);

-- ----------------------------------------------------------------------------
-- allergies  (internal/allergies.Allergy)
-- ----------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS allergies (
    id           BIGSERIAL PRIMARY KEY,
    created_at   TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at   TIMESTAMPTZ NOT NULL DEFAULT now(),
    patient_id   BIGINT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    doctor_id    BIGINT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    allergen     VARCHAR(255) NOT NULL,
    type         VARCHAR(50) NOT NULL,
    severity     VARCHAR(20) NOT NULL,
    reaction     TEXT,
    status       VARCHAR(20) NOT NULL DEFAULT 'Active',
    notes        TEXT,
    recorded_at  TIMESTAMPTZ NOT NULL
);
CREATE INDEX IF NOT EXISTS idx_allergies_patient_id ON allergies (patient_id);
CREATE INDEX IF NOT EXISTS idx_allergies_doctor_id  ON allergies (doctor_id);

-- ----------------------------------------------------------------------------
-- lab_requests & lab_results  (internal/lab_requests, internal/lab_results)
-- ----------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS lab_requests (
    id              BIGSERIAL PRIMARY KEY,
    created_at      TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at      TIMESTAMPTZ NOT NULL DEFAULT now(),
    patient_id      BIGINT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    doctor_id       BIGINT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    appointment_id  BIGINT REFERENCES appointments(id) ON DELETE SET NULL,
    test_name       VARCHAR(255) NOT NULL,
    category        VARCHAR(100),
    priority        VARCHAR(20) NOT NULL DEFAULT 'Routine',
    clinical_notes  TEXT,
    reason          TEXT,
    status          VARCHAR(20) NOT NULL DEFAULT 'Pending',
    requested_at    TIMESTAMPTZ NOT NULL
);
CREATE INDEX IF NOT EXISTS idx_lab_requests_patient_id     ON lab_requests (patient_id);
CREATE INDEX IF NOT EXISTS idx_lab_requests_doctor_id      ON lab_requests (doctor_id);
CREATE INDEX IF NOT EXISTS idx_lab_requests_appointment_id ON lab_requests (appointment_id);

CREATE TABLE IF NOT EXISTS lab_results (
    id               BIGSERIAL PRIMARY KEY,
    created_at       TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at       TIMESTAMPTZ NOT NULL DEFAULT now(),
    lab_request_id   BIGINT NOT NULL REFERENCES lab_requests(id) ON DELETE CASCADE,
    patient_id       BIGINT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    doctor_id        BIGINT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    test_name        VARCHAR(255) NOT NULL,
    result           TEXT NOT NULL,
    reference_range  VARCHAR(255),
    units            VARCHAR(50),
    interpretation   TEXT,
    status           VARCHAR(20) NOT NULL DEFAULT 'Completed',
    remarks          TEXT,
    performed_at     TIMESTAMPTZ NOT NULL,
    verified_at      TIMESTAMPTZ
);
CREATE INDEX IF NOT EXISTS idx_lab_results_lab_request_id ON lab_results (lab_request_id);
CREATE INDEX IF NOT EXISTS idx_lab_results_patient_id     ON lab_results (patient_id);
CREATE INDEX IF NOT EXISTS idx_lab_results_doctor_id      ON lab_results (doctor_id);

-- ============================================================================
-- End of initial schema
-- ============================================================================
