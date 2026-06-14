Architecture Design : 

                    React UI
                        |
                        |
                        ▼
                 Gin REST API
                        |
                        |
         +--------------+-------------+
         |                            |
         ▼                            ▼
    PostgreSQL                  Scheduler
                                      |
                                      |
                                      ▼
                              Ingestion Worker
                                      |
                                      |
                                      ▼
                               RSS Providers
                                      |
                                      ▼
                              Sentiment Service


CREATE TABLE entities (
    id UUID PRIMARY KEY,

    name VARCHAR(255) NOT NULL,

    aliases TEXT[],

    status VARCHAR(20) NOT NULL,

    ingestion_interval_minutes INT NOT NULL,

    last_ingested_at TIMESTAMP,

    created_at TIMESTAMP,

    updated_at TIMESTAMP
);

CREATE TABLE news_articles(
    id UUID PRIMARY KEY,

    entity_id UUID NOT NULL,

    title TEXT NOT NULL,

    description TEXT,

    url TEXT NOT NULL,

    source VARCHAR(255),

    published_date TIMESTAMP,

    sentiment VARCHAR(20),

    sentiment_score FLOAT,

    url_hash VARCHAR(64),

    dedupe_key VARCHAR(64),

    created_at TIMESTAMP,

    updated_at TIMESTAMP
);

CREATE UNIQUE INDEX idx_url_hash
ON news_articles(url_hash);

CREATE UNIQUE INDEX idx_entity_dedupe
ON news_articles(entity_id,dedupe_key);