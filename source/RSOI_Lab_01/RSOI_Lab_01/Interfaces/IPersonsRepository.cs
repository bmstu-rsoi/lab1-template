using RSOI_Lab_01.Entities;

namespace RSOI_Lab_01.Interfaces
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
