create table if not exists SystemResourceTypes (
    id int unsigned auto_increment primary key,
    name varchar(128) charset utf8
);

create table if not exists SystemResources (
    id bigint unsigned auto_increment primary key,
    resourcetype int unsigned,
    name varchar(128) charset utf8,
    rarity float unsigned,

    constraint `fk_SystemResources_SystemResourceTypes`
        foreign key (resourcetype) references SystemResourceTypes (id)
        on delete cascade
);

create table if not exists SystemLocations (
    id bigint unsigned auto_increment primary key,
    coordx bigint,
    coordy bigint
);

create table if not exists GenResourceLocations (
    resourcetype int unsigned,
    location bigint unsigned,
    capacity int unsigned,

    constraint `fk_GenResourceLocations_SystemResourceTypes`
        foreign key (resourcetype) references SystemResourceTypes (id)
        on delete cascade,
        foreign key (location) references SystemLocations (id)
        on delete cascade
);