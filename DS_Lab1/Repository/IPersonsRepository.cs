using DS_Lab1.Models;

namespace DS_Lab1.Repository
{
    public interface IPersonsRepository
    {
        Task<IEnumerable<Person>> GetAllAsync();
        Task<Person?> GetAsync(int id);
        Task CreateAsync(Person person);
        void Update(Person person);
        Task DeleteAsync(int id);
        Task SaveAsync();
    }
}
