using Microsoft.VisualStudio.TestTools.UnitTesting;
using Controllers;
using System;
using Microsoft.EntityFrameworkCore;
using Microsoft.Extensions.Logging.Abstractions;
using Moq;
using System.Collections.Generic;
using MyBD;
using lab1.Repository.Interfaces;
using lab1.Repository.Implementation;
using Entity;

namespace TestProject
{
    [TestClass]
    public class UnitTests
    {

        private BDContext db;
        private IPersonRepository personRep;

        [TestInitialize]
        public void Init()
        {
            var options = new DbContextOptionsBuilder<BDContext>()
                .UseInMemoryDatabase("PersonsControllerTest")
                .Options;
            db = new BDContext(options);
            db.Database.EnsureDeleted();
            db.Database.EnsureCreated();
            personRep = new PersonRepository(db);

        }

        private Boolean PersonCompare(Person first, Person second)
        {
            return first.Name == second.Name && first.Age == second.Age && first.Address == second.Address && first.Work == second.Work ? true : false;
        }

        [TestMethod]
        public void TestPatch()
        {
            Person changedPerson = new Person { Id = 2, Name = "Alan", Age = 47, Address = "Erevan", Work = "Yandex" };
            Person checkPerson;
            List<Person> persons = new List<Person>();
            persons.Add(new Person { Id = 1, Name = "Alex", Age = 24, Address = "Los Angeles", Work = "Google" });
            persons.Add(new Person { Id = 2, Name = "Ivan", Age = 20, Address = "Moscow", Work = "Yandex" });
            persons.Add(new Person { Id = 3, Name = "Dasha", Age = 30, Address = "Beijing", Work = "TikTok" });

            foreach (Person person in persons)
            {
                checkPerson = personRep.Add(person);
                Assert.IsTrue(PersonCompare(person, checkPerson));
            }

            checkPerson = personRep.Update(changedPerson);
            Assert.IsTrue(PersonCompare(changedPerson, checkPerson));

            checkPerson = personRep.GetUserByID(changedPerson.Id);
            Assert.IsTrue(PersonCompare(changedPerson, checkPerson));
        }

        [TestMethod]
        public void TestOkDelete()
        {
            Person checkPerson;
            List<Person> persons = new List<Person>();
            persons.Add(new Person { Id = 1, Name = "Alex", Age = 24, Address = "Los Angeles", Work = "Google" });
            persons.Add(new Person { Id = 2, Name = "Ivan", Age = 20, Address = "Moscow", Work = "Yandex" });
            persons.Add(new Person { Id = 3, Name = "Dasha", Age = 30, Address = "Beijing", Work = "TikTok" });

            foreach (Person person in persons)
            {
                checkPerson = personRep.Add(person);
                Assert.IsTrue(PersonCompare(person, checkPerson));
            }

            int code = personRep.Delete(persons[1].Id);
            Assert.IsTrue(code == 0);

            checkPerson = personRep.GetUserByID(persons[1].Id);
            Assert.IsNull(checkPerson);
        }

        [TestMethod]
        public void TestNotFoundDelete()
        {
            Person checkPerson;
            List<Person> persons = new List<Person>();
            persons.Add(new Person { Id = 1, Name = "Alex", Age = 24, Address = "Los Angeles", Work = "Google" });
            persons.Add(new Person { Id = 2, Name = "Ivan", Age = 20, Address = "Moscow", Work = "Yandex" });
            persons.Add(new Person { Id = 3, Name = "Dasha", Age = 30, Address = "Beijing", Work = "TikTok" });

            foreach (Person person in persons)
            {
                checkPerson = personRep.Add(person);
                Assert.IsTrue(PersonCompare(person, checkPerson));
            }

            int code = personRep.Delete(5);
            Assert.IsTrue(code == 1);
        }

        [TestMethod]
        public void TestGetAll()
        {
            Person checkPerson;
            List<Person> persons = new List<Person>();
            persons.Add(new Person { Id = 1, Name = "Alex", Age = 24, Address = "Los Angeles", Work = "Google" });
            persons.Add(new Person { Id = 2, Name = "Ivan", Age = 20, Address = "Moscow", Work = "Yandex" });
            persons.Add(new Person { Id = 3, Name = "Dasha", Age = 30, Address = "Beijing", Work = "TikTok" });

            foreach (Person person in persons)
            {
                checkPerson = personRep.Add(person);
                Assert.IsTrue(PersonCompare(person, checkPerson));
            }

            foreach (Person person in persons)
            {
                checkPerson = personRep.GetUserByID(person.Id);
                Assert.IsTrue(PersonCompare(person, checkPerson));
            }
        }

        [TestMethod]
        public void TestGetById()
        {
            Person checkPerson;
            List<Person> persons = new List<Person>();
            persons.Add(new Person { Id = 1, Name = "Alex", Age = 24, Address = "Los Angeles", Work = "Google" });
            persons.Add(new Person { Id = 2, Name = "Ivan", Age = 20, Address = "Moscow", Work = "Yandex" });
            persons.Add(new Person { Id = 3, Name = "Dasha", Age = 30, Address = "Beijing", Work = "TikTok" });

            foreach (Person person in persons)
            {
                checkPerson = personRep.Add(person);
                Assert.IsTrue(PersonCompare(person, checkPerson));
            }

            checkPerson = personRep.GetUserByID(persons[1].Id);
            Assert.IsTrue(PersonCompare(persons[1], checkPerson));

        }

        [TestMethod]
        public void TestPost()
        {
            Person person = new Person { Id = 1, Name = "Alan", Age = 20, Address = "Moscow", Work = "Yandex" };

            person = personRep.Add(person);

            Person checkPerson = personRep.GetUserByID(person.Id);
            Assert.IsTrue(PersonCompare(person, checkPerson));

        }
    }
}
