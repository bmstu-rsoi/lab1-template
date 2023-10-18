using DS_Lab1.Models;

namespace DS_Lab1.Repository
{
    public interface IPersonsRepository
    {
        IEnumerable<Person> GetAll();
        Person? Get(int id);
        void Create(Person person);
        void Update(Person person);
        void Delete(int id);
    }
}
