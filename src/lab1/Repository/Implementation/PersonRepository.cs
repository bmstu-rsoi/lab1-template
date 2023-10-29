using System;
using System.Collections.Generic;
using System.Linq;
using lab1.Repository.Interfaces;
using Entity;
using MyBD;

namespace lab1.Repository.Implementation
{
    public class PersonRepository : IPersonRepository, IDisposable
    {
        private readonly BDContext db;
        public PersonRepository(BDContext curDb)
        {
            db = curDb;
        }
        public Person Add(Person element)
        {
            element.Address ??= "";
            element.Name ??= "";
            element.Work ??= "";
            db.ChangeTracker.Clear();
            db.Person.Attach(element);
            db.Person.Add(element);
            db.SaveChanges();
            return element;

        }
        public List<Person> GetAll()
        {
            IQueryable<Person> users = db.Person;
            return users == null || users.Count() == 0 ? null : users.ToList();
        }
        public Person Update(long id, Person element)
        {
  
            Person person = db.Person.Find(id);
            if (person != null)
            {
                person.Id = element.Id;
                person.Name = element.Name ?? person.Name;
                person.Address = element.Address ?? person.Address;
                person.Work = element.Work ?? person.Work;
                person.Age = (element.Age < 0) ? person.Age : element.Age;
                db.ChangeTracker.Clear();
                db.Person.Attach(person);
                db.Person.Update(person);
                db.SaveChanges();
                return person;
            }
            else
            {
                return null;
            }

        }

        public int Delete(long id)
        {
            Person element = db.Person.Find(id);
            if (element == null)
            {
                return 1;
            }
            db.ChangeTracker.Clear();
            db.Person.Attach(element);
            db.Person.Remove(element);
            db.SaveChanges();
            return 0;

        }

        public Person GetUserByID(long id)
        {

            Person person = db.Person.Find(id);
            db.ChangeTracker.Clear();
            return person;
        }


        public void Dispose()
        {
            db.Dispose();
        }
    }
}
