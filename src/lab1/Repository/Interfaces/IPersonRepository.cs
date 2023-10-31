using Entity;

namespace lab1.Repository.Interfaces
{
    public interface IPersonRepository : IRepository<Person>
    {
        Person GetUserByID(long id);

    }
}
