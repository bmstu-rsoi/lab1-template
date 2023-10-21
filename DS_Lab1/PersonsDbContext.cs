﻿using DS_Lab1.Models;
using Microsoft.EntityFrameworkCore;

namespace DS_Lab1
{
    public class PersonsDbContext : DbContext
    {
        public DbSet<Person> Persons { get; set; }

        public PersonsDbContext()
        {
            Database.EnsureCreated();
        }

        public PersonsDbContext(DbContextOptions<PersonsDbContext> options) : base(options)
        {
            Database.EnsureCreated();
        }

        protected override void OnModelCreating(ModelBuilder modelBuilder)
        {
            modelBuilder.Entity<Person>(entity =>
            {
                entity.ToTable("persons_table");

                entity.HasKey(p => p.Id).HasName("id");

                entity.Property(e => e.Name).IsRequired().HasColumnName("name");

                entity.Property(e => e.Age).HasColumnName("age");

                entity.Property(e => e.Adress).HasColumnName("adress");

                entity.Property(e => e.Work).HasColumnName("work");
            });
        }
    }
}
