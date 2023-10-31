using Microsoft.EntityFrameworkCore;
using RSOI_Lab_01.Entities;
using RSOI_Lab_01.Models;
using System.Collections.Generic;
using System.Reflection.Emit;

namespace RSOI_Lab_01
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
                entity.HasKey(p => p.Id).HasName("Id");

                entity.ToTable("persons_table");

                entity.Property(p => p.Id).ValueGeneratedOnAdd().HasColumnName("Id");

                //NpgsqlPropertyBuilderExtensions.UseIdentityByDefaultColumn(entity.Property(p => p.Id));

                entity.Property(e => e.Name).IsRequired().HasColumnName("name");

                entity.Property(e => e.Age).HasColumnName("age");

                entity.Property(e => e.Adress).HasColumnName("adress");

                entity.Property(e => e.Work).HasColumnName("work");
            });
        }
    }
}
