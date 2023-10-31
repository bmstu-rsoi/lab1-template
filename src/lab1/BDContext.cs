using System;
using System.Linq;
using Microsoft.EntityFrameworkCore;
using Entity;
#nullable disable

namespace MyBD
{
    public partial class BDContext : DbContext
    {
        private string ConnectionString { get; set; }
        public BDContext()
        {
            ConnectionString = "Host = postgres; Port = 5432; Database = persons; Username = postgres; Password = 1234";
        }

        public BDContext(DbContextOptions<BDContext> options)
            : base(options)
        {
        }

        public virtual DbSet<Person> Person { get; set; }
        
        protected override void OnConfiguring(DbContextOptionsBuilder optionsBuilder)
        {
            if (!optionsBuilder.IsConfigured)
            {
//#warning To protect potentially sensitive information in your connection string, you should move it out of source code. You can avoid scaffolding the connection string by using the Name= syntax to read it from configuration - see https://go.microsoft.com/fwlink/?linkid=2131148. For more guidance on storing connection strings, see http://go.microsoft.com/fwlink/?LinkId=723263.
                optionsBuilder.UseNpgsql(ConnectionString);
            }
        }

        protected override void OnModelCreating(ModelBuilder modelBuilder)
        {
            
           

            modelBuilder.HasAnnotation("Relational:Collation", "Russian_Russia.1251");

           


            modelBuilder.Entity<Person>(entity =>
            {
                entity.ToTable("persons");

                entity.Property(e => e.Id)
                    .ValueGeneratedNever()
                    .HasColumnName("id");

                entity.Property(e => e.Name)
                    .IsRequired()
                    .HasMaxLength(60)
                    .HasColumnName("name");

                entity.Property(e => e.Address)
                    .IsRequired()
                    .HasMaxLength(60)
                    .HasColumnName("address");

                entity.Property(e => e.Work)
                    .IsRequired()
                    .HasMaxLength(60)
                    .HasColumnName("work");

                entity.Property(e => e.Age).HasColumnName("age");
                
            });

            OnModelCreatingPartial(modelBuilder);
        }

        partial void OnModelCreatingPartial(ModelBuilder modelBuilder);
    }
}
