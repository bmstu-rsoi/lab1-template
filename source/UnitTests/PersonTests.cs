using Moq;
using RSOI_Lab_01.Entities;
using RSOI_Lab_01.Interfaces;
using RSOI_Lab_01.Models;
using RSOI_Lab_01.Services;

namespace UnitTests
{
    public class PersonTests
    {
        [Fact]
        public async Task GetAllPersonsTest()
        {
            List<Person> expected = new() { new Person(1, "Nikita", 22, "Address", "MTS_Digital") };
            var mock = new Mock<IPersonsRepository>();
            mock.Setup(a => a.GetAllAsync()).ReturnsAsync(expected);
            var service = new PersonsService(mock.Object);

            var result = await service.GetAllAsync();

            Assert.NotNull(result);
        }

        [Fact]
        public async Task GetUndeclaretedPersonTest()
        {
            var mock = new Mock<IPersonsRepository>();
            mock.Setup(a => a.GetAsync(-1));
            var service = new PersonsService(mock.Object);

            var person = await service.GetAsync(-1);

            Assert.Null(person?.Id);
        }

        [Fact]
        public async Task GetPersonTest()
        {
            var expected = new Person(1, "Nikita", 22, "Address", "MTS_Digital");
            var mock = new Mock<IPersonsRepository>();
            mock.Setup(a => a.GetAsync(1)).ReturnsAsync(expected);
            var service = new PersonsService(mock.Object);

            var actual = await service.GetAsync(1);

            Assert.Equal(expected.Id, actual?.Id);
        }

        [Fact]
        public async Task CreatePersonTest()
        {
            var person = new Person(1, "Nikita", 22, "Address", "MTS_Digital");
            var mock = new Mock<IPersonsRepository>();
            var service = new PersonsService(mock.Object);

            var result = await service.CreateAsync(new PersonRequest(person.Name, person.Age, person.Adress, person.Work));

            mock.Verify(a => a.CreateAsync(person));
        }

        [Fact]
        public async Task DeletePersonTest()
        {
            List<Person> expected = new();
            var person = new Person(1, "Nikita", 22, "Address", "MTS_Digital");
            var mock = new Mock<IPersonsRepository>();
            var service = new PersonsService(mock.Object);

            await service.DeleteAsync((int)person.Id);

            mock.Verify(a => a.DeleteAsync((int)person.Id));
        }
    }
}